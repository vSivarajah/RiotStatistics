package api

type ParticipantIdentities struct {
	ParticipantID int    `json:"participantId" bson:"participantId"`
	Player        Player `json:"player" bson:"player"`
}

type Player struct {
	MatchHistoryURI string `json:"matchHistoryUri" bson:"matchHistoryUri"`
	ProfileIcon     int    `json:"profileIcon" bson:"profileIcon"`
	SummonerID      int64  `json:"summonerId" bson:"summonerId"`
	SummonerName    string `json:"summonerName" bson:"summonerName"`
}

type ParticipantDTO struct {
	ChampionID                int                 `json:"championId"`
	HighestAchievedSeasonTier string              `json:"highestAchievedSeasonTier"`
	Masteries                 []Mastery           `json:"masteries"`
	ParticipantID             int                 `json:"participantId"`
	Runes                     []Rune              `json:"runes"`
	Spell1ID                  int                 `json:"spell1Id"`
	Spell2ID                  int                 `json:"spell2Id"`
	Stats                     ParticipantStats    `json:"stats"`
	TeamID                    int                 `json:"teamId"`
	Timeline                  ParticipantTimeline `json:"timeline"`
}

type ParticipantTimeline struct {
	AncientGolemAssistsPerMinCounts ParticipantTimelineData `json:"ancientGolemAssistsPerMinCounts"`
	AncientGolemKillsPerMinCounts   ParticipantTimelineData `json:"ancientGolemKillsPerMinCounts"`
	AssistedLaneDeathsPerMinDeltas  ParticipantTimelineData `json:"assistedLaneDeathsPerMinDeltas"`
	AssistedLaneKillsPerMinDeltas   ParticipantTimelineData `json:"assistedLaneKillsPerMinDeltas"`
	BaronAssistsPerMinCounts        ParticipantTimelineData `json:"baronAssistsPerMinCounts"`
	BaronKillsPerMinCounts          ParticipantTimelineData `json:"baronKillsPerMinCounts"`
	CreepsPerMinDeltas              ParticipantTimelineData `json:"creepsPerMinDeltas"`
	CsDiffPerMinDeltas              ParticipantTimelineData `json:"csDiffPerMinDeltas"`
	DamageTakenDiffPerMinDeltas     ParticipantTimelineData `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas         ParticipantTimelineData `json:"damageTakenPerMinDeltas"`
	DragonAssistsPerMinCounts       ParticipantTimelineData `json:"dragonAssistsPerMinCounts"`
	DragonKillsPerMinCounts         ParticipantTimelineData `json:"dragonKillsPerMinCounts"`
	ElderLizardAssistsPerMinCounts  ParticipantTimelineData `json:"elderLizardAssistsPerMinCounts"`
	ElderLizardKillsPerMinCounts    ParticipantTimelineData `json:"elderLizardKillsPerMinCounts"`
	GoldPerMinDeltas                ParticipantTimelineData `json:"goldPerMinDeltas"`
	InhibitorAssistsPerMinCounts    ParticipantTimelineData `json:"inhibitorAssistsPerMinCounts"`
	InhibitorKillsPerMinCounts      ParticipantTimelineData `json:"inhibitorKillsPerMinCounts"`
	Lane                            string                  `json:"lane"`
	Role                            string                  `json:"role"`
	TowerAssistsPerMinCounts        ParticipantTimelineData `json:"towerAssistsPerMinCounts"`
	TowerKillsPerMinCounts          ParticipantTimelineData `json:"towerKillsPerMinCounts"`
	TowerKillsPerMinDeltas          ParticipantTimelineData `json:"towerKillsPerMinDeltas"`
	VilemawAssistsPerMinCounts      ParticipantTimelineData `json:"vilemawAssistsPerMinCounts"`
	VilemawKillsPerMinCounts        ParticipantTimelineData `json:"vilemawKillsPerMinCounts"`
	WardsPerMinDeltas               ParticipantTimelineData `json:"wardsPerMinDeltas"`
	XpDiffPerMinDeltas              ParticipantTimelineData `json:"xpDiffPerMinDeltas"`
	XpPerMinDeltas                  ParticipantTimelineData `json:"xpPerMinDeltas"`
}

type ParticipantTimelineData struct {
	TenToTwenty    float64 `json:"tenToTwenty" bson:"tenToTwenty"`
	ThirtyToTen    float64 `json:"thirtyToTen" bson:"thirtyToTen"`
	TwentyToThirty float64 `json:"twentyToThirty" bson:"twentyToThirty"`
	ZeroToTen      float64 `json:"zeroToTen" bson:"zeroToTen"`
}
type ParticipantStats struct {
	Assists                         int64 `json:"assists" bson:"assists"`
	ChampLevel                      int64 `json:"champLevel" bson:"champLevel"`
	CombatPlayerScore               int64 `json:"combatPlayerScore" bson:"combatPlayerScore"`
	Deaths                          int64 `json:"deaths" bson:"deaths"`
	DoubleKills                     int64 `json:"doubleKills" bson:"doubleKills"`
	FirstBloodAssist                bool  `json:"firstBloodAssist" bson:"firstBloodAssist"`
	FirstBloodKill                  bool  `json:"firstBloodKill" bson:"firstBloodKill"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist" bson:"firstInhibitorAssist"`
	FirstInhibitorKill              bool  `json:"firstInhibitorKill" bson:"firstInhibitorKill"`
	FirstTowerAssist                bool  `json:"firstTowerAssist" bson:"firstTowerAssist"`
	FirstTowerKill                  bool  `json:"firstTowerKill" bson:"firstTowerKill"`
	GoldEarned                      int64 `json:"goldEarned" bson:"goldEarned"`
	GoldSpent                       int64 `json:"goldSpent" bson:"goldSpent"`
	InhibitorKills                  int64 `json:"inhibitorKills" bson:"inhibitorKills"`
	Item0                           int64 `json:"item0" bson:"item0"`
	Item1                           int64 `json:"item1" bson:"item1"`
	Item2                           int64 `json:"item2" bson:"item2"`
	Item3                           int64 `json:"item3" bson:"item3"`
	Item4                           int64 `json:"item4" bson:"item4"`
	Item5                           int64 `json:"item5" bson:"item5"`
	Item6                           int64 `json:"item6" bson:"item6"`
	KillingSprees                   int64 `json:"killingSprees" bson:"killingSprees"`
	Kills                           int64 `json:"kills" bson:"kills"`
	LargestCriticalStrike           int64 `json:"largestCriticalStrike" bson:"largestCriticalStrike"`
	LargestKillingSpree             int64 `json:"largestKillingSpree" bson:"largestKillingSpree"`
	LargestMultiKill                int64 `json:"largestMultiKill" bson:"largestMultiKill"`
	MagicDamageDealt                int64 `json:"magicDamageDealt" bson:"magicDamageDealt"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions" bson:"magicDamageDealtToChampions"`
	MagicDamageTaken                int64 `json:"magicDamageTaken" bson:"magicDamageTaken"`
	TotalMinionsKilled              int64 `json:"totalMinionsKilled" bson:"totalMinionsKilled"`
	NeutralMinionsKilled            int64 `json:"neutralMinionsKilled" bson:"neutralMinionsKilled"`
	NeutralMinionsKilledEnemyJungle int64 `json:"neutralMinionsKilledEnemyJungle" bson:"neutralMinionsKilledEnemyJungle"`
	NeutralMinionsKilledTeamJungle  int64 `json:"neutralMinionsKilledTeamJungle" bson:"neutralMinionsKilledTeamJungle"`
	NodeCapture                     int64 `json:"nodeCapture" bson:"nodeCapture"`
	NodeCaptureAssist               int64 `json:"nodeCaptureAssist" bson:"nodeCaptureAssist"`
	NodeNeutralize                  int64 `json:"nodeNeutralize" bson:"nodeNeutralize"`
	ObjectivePlayerScore            int64 `json:"objectivePlayerScore" bson:"objectivePlayerScore"`
	PentaKills                      int64 `json:"pentaKills" bson:"pentaKills"`
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt" bson:"physicalDamageDealt"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions" bson:"physicalDamageDealtToChampions"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken" bson:"physicalDamageTaken"`
	QuadraKills                     int64 `json:"quadraKills" bson:"quadraKills"`
	SightWardsBoughtInGame          int64 `json:"sightWardsBoughtInGame" bson:"sightWardsBoughtInGame"`
	TeamObjective                   int64 `json:"teamObjective" bson:"teamObjective"`
	TotalDamageDealt                int64 `json:"totalDamageDealt" bson:"totalDamageDealt"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions" bson:"totalDamageDealtToChampions"`
	TotalDamageTaken                int64 `json:"totalDamageTaken" bson:"totalDamageTaken"`
	TotalHeal                       int64 `json:"totalHeal" bson:"totalHeal"`
	TotalPlayerScore                int64 `json:"totalPlayerScore" bson:"totalPlayerScore"`
	TotalScoreRank                  int64 `json:"totalScoreRank" bson:"totalScoreRank"`
	TotalTimeCrowdControlDealt      int64 `json:"totalTimeCrowdControlDealt" bson:"totalTimeCrowdControlDealt"`
	TotalUnitsHealed                int64 `json:"totalUnitsHealed" bson:"totalUnitsHealed"`
	TowerKills                      int64 `json:"towerKills" bson:"towerKills"`
	TripleKills                     int64 `json:"tripleKills" bson:"tripleKills"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt" bson:"trueDamageDealt"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions" bson:"trueDamageDealtToChampions"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken" bson:"trueDamageTaken"`
	UnrealKills                     int64 `json:"unrealKills" bson:"unrealKills"`
	VisionWardsBoughtInGame         int64 `json:"visionWardsBoughtInGame" bson:"visionWardsBoughtInGame"`
	WardsKilled                     int64 `json:"wardsKilled" bson:"wardsKilled"`
	WardsPlaced                     int64 `json:"wardsPlaced" bson:"wardsPlaced"`
	Win                             bool  `json:"win" bson:"win"`
}
