package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func reader(conn *websocket.Conn) {
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(p))

	if err := conn.WriteMessage(messageType, p); err != nil {
		log.Println(err)
		return
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client successfully connected")
	reader(ws)
}

func SetupRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ws", wsEndpoint)
}

func Listen() {
	port := ":7000"
	fmt.Println("Server is running on port ", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
