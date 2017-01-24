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

var upgrader = websocket.Upgrader{} // use default options
type mmg struct {
	name   string
	c      *websocket.Conn
	quit   chan error
	sender chan bool
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
func (m *mmg) RUN(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	m.c = c

	for {
		select {
		case s := <-m.sender:
			log.Println(s) // Send data
		case q := <-m.quit: // IF socket error close to websocket
			log.Println(q)
			break
		default:
			mt, message, err := c.ReadMessage()
			if err != nil {
				m.quit <- err
			}
			log.Printf("recv: %s", message)

			var req procesJob
			req.m = m

			json.Unmarshal(matchmaking.Decoder(message), &req.req) // Decode message
			req.req.MessageType = mt                               // Message Type Added

			jobs <- req //  Job send to channel

		}

	}
}
