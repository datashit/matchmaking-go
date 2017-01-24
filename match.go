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
	name string
	c    *websocket.Conn
	quit chan error
}

func (m *mmg) proces(r *matchmaking.Request) {

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
		case q := <-m.quit: // IF fatal error close websocket
			log.Println(q)
			break
		default:
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)

			var req matchmaking.Request
			json.Unmarshal(matchmaking.Decoder(message), &req) // Decode message
			req.MessageType = mt                               // Message Type Added

			m.proces(&req)
		}

	}
}
