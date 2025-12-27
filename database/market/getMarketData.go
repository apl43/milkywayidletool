package market

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

func GetMarketData() {

	var marketData map[string]map[int]MarketItem
	var dataTimestamp int

	// 初始化Collector
	c := colly.NewCollector()

	// 访问每个页面时执行
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Printf("Link found: %q\n", e.Text)
	})

	// 在访问每个页面之前执行
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))

		var res ResponseJson
		if err := json.Unmarshal(r.Body, &res); err != nil {
			fmt.Println("解析失败: ", err)
			return
		}

		dataTimestamp = res.Timestamp
		marketData = make(map[string]map[int]MarketItem)
		for itemName, levelPrice := range res.MarketData {
			marketData[itemName] = make(map[int]MarketItem)
			for level, price := range levelPrice {
				levelInt, err := strconv.Atoi(level)
				if err != nil {
					fmt.Printf("等级转整数失败: %v\n", err)
					continue
				}
				marketData[itemName][levelInt] = MarketItem{
					Level: levelInt,
					Sell:  int64(price["a"]),
					Buy:   int64(price["b"]),
				}
			}
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})

	c.Visit("https://www.milkywayidle.com/game_data/marketplace.json")

	SaveMarketDataAsGob(&MarketDataFile{
		SaveTime:      time.Now(),
		DataTimestamp: dataTimestamp,
		MarketItems:   marketData,
	})
}

// 保存为GoB格式（二进制，高效）
func SaveMarketDataAsGob(data *MarketDataFile) error {
	// 创建目录
	historicalDataDir := "./database/market/historicalData"
	if err := os.MkdirAll(historicalDataDir, 0755); err != nil {
		return err
	}

	// 生成文件名
	currentLocalTime := time.Now().Local().Format("20060102_15:04:05")
	currentTimestamp := time.Now().Unix()
	filename := fmt.Sprintf("market_%s_%d.gob", currentLocalTime, currentTimestamp)
	filepath := filepath.Join(historicalDataDir, filename)

	// 创建文件
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 编码并保存
	encoder := gob.NewEncoder(file)
	return encoder.Encode(data)
}

// 加载最新的GoB文件
func LoadLatestMarketData() (*MarketDataFile, error) {
	historicalDataDir := "./database/market/historicalData"

	// 查找最新的gob文件
	files, err := filepath.Glob(filepath.Join(historicalDataDir, "market_*.gob"))
	if err != nil || len(files) == 0 {
		return nil, fmt.Errorf("未找到数据文件")
	}

	// 按时间排序，取最新的
	latestFile := ""
	for _, file := range files {
		if latestFile == "" || file > latestFile {
			latestFile = file
		}
	}

	// 加载文件
	return LoadMarketDataFromGob(latestFile)
}

// 从GoB文件加载
func LoadMarketDataFromGob(filename string) (*MarketDataFile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data MarketDataFile
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func PrintAllItems(mi *map[string]map[int]MarketItem) {
	for n, md := range *mi {
		for l, mi := range md {
			fmt.Printf("物品名称: %s, 等级: %d, 左一: %d, 右一: %d\n", n, l, mi.Sell, mi.Buy)
		}
	}
}

type ResponseJson struct {
	MarketData map[string]map[string]map[string]int
	Timestamp  int
}

// 市场里物品的结构
type MarketItem struct {
	Level int   // 等级
	Sell  int64 // 左一
	Buy   int64 // 右一
}

type PriceInfo struct {
}

// 完整的存储结构
type MarketDataFile struct {
	SaveTime      time.Time                     // 保存时间
	DataTimestamp int                           // 数据时间戳
	MarketItems   map[string]map[int]MarketItem // 物品结构
}
