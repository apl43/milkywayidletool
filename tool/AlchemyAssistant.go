package tool

import (
	"fmt"
	"math"
	"math/rand"
	"milkywayidletool/database"
	"milkywayidletool/models"
	"time"
)

// // 打印玩家信息
// func PrintPlayerInfo(p models.PlayerInfo) {
// 	fmt.Printf("===== 打印玩家 %s 信息 =====\n", p.Name)
// 	fmt.Println("===== 各技能经验等级 =====")
// 	for i, x := range p.Exp {
// 		fmt.Printf("%s 经验: %d\n", i, x)
// 	}
// }

// 炼金项目列表
// var TransformationList = map[string]models.ConversionSolution{
// 	"点金催化剂": {
// 		Catalyst: 0,
// 		Product:  []string{"转化催化剂", "至高催化剂"},
// 	},
// 	"分解催化剂": {
// 		Catalyst: 0,
// 		Product:  []string{"转化催化剂", "至高催化剂"},
// 	},
// }

// 炼金转化项目列表
var ConversionOperationList = []models.ConversionOperation{
	{RecipeName: "催化剂项目-买黄催", InputName: "点金催化剂", InputQuantity: 20000, Catalyst: "", RetainProducts: []string{"转化催化剂", "至高催化剂"}},
	{RecipeName: "催化剂项目-买绿催", InputName: "分解催化剂", InputQuantity: 20000, Catalyst: "", RetainProducts: []string{"转化催化剂", "至高催化剂"}},
}

// 玩家炼金等级
var playAlchemyLevel = 136

// 计算全部项目收益
func CalculateAllRevenue() {
	rand.Seed(time.Now().UnixNano())
	// 转化

	// 1. 缓存社区buff, 装备信息, 房屋, 成就加成, mo卡
	// 2. 获取转化项目列表
	// 3. 枚举每一个转化项目，计算转化基础信息和加成
	// 4. 获取结果，计算收益

	// 1.
	// 基础加成
	var baseBonus = models.AlchemyBonus{
		EfficiencyBonus:    0.197,
		SuccessRateBonus:   0,
		ActionSpeedBonus:   0,
		ExperienceBonus:    0.295,
		RareDiscoveryBonus: 0,
	}

	// 先计算基础加成, 暴饮之囊 +10 浓度 +0.129
	// 效率
	baseBonus.EfficiencyBonus += 0.129        // 炼金师上衣 +10
	baseBonus.EfficiencyBonus += 0.129        // 炼金师下装 +10
	baseBonus.EfficiencyBonus += 0.129        // 附魔手套 +10
	baseBonus.EfficiencyBonus += 0.075        // 实验室 +5
	baseBonus.EfficiencyBonus += 0.02         // 成就 熟练者
	baseBonus.EfficiencyBonus += 0.10 * 1.129 // 效率茶
	// 炼金成功率
	baseBonus.SuccessRateBonus += 0.05 * 1.129 // 催化茶
	// 行动速度
	baseBonus.ActionSpeedBonus += 1.3545 // 星空蒸馏器 +10
	baseBonus.ActionSpeedBonus += 0.064  // 速度项链 +5
	// 经验
	baseBonus.ExperienceBonus += 0.0516       // 炼金师下装 +10
	baseBonus.ExperienceBonus += 0.0516       // 星空蒸馏器 +10
	baseBonus.ExperienceBonus += 0.08         // 专家炼金护符 +5
	baseBonus.ExperienceBonus += 0.0105       // 一堆房屋
	baseBonus.ExperienceBonus += 0.02         // 成就 新手
	baseBonus.ExperienceBonus += 0.12 * 1.129 // 经验茶
	baseBonus.ExperienceBonus += 0.05         // mo卡
	// 稀有发现
	baseBonus.RareDiscoveryBonus += 0.1935 // 炼金师上衣 +10
	baseBonus.RareDiscoveryBonus += 0.1935 // 星空蒸馏器 +10
	baseBonus.RareDiscoveryBonus += 0.08   // 稀有发现耳环 +0
	baseBonus.RareDiscoveryBonus += 0.08   // 稀有发现戒指 +0
	baseBonus.RareDiscoveryBonus += 0.042  // 一堆房屋

	// fmt.Printf("%+v", baseBonus)

	for _, item := range ConversionOperationList {
		cost, revenue, time, err := CalcConversionItemProcess(item, baseBonus)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("消耗: %+v, 收入: %+v, 耗时: %.2fs\n", cost, revenue, time)
		fmt.Println()
	}
}

// 物品数量结构
type ItemQuantity struct {
	Item     string // 物品
	Quantity int64  // 数量
}

// 计算转化项目过程
// 传入(单个转化项目的转化操作结构, 加成信息)
// 输出(消耗, 产出, 耗时, err)
// 1. 原料列表(实时更新的, 产物可能也是原料), 消耗列表(金币, 催化剂...), 产出列表
// 2. 枚举原料列表里的物品
// 2.1 计算成功率(受催化剂影响) & 效率(受推荐等级影响) & 行动时间
// 2.2 计算物品全部转完的过程
// 2.2.1 计算效率加成下的真实行动次数c
// 2.2.2 计算c次行动的产物
func CalcConversionItemProcess(co models.ConversionOperation, ab models.AlchemyBonus) (map[string]int64, map[string]int64, float64, error) {
	fmt.Printf("项目[%s], 消耗%d个%s, 保留%s\n", co.RecipeName, co.InputQuantity, co.InputName, co.RetainProducts)

	// 1. 原料列表(实时更新的, 产物可能也是原料), 消耗列表(金币, 催化剂...), 产出列表
	rawMaterial := map[string]int64{co.InputName: int64(co.InputQuantity)}
	cost := make(map[string]int64)
	revenue := make(map[string]int64)
	retainProductsSet := make(map[string]bool)
	for _, x := range co.RetainProducts {
		retainProductsSet[x] = true
	}
	var time float64

	// 2. 枚举原料列表里的物品
	for len(rawMaterial) > 0 {
		// fmt.Printf("原料: %+v\n", rawMaterial)
		var n string
		var q int64
		for n, q = range rawMaterial {
			delete(rawMaterial, n)
			break
		}
		// 2.1 计算成功率(受催化剂影响) & 效率(受推荐等级影响) & 行动时间
		successRate := database.ConversionRecipeMap[n].SuccessRate * (1 + ab.SuccessRateBonus)
		efficiency := ab.EfficiencyBonus
		diff := playAlchemyLevel - database.ConversionRecipeMap[n].RecommendedLevel
		if diff > 0 {
			efficiency += float64(diff) * 0.01
		}
		actionTime := database.ConversionRecipeMap[n].ActionTime / (1 + ab.ActionSpeedBonus)
		// 2.2 计算物品全部转完的过程
		for q > 0 {
			// fmt.Printf("物品: %s, 数量: %d\n", n, q)
			// 2.2.1 计算效率加成下的真实行动次数c
			c := int64(1 + math.Floor(efficiency))
			if rand.Float64() < math.Mod(efficiency, 1) {
				c++
			}
			if q > c {
				q -= c 
			} else {
				q = 0
			}
			// 2.2.2 计算c次行动的产物
			resCost, resRevenue, err := CalcConversionProcess(co.InputName, co.Catalyst, successRate)
			if err != nil {
				return nil, nil, 0, err
			}
			for _, rc := range resCost {
				cost[rc.Item] += rc.Quantity
			}
			for _, rr := range resRevenue {
				if retainProductsSet[rr.Item] {
					revenue[rr.Item] += rr.Quantity
				} else {
					rawMaterial[rr.Item] += rr.Quantity
				}
			}
			time += actionTime
		}

	}
	return cost, revenue, time, nil
}

// 计算转化过程的函数
// 传入(转化物品, 催化剂, 成功率)
// 输出(消耗, 产出, err)
func CalcConversionProcess(name string, catalyst string, rate float64) ([]ItemQuantity, []ItemQuantity, error) {
	var cost []ItemQuantity
	var revenue []ItemQuantity
	cost = append(cost, ItemQuantity{Item: "金币", Quantity: database.ConversionRecipeMap[name].GoldCost})

	// 计算成功率, 失败只消耗金币
	if rand.Float64() > rate {
		return cost, revenue, nil
	}
	// 成功先消耗催化剂
	if catalyst != "" {
		cost = append(cost, ItemQuantity{Item: catalyst, Quantity: 1})
	}

	// 计算产物
	// 累积概率法
	// 1. 保证所有产物概率之和为1
	// 2. 随机一个[0, 1)的浮点数f
	// 3. 依次累加所以产物的概率, 第一个累加后大于f的产物就是抽中的
	var sum float64
	for _, pp := range database.ConversionRecipeMap[name].MainProducts {
		sum += pp.Chance
	}
	// 1. 保证所有产物概率之和为1
	if sum != float64(1) {
		return nil, nil, fmt.Errorf("[概率之和不为1, 请检查录入数据] Name: %s, MainProducts: %+v", name, database.ConversionRecipeMap[name].MainProducts)
	}
	// 2. 随机一个[0, 1)的浮点数f
	f := rand.Float64()
	// 3. 依次累加所以产物的概率, 第一个累加后大于f的产物就是抽中的
	sum = 0
	var flag bool
	for _, pp := range database.ConversionRecipeMap[name].MainProducts {
		sum += pp.Chance
		if sum > f {
			revenue = append(revenue, ItemQuantity{Item: pp.Name, Quantity: int64(pp.Quantity)})
			flag = true
			break
		}
	}
	if !flag {
		return nil, nil, fmt.Errorf("[概率总和小于随机数] f: %.6f, sum: %.6f", f, sum)
	}
	return cost, revenue, nil
}
