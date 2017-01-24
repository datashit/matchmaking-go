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

var jobs = make(chan procesJob, 1000)

type procesJob struct {
	m   *mmg
	req matchmaking.Request
}

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
