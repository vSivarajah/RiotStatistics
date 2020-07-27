package api

type TeamStatsDTO struct {
	BaronKills           int    `json:"baronKills"`
	DominionVictoryScore int64  `json:"dominionVictoryScore"`
	DragonKills          int    `json:"dragonKills"`
	FirstBaron           bool   `json:"firstBaron"`
	FirstBlood           bool   `json:"firstBlood"`
	FirstDragon          bool   `json:"firstDragon"`
	FirstInhibitor       bool   `json:"firstInhibitor"`
	FirstRiftHerald      bool   `json:"firstRiftHerald"`
	FirstTower           bool   `json:"firstTower"`
	InhibitorKills       int    `json:"inhibitorKills"`
	RiftHeraldKills      int    `json:"riftHeraldKills"`
	TeamID               int    `json:"teamId"`
	TowerKills           int    `json:"towerKills"`
	VilemawKills         int    `json:"vilemawKills"`
	Win                  string `json:"win"`
}
