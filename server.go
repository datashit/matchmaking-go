// Copyright 2017 The Yiğit YILDIRIM<yigit@yildirim.me> Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matchmaking

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/datashit/matchmaking-go/Message"
)

type procesJob struct {
	m   *mmg
	req matchmaking.Request
}

var jobs = make(chan procesJob, 1000)

// ServerName server name
var ServerName string

// CreateProcesWorker create pool for proces function
// workerSize The number of Goroutine will work
func CreateProcesWorker(workerSize int) {
	for w := 1; w <= workerSize; w++ {
		go proces(jobs)
	}
	CreateMatchWorker(workerSize)
}

func proces(job <-chan procesJob) {
	for j := range job {
		switch j.req.Function {
		case "MATCH":
			match(&j)
		case "LOBBY":
		case "CHAT":
		case "CONTACTS":
		case "PARTY":
		default:
			var x matchmaking.Response
			x.Command = "BACKDEF"
			x.MessageType = j.req.MessageType
			x.Data = []byte("NO DATA")
			x.ServerID = j.m.name
			j.m.sender <- x
		}
	}
}

var simultaneous int32

// Accept function incomming client accept and join matchmaking server.
func Accept(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Simultaneous: %v \r\n", atomic.AddInt32(&simultaneous, 1))

	m := newClient("Client")
	m.RUN(w, r)

	fmt.Printf("Simultaneous: %v \r\n", atomic.AddInt32(&simultaneous, -1))
}

// INIT Initializes mmg server name: Server Name workerSize: worker pool size
func INIT(name string, workerSize int) {
	ServerName = name
	CreateProcesWorker(workerSize)
}
