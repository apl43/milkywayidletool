package database

import "milkywayidletool/models"

// 存放所有转化配方
var ConversionRecipeMap = map[string]models.ConversionRecipe{
	"点金催化剂": {
		Name:        "点击催化剂",
		Quantity:    1,
		GoldCost:    100,
		ActionTime:  20,
		SuccessRate: 0.45,
		MainProducts: []models.ProductProbability{
			{Name: "点金催化剂", Quantity: 1, Chance: 0.36},
			{Name: "分解催化剂", Quantity: 1, Chance: 0.33},
			{Name: "转化催化剂", Quantity: 1, Chance: 0.27},
			{Name: "至高催化剂", Quantity: 1, Chance: 0.04},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
	},
	"分解催化剂": {
		Name:        "分解催化剂",
		Quantity:    1,
		GoldCost:    100,
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "点金催化剂", Quantity: 1, Chance: 0.36},
			{Name: "分解催化剂", Quantity: 1, Chance: 0.33},
			{Name: "转化催化剂", Quantity: 1, Chance: 0.27},
			{Name: "至高催化剂", Quantity: 1, Chance: 0.04},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
	},
}
