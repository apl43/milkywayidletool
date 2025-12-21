package config

import (
	"milkywayidletool/models"
)

// 玩家信息配置
var P1 = models.PlayerInfo{
	Name: "apl43",
	Exp: models.Experience{
		AlchemyExperience: 382_493_094,
	},
	Equipment: models.EquipmentList{
		Body: []models.EquipmentInfo{
			// todo: add body equipment info
		},
		Tools: []models.EquipmentInfo{
			// todo: add tools equipment info
		},
		Accessories: []models.EquipmentInfo{
			// todo: add accessories equipment info
		},
	},
}
