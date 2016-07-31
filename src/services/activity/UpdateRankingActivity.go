package activity

import "services/client"

const (
	requestedState = State{
		Description:"Requested update ranking",
	}
	getArenaDataState = State{
		Description:"Getting arena data",
	}
	committingChangesState = State{
		Description:"Committing updates to database",
	}
	completedState = State{
		Description:"Completed ranking update, returning results",
	}
	faultState = State{
		Description:"An error occured",
	}
)

type UpdateRankingActivity struct {
	currentStatus Status
	currentState State
	percentComplete Percent
}

func (ua *UpdateRankingActivity) Start(btnClient client.BattleNetClient, dbClient client.DatabaseClient) {
	ua.currentState = requestedState
	ua.currentStatus = Running
	ua.percentComplete = 0

	oldArenaData := make(chan []string)
	newArenaData := make(chan []string)
	go getOldArenaData(dbClient, oldArenaData)
	go getNewArenaData(btnClient, newArenaData)


}

func compareAndExtract() {

}

func getOldArenaData(dbClient client.DatabaseClient, results chan []string) {
	results <- string{""}
}

func getNewArenaData(btnClient client.BattleNetClient, results chan []string) {
	results <- string{""}
}