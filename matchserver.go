package matchmaking

import (
	"log"

	matchmaking "github.com/datashit/matchmaking-go/Message"
)

var findjobs = make(chan sfindJob, 1000)

// CreateMatchWorker create pool for proces function
// workerSize The number of Goroutine will work
func CreateMatchWorker(workerSize int) {
	for w := 1; w <= workerSize; w++ {
		go matchProces(findjobs)
	}
}

func matchProces(job <-chan sfindJob) {
	for j := range job {
		log.Println("Work Match Proces")

		var res matchmaking.Response
		res.MessageType = 1
		res.Function = "MATCH"
		res.Command = "FIND"
		j.m.sender <- res
	}
}
