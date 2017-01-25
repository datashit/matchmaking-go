package matchmaking

import "encoding/json"

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

	var sf sFind
	json.Unmarshal(job.req.Data, &sf)

}
