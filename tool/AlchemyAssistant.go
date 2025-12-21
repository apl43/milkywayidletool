package tool

import (
	"fmt"
	"github.com/apl43/milkywayidletool/model/alchemymodel"
)

// // 获取炼金经验
// func GetAlchemyExperience() float64 {

// }

// 计算转化产出
func CalculateTheTransformationOutput(i *alchemymodel.AlchemyTransformationConfig) {
	i.Cost = alchemymodel.AlchemyCost{ItemID: 1, Quantity: 10, GoldCost: 100, CatalystID: 0}
	fmt.Println("111", i.Cost)
}
