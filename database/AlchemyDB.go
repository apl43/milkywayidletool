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
		RecommendedLevel: 40,
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
		RecommendedLevel: 50,
	},
	"复活": {
		Name:        "复活",
		Quantity:    1,
		GoldCost:    2000, 
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "复活", Quantity: 1, Chance: 0.125},
			{Name: "疯狂", Quantity: 1, Chance: 0.125},
			{Name: "无敌", Quantity: 1, Chance: 0.125},
			{Name: "速度光环", Quantity: 1, Chance: 0.125},
			{Name: "守护光环", Quantity: 1, Chance: 0.125},
			{Name: "物理光环", Quantity: 1, Chance: 0.125},
			{Name: "暴击光环", Quantity: 1, Chance: 0.125},
			{Name: "元素光环", Quantity: 1, Chance: 0.125},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
		RecommendedLevel: 60,
	},
	"疯狂": {
		Name:        "疯狂",
		Quantity:    1,
		GoldCost:    2000, 
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "复活", Quantity: 1, Chance: 0.125},
			{Name: "疯狂", Quantity: 1, Chance: 0.125},
			{Name: "无敌", Quantity: 1, Chance: 0.125},
			{Name: "速度光环", Quantity: 1, Chance: 0.125},
			{Name: "守护光环", Quantity: 1, Chance: 0.125},
			{Name: "物理光环", Quantity: 1, Chance: 0.125},
			{Name: "暴击光环", Quantity: 1, Chance: 0.125},
			{Name: "元素光环", Quantity: 1, Chance: 0.125},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
		RecommendedLevel: 60,
	},
	"无敌": {
		Name:        "无敌",
		Quantity:    1,
		GoldCost:    2000, 
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "复活", Quantity: 1, Chance: 0.125},
			{Name: "疯狂", Quantity: 1, Chance: 0.125},
			{Name: "无敌", Quantity: 1, Chance: 0.125},
			{Name: "速度光环", Quantity: 1, Chance: 0.125},
			{Name: "守护光环", Quantity: 1, Chance: 0.125},
			{Name: "物理光环", Quantity: 1, Chance: 0.125},
			{Name: "暴击光环", Quantity: 1, Chance: 0.125},
			{Name: "元素光环", Quantity: 1, Chance: 0.125},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
		RecommendedLevel: 60,
	},
	"速度光环": {
		Name:        "速度光环",
		Quantity:    1,
		GoldCost:    2000, 
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "复活", Quantity: 1, Chance: 0.125},
			{Name: "疯狂", Quantity: 1, Chance: 0.125},
			{Name: "无敌", Quantity: 1, Chance: 0.125},
			{Name: "速度光环", Quantity: 1, Chance: 0.125},
			{Name: "守护光环", Quantity: 1, Chance: 0.125},
			{Name: "物理光环", Quantity: 1, Chance: 0.125},
			{Name: "暴击光环", Quantity: 1, Chance: 0.125},
			{Name: "元素光环", Quantity: 1, Chance: 0.125},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
		RecommendedLevel: 40,
	},
	"守护光环": {
		Name:        "守护光环",
		Quantity:    1,
		GoldCost:    2000, 
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "复活", Quantity: 1, Chance: 0.125},
			{Name: "疯狂", Quantity: 1, Chance: 0.125},
			{Name: "无敌", Quantity: 1, Chance: 0.125},
			{Name: "速度光环", Quantity: 1, Chance: 0.125},
			{Name: "守护光环", Quantity: 1, Chance: 0.125},
			{Name: "物理光环", Quantity: 1, Chance: 0.125},
			{Name: "暴击光环", Quantity: 1, Chance: 0.125},
			{Name: "元素光环", Quantity: 1, Chance: 0.125},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
		RecommendedLevel: 40,
	},
	"物理光环": {
		Name:        "物理光环",
		Quantity:    1,
		GoldCost:    2000, 
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "复活", Quantity: 1, Chance: 0.125},
			{Name: "疯狂", Quantity: 1, Chance: 0.125},
			{Name: "无敌", Quantity: 1, Chance: 0.125},
			{Name: "速度光环", Quantity: 1, Chance: 0.125},
			{Name: "守护光环", Quantity: 1, Chance: 0.125},
			{Name: "物理光环", Quantity: 1, Chance: 0.125},
			{Name: "暴击光环", Quantity: 1, Chance: 0.125},
			{Name: "元素光环", Quantity: 1, Chance: 0.125},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
		RecommendedLevel: 40,
	},
	"暴击光环": {
		Name:        "暴击光环",
		Quantity:    1,
		GoldCost:    2000, 
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "复活", Quantity: 1, Chance: 0.125},
			{Name: "疯狂", Quantity: 1, Chance: 0.125},
			{Name: "无敌", Quantity: 1, Chance: 0.125},
			{Name: "速度光环", Quantity: 1, Chance: 0.125},
			{Name: "守护光环", Quantity: 1, Chance: 0.125},
			{Name: "物理光环", Quantity: 1, Chance: 0.125},
			{Name: "暴击光环", Quantity: 1, Chance: 0.125},
			{Name: "元素光环", Quantity: 1, Chance: 0.125},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
		RecommendedLevel: 40,
	},
	"元素光环": {
		Name:        "元素光环",
		Quantity:    1,
		GoldCost:    2000, 
		ActionTime:  20,
		SuccessRate: 0.5,
		MainProducts: []models.ProductProbability{
			{Name: "复活", Quantity: 1, Chance: 0.125},
			{Name: "疯狂", Quantity: 1, Chance: 0.125},
			{Name: "无敌", Quantity: 1, Chance: 0.125},
			{Name: "速度光环", Quantity: 1, Chance: 0.125},
			{Name: "守护光环", Quantity: 1, Chance: 0.125},
			{Name: "物理光环", Quantity: 1, Chance: 0.125},
			{Name: "暴击光环", Quantity: 1, Chance: 0.125},
			{Name: "元素光环", Quantity: 1, Chance: 0.125},
		},
		ExtraOutput: models.ExtraOutput{
			EssenceName:    "炼金精华",
			EssenceChance:  0, // 待修改
			RareDropName:   "中工匠匣",
			RareDropChance: 0, // 待修改
			ExperienceGain: 0, // 待修改
		},
		RecommendedLevel: 40,
	},
}
