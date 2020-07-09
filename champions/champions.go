package champions

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ChampionData struct {
	Type      string              `json:"type"`
	Format    string              `json:"format"`
	Version   string              `json:"version"`
	Champions map[string]Champion `json:"data"`
	Keys      map[string]string   `json:"keys"`
}

type Champion struct {
	Version string          `json:"version"`
	Id      string          `json:"id"`
	Key     string          `json:"key"`
	Name    string          `json:"name"`
	Title   string          `json:"title"`
	Blurb   string          `json:"blurb"`
	ParType string          `json:"partype"`
	Passive ChampionPassive `json:"passive"`
	Info    ChampionInfo    `json:"info"`
	Tags    []string        `json:"tags"`
}
type ChampionInfo struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty" gorethink:"difficulty"`
}

type ChampionStats struct {
	HP                   float32 `json:"hp"`
	HPPerLevel           float32 `json:"hpperlevel"`
	MP                   float32 `json:"mp"`
	MPPerLevel           float32 `json:"mpperlevel"`
	MoveSpeed            float32 `json:"movespeed"`
	Armor                float32 `json:"armor"`
	ArmorPerLevel        float32 `json:"armorperlevel"`
	SpellBlock           float32 `json:"spellblock"`
	SpellBlockPerLevel   float32 `json:"spellblockperlevel"`
	AttackRange          float32 `json:"attackrange"`
	HPRegen              float32 `json:"hpregen"`
	HPRegenPerLevel      float32 `json:"hpregenperlevel"`
	MPRegen              float32 `json:"mpregen"`
	MPRegenPerLevel      float32 `json:"mpregenperlevel"`
	Crit                 float32 `json:"crit"`
	CritPerLevel         float32 `json:"critperlevel"`
	AttackDamage         float32 `json:"attackdamage"`
	AttackDamagePerLevel float32 `json:"attackdamageperlevel"`
	AttackSpeedOffset    float32 `json:"attackspeedoffset"`
	AttackSpeedPerLevel  float32 `json:"attackspeedperlevel"`
}

type ChampionPassive struct {
	Description          string `json:"description"`
	SanitizedDescription string `json:"sanitizedDescription"`
	Name                 string `json:"name"`
	Image                Image  `json:"image"`
}
type Image struct {
	Full    string `json:"full"`
	Sprite  string `json:"sprite"`
	Group   string `json:"group"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	W       int    `json:"w"`
	H       int    `json:"h"`
	Encoded string
}

func GetChampions() ChampionData {
	plan, _ := ioutil.ReadFile("champions.json")
	data := ChampionData{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
