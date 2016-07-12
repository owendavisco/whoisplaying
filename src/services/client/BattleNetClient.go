package client

import (
	shared "services/shared"
	model "services/model"
	"net/http"
	"encoding/json"
)

type BattleNetClient struct {
	apiKey string
}

func (btnClient BattleNetClient) Get3v3Leaderboard() [] model.CharacterPVPStats {
	leaderboardUrl := shared.BuildLeaderboardUrl(shared.US, shared.THREE_VS_THREE, btnClient.apiKey, shared.DEFAULT_LOCALE)
	leaderboard := getJson(leaderboardUrl, model.Leaderboard{})

	return leaderboard.Rows
}

func getJson(url string, target interface{}) model.Leaderboard {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}