package service

import (
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/websocket"
)

type Sock struct {
}

func echoHandler(ws *websocket.Conn) {

	// n, err := ws.Read(msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Receive: %s\n", msg[:n])

	// send_msg := "[" + string(msg[:n]) + "]"
	for {
		i := <-Chant

		send_msg := strconv.Itoa(i)
		_, err := ws.Write([]byte(send_msg))
		if err != nil {
			log.Println(err)
		}

	}

}

func (this *Sock) Start() {
	http.Handle("/echo", websocket.Handler(echoHandler))

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
