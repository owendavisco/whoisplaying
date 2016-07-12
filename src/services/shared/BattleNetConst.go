package services

// Battle.net url constants
const (
	BASE_URL = "%s.api.battle.net"
	DEFAULT_SCHEME = "https"
)

// Header constants
const (
	DEFAULT_ACCEPT = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"
)

// Path constants
const (
	WOW_LEADERBOARD_PATH  = "/wow/leaderboard/%s"
	CHARACTOR_PATH = "/character/%s"
)

// Query keys
const (
	LOCALE = "locale"
	API_KEY = "apiKey"
)

// Query values
const (
	DEFAULT_LOCALE = "en_Us"
	THREE_VS_THREE = "3v3"
	TWO_VS_TWO = "2v2"
)

// Regions
const (
	US = "us"
	EU = "eu"
	KR = "kr"
	TW = "tw"
)
