package model

// 炼金转化配置
type AlchemyTransformationConfig struct {
	Cost        AlchemyCost // 消耗
	SuccessRate float64     // 成功率
	Output      float64     // 产出
	ActionTime  float64     // 行动时间(单位:秒)
}

// 炼金消耗
type AlchemyCost struct {
	ItemID     int // 物品ID
	Quantity   int // 物品数量
	GoldCost   int // 消耗金币数量
	CatalystID int // 催化剂ID  0: 无｜1: 黄催｜2: 绿催｜3: 紫催｜4: 红催
}

// 炼金产出
type AlchemyOutput struct {
	OutputList       map[int]AlchemyOutputItem // 产出列表
	EssenceOutput    map[int]AlchemyOutputItem // 精华
	RareOutput       map[int]AlchemyOutputItem // 稀有
	ExperienceOutput float64                   // 经验
}

// 炼金产出物品
type AlchemyOutputItem struct {
	Id          int     // 物品ID
	Probability float64 // 产出概率
	Quantity    int     // 产出数量
}

// 炼金加成
type AlchemyBonus struct {
	EfficiencyBonus    float64 // 效率加成
	SuccessRateBonus   float64 // 炼金成功率加成
	ActionSpeedBonus   float64 // 行动速度加成
	ExperienceBonus    float64 // 经验加成
	RareDiscoveryBonus float64 // 稀有发现加成
}
