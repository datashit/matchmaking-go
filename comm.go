// Copyright 2017 The Yiğit YILDIRIM<yigit@yildirim.me> Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matchmaking

import (
	"encoding/json"
	"log"
	"net/http"

	matchmaking "github.com/datashit/matchmaking-go/Message"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options
type mmg struct {
	name   string
	c      *websocket.Conn
	quit   chan error
	sender chan matchmaking.Response
}

func (m *mmg) Close() {
	close(m.quit)
	close(m.sender)
}

// NewMMG create mmg and  return created mmg
func newClient(Name string) mmg {
	return mmg{name: Name, // Set client name
		quit:   make(chan error, 1),                // Open channel error
		sender: make(chan matchmaking.Response, 1)} // Open channel Response

}

func (m *mmg) WriteMessage(r matchmaking.Response) error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	msg := matchmaking.Encoder(data)

	err = m.c.WriteMessage(r.MessageType, msg)
	return err
}

func (m *mmg) Writer() {

}

func (m *mmg) RUN(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	m.c = c // Added Conn

	go func() { // Start Reader Goroutine
		log.Printf("%v Read GO", m.name)
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				m.quit <- err // Send close command
				<-m.quit      // Wait closing socket
				break         // closed Goroutine
			}
			log.Printf("%v - recv: %s", m.name, message)

			var req procesJob
			req.m = m

			json.Unmarshal(matchmaking.Decoder(message), &req.req) // Decode message
			req.req.MessageType = mt                               // Message Type Added

			jobs <- req //  send to Job
		}
	}()

	for {
		select {
		case s := <-m.sender:
			data, err := json.Marshal(s)
			if err != nil {
				continue
			}
			c.WriteMessage(s.MessageType, data)
			log.Printf("%s - send: %v", m.name, s) // Send data
		case q := <-m.quit: // IF socket error close to websocket
			log.Println(q)
			m.quit <- q // send to close flag for runing goroutine
			return
		}
	}
}
