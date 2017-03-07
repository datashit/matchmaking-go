package matchmaking

import (
	"encoding/json"
	"log"
)

type sFind struct {
	Location string // Game server Location
	Mode     string // Game mode
	Map      string // Game map
	Rank     bool   // Rank mode
	PartyID  string // Party id
}

type sfindJob struct {
	UserID string // User id
	m      *mmg   // websocket
	sfind  sFind  // Find Settings
}

func matchFind(job *procesJob) {

	var sfjob sfindJob
	sfjob.UserID = job.req.UserID
	sfjob.m = job.m

	json.Unmarshal(job.req.Data, &sfjob.sfind)

	log.Println("matchFind")
	findjobs <- sfjob
	log.Println("matchFind job sent")

}
