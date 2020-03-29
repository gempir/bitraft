package main

import (
	"log"

	"code.cerinuts.io/libs/goPurple/firehose"
)

var broadcast = make(chan socketMessage)

type socketMessage struct {
	Channel string `json:"channel"`
	Message string `json:"message"`
}

func main() {

	firehose := &firehose.Firehose{}
	msgQ, err := firehose.Connect("oauth")
	if err != nil {
		log.Fatal("firehose error: " + err.Error())
	}

	for event := range msgQ {
		log.Println(event)
	}

	startWebsocketServer()
}
