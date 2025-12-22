package models

// 转化models
// 转化操作结构
type Conversionoperation struct {
	RecipeName     string   // 转化项目名称
	InputName      string   // 原料名称
	InputQuantity  int      // 原料数量
	Catalyst       string   // 选择催化剂
	RetainProducts []string // 保留产物
}

// 转化配方结构
type ConversionRecipe struct {
	Name       string // 消耗物品名称
	Quantity   int    // 消耗物品数量
	GoldCost   int    // 消耗金币
	ActionTime int    // 行动时间(单位: 秒)

	SuccessRate float64 // 成功率

	MainProducts []ProductProbability // 主要产物
	ExtraOutput  ExtraOutput          // 额外产出
}

// 主要产物
type ProductProbability struct {
	Name     string  // 产物名称
	Quantity int     // 产物数量
	Chance   float64 // 获得概率
}

// 额外产出
type ExtraOutput struct {
	EssenceName   string  // 精华名称
	EssenceChance float64 // 精华掉落概率

	RareDropName   string  // 稀有掉落名称
	RareDropChance float64 // 稀有掉落概率

	ExperienceGain float64 // 经验
}

// 炼金加成
type AlchemyBonus struct {
	EfficiencyBonus    float64 // 效率加成
	SuccessRateBonus   float64 // 炼金成功率加成
	ActionSpeedBonus   float64 // 行动速度加成
	ExperienceBonus    float64 // 经验加成
	RareDiscoveryBonus float64 // 稀有发现加成
}
