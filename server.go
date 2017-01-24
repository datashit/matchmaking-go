// Copyright 2017 The Yiğit YILDIRIM<yigit@yildirim.me> Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matchmaking

import "github.com/datashit/matchmaking-go/Message"

type procesJob struct {
	m   *mmg
	req matchmaking.Request
}

var jobs = make(chan procesJob, 1000)

// CreateProcesWorker create pool for proces function
// workerSize The number of Goroutine will work
func CreateProcesWorker(workerSize int) {
	for w := 1; w <= workerSize; w++ {
		go proces(jobs)
	}
}

func proces(job <-chan procesJob) {
	for j := range job {
		switch j.req.Command {
		case "MATCH":
		case "LOBBY":
		case "CHAT":
		case "CONTACTS":
		case "PARTY":
		}
	}
}
