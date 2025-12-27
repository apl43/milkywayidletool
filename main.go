package main

import (
	"fmt"
	"milkywayidletool/database/market"
	// "milkywayidletool/tool"
)

func main() {
	fmt.Println("Hello, Milky Way Idle Tool!")
	fmt.Println()
	// tool.CalculateAllRevenue()

	// // 获取市场数据
	// market.GetMarketData()

	// 加载最新市场文件
	latestFile, err := market.LoadLatestMarketData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("保存时间: %s\n", latestFile.SaveTime)
	fmt.Printf("数据时间戳: %d\n", latestFile.DataTimestamp)
	// market.PrintAllItems(&latestFile.MarketItems)
}
