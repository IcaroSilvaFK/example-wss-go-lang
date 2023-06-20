package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	clientId string
	conn     *websocket.Conn
}

type MessagePayload struct {
	ClientId string `json:"client_id"`
	Message  string `json:"message"`
}

type Room struct {
	clients []Client
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
	userId := uuid.NewString()
	room.clients = append(room.clients, Client{clientId: userId, conn: conn})

	conn.WriteJSON(MessagePayload{
		ClientId: userId,
		Message:  "Connected",
	})

	for {

		var message MessagePayload

		err := conn.ReadJSON(&message)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			break
		}

		for _, client := range room.clients {

			if message.ClientId != client.clientId {
				if err := client.conn.WriteJSON(message); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					break
				}
			}

		}
	}

}
