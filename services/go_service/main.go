package main

import (
	"fmt"

	"github.com/gmbh-micro/gmbh"
)

/*
 *  This service is an example of an interal service that has two reg
 */

func main() {

	// configure runtime settings
	runtime := gmbh.SetRuntime(gmbh.RuntimeOptions{Blocking: true, Verbose: true})

	// configure service setings
	service := gmbh.SetService(gmbh.ServiceOptions{Name: "gos", PeerGroups: []string{"universal"}})

	// instantiate the client
	client, err := gmbh.NewClient(runtime, service)
	if err != nil {
		panic(err)
	}

	// register routes
	client.Route("testOne", handleOne)
	client.Route("testTwo", handleTwo)

	// start the client
	client.Start()
}

// handle functions in gmbh-go behave similarly to how the default http package works.
// All of the incoming data is encapsulated by the gmbh.Request object and then assign
// return fields in the gmbh.Responder object to be sent back to the original caller.
func handleOne(req gmbh.Request, resp *gmbh.Responder) {

	// get the payload from the request
	data := req.GetPayload()

	// instantiate the return payload with the result data
	payload := gmbh.NewPayload()

	// add the resulting data the the client
	payload.AppendStringField(
		"result", fmt.Sprintf("hello from gos test 1; returning same message; message=%s", data.GetStringField("xid")),
	)

	// attach the payload to the responder
	resp.SetPayload(payload)
}

// see handleOne above
func handleTwo(req gmbh.Request, resp *gmbh.Responder) {
	data := req.GetPayload()
	payload := gmbh.NewPayload()
	payload.AppendStringField(
		"result", fmt.Sprintf("hello from gos test 2; returning same message; message=%s", data.GetStringField("xid")),
	)
	resp.SetPayload(payload)
}
