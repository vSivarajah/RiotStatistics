package api

type TeamStatsDTO struct {
	BaronKills           int    `json:"baronKills" bson:"baronKills"`
	DominionVictoryScore int64  `json:"dominionVictoryScore" bson:"dominionVictoryScore"`
	DragonKills          int    `json:"dragonKills" bson:"dragonKills"`
	FirstBaron           bool   `json:"firstBaron" bson:"firstBaron"`
	FirstBlood           bool   `json:"firstBlood" bson:"firstBlood"`
	FirstDragon          bool   `json:"firstDragon" bson:"firstDragon"`
	FirstInhibitor       bool   `json:"firstInhibitor" bson:"firstInhibitor"`
	FirstRiftHerald      bool   `json:"firstRiftHerald" bson:"firstRiftHerald"`
	FirstTower           bool   `json:"firstTower" bson:"firstTower"`
	InhibitorKills       int    `json:"inhibitorKills" bson:"inhibitorKills"`
	RiftHeraldKills      int    `json:"riftHeraldKills" bson:"riftHeraldKills"`
	TeamID               int    `json:"teamId" bson:"teamId"`
	TowerKills           int    `json:"towerKills" bson:"towerKills"`
	VilemawKills         int    `json:"vilemawKills" bson:"vilemawKills"`
	Win                  string `json:"win" bson:"win"`
}
