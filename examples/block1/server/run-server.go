package main

import (
	"github.com/zubairhamed/canopus"
	"log"
)

func main() {
	server := canopus.NewLocalServer()

	server.Get("/blockinfo", func(req canopus.CoapRequest) canopus.CoapResponse {
		log.Println("Hello Called")
		// canopus.PrintMessage(req.GetMessage())
		msg := canopus.ContentMessage(req.GetMessage().MessageID, canopus.MessageAcknowledgment)
		msg.SetStringPayload("Acknowledged: " + req.GetMessage().Payload.String())
		res := canopus.NewResponse(msg, nil)

		return res
	})

	server.Post("/blockupload", func(req canopus.CoapRequest) canopus.CoapResponse {
		log.Println("Hello Called via POST")
		// canopus.PrintMessage(req.GetMessage())
		msg := canopus.ContentMessage(req.GetMessage().MessageID, canopus.MessageAcknowledgment)
		msg.SetStringPayload("Acknowledged: " + req.GetMessage().Payload.String())
		res := canopus.NewResponse(msg, nil)

		return res
	})

	server.OnBlockMessage(func(msg *canopus.Message, inbound bool) {
		log.Println("Incoming Block Message:")
		canopus.PrintMessage(msg)
	})

	server.Start()
}