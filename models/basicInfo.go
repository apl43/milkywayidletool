package models

// 玩家基本信息
type PlayerInfo struct {
	Name	  string        // 玩家名称
	Exp       Experience    // 经验
	Equipment EquipmentList // 装备列表
}

// 配装列表
type EquipmentList struct {
	Body        []EquipmentInfo // 装备
	Tools       []EquipmentInfo // 工具
	Accessories []EquipmentInfo // 配饰
}

// 装备信息
type EquipmentInfo struct {
	ID     int              // 装备ID
	Slot   string           // 穿戴位置
	Level  int              // 装备等级
	Effect EquipmentEffects // 装备效果
}

// buff加成
type EquipmentEffects struct {
	EfficiencyBonus    BuffRangeAndEffect // 效率加成
	RareDiscoveryBonus BuffRangeAndEffect // 稀有发现加成
	SuccessRateBonus   BuffRangeAndEffect // 成功率加成
	ActionSpeedBonus   BuffRangeAndEffect // 行动速度加成
	ExperienceBonus    BuffRangeAndEffect // 经验加成
}

// buff加成范围和效果
type BuffRangeAndEffect struct {
	Range  string  // 加成范围描述
	Effect float64 // 加成效果
}

// 各技能等级
type Experience struct {
	AlchemyExperience float64
}

// 社区buff加成
type CommunityBuffs struct {
	ExperienceBonus        CommunityBuffLevels // 经验加成
	GatheringQuantityBonus CommunityBuffLevels // 采集数量加成
	ProductionEfficiency   CommunityBuffLevels // 生产效率加成
	EnhancementSuccessRate CommunityBuffLevels // 强化成功率加成
	BattleDropQuantity     CommunityBuffLevels // 战斗掉落数量加成
}

// 社区buff等级和效果
type CommunityBuffLevels struct {
	Level  int     // buff等级
	Effect float64 // buff效果
}
