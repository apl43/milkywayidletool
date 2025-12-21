package tool

import (
	"fmt"
	"milkywayidletool/models"
)

// // 获取炼金经验
// func GetAlchemyExperience() float64 {

// }

// // 计算转化产出
// func CalculateTheTransformationOutput(i *models.AlchemyTransformationConfig) {
// 	i.Cost = models.AlchemyCost{ItemID: 1, Quantity: 10, GoldCost: 100, CatalystID: 0}
// 	fmt.Println("111", i.Cost)
// }


func Te() {
	fmt.Println("测试")
	var a models.AlchemyTransformationConfig
	a = models.AlchemyTransformationConfig{
		Cost: models.AlchemyCost{ItemID: 1, Quantity: 10, GoldCost: 100, CatalystID: 0},
	}
	fmt.Println(a)
}