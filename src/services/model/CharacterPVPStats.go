package model

type CharacterPVPStats struct {
	Ranking int `json:"ranking"`
	Rating int `json:"rating"`
	Name string `json:"name"`
	RealmID int `json:"realmId"`
	RealmName string `json:"realmName"`
	RealmSlug string `json:"realmSlug"`
	RaceID int `json:"raceId"`
	ClassID int `json:"classId"`
	SpecID int `json:"specId"`
	FactionID int `json:"factionId"`
	GenderID int `json:"genderId"`
	SeasonWins int `json:"seasonWins"`
	SeasonLosses int `json:"seasonLosses"`
	WeeklyWins int `json:"weeklyWins"`
	WeeklyLosses int `json:"weeklyLosses"`
}
