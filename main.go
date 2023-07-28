package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options
var state State

func reader(c *websocket.Conn) {
	defer c.Close()
	msgchan := handler()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		msgchan <- string(message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func handler() chan string {
	msgchan := make(chan string)
	go func() {
		for {
			select {
			case m := <-msgchan:
				log.Printf("message: %s", m)
				var msg Event

				switch msg.Type {
				case Login:
					state.Players["344234234234224"] = Player{}
				}
			default:
				fmt.Println("fart")
			}
		}
	}()
	return msgchan
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	reader(c)

}

func main() {
	state = State{}
	http.HandleFunc("/ws", handleWebsocket)

	http.ListenAndServe(":8080", nil)
}
