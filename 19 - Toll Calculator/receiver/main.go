package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/grayjunzi/tolling/types"
)

func main() {
	recv := NewReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":3000", nil)
}

type Receiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
}

func NewReceiver() *Receiver {
	return &Receiver{
		msgch: make(chan types.OBUData, 128),
	}
}

func (recv *Receiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	recv.conn = conn
	go recv.wsReceiveLoop()
}

func (recv *Receiver) wsReceiveLoop() {
	fmt.Println("NEW OBU client connected")
	for {
		var data types.OBUData
		if err := recv.conn.ReadJSON(&data); err != nil {
			log.Printf("read error: %w+", err)
			continue
		}
		fmt.Printf("received OBU data from %d :: <lat %.2f, long %.2f>\n", data.OBUID, data.Lat, data.Long)
		recv.msgch <- data
	}
}
