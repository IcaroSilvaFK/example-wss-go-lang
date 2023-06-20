package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Room struct {
	clients []*websocket.Conn
}

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var room Room

func NewWssController(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	defer conn.Close()
	room.clients = append(room.clients, conn)

	for {

		_, message, err := conn.ReadMessage()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			break
		}

		msg := string(message)

		fmt.Println(msg)

		for _, client := range room.clients {
			if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				break
			}
		}
	}

}
