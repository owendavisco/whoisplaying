package services

import (
	"net/url"
	"fmt"
)

func BuildLeaderboardUrl(region string, bracket string, apiKey string, locale string) string {
	// Build encoded url
	var leaderboardUrl = url.URL{
		Scheme: DEFAULT_SCHEME,
		Host: fmt.Sprint(BASE_URL, region),
		Path: fmt.Sprint(WOW_LEADERBOARD_PATH, bracket),
		RawQuery: url.Values{
			LOCALE: locale,
			API_KEY: apiKey,
		}.Encode(),
	}

	// Return to string of the encoded url
	return leaderboardUrl.String()
}