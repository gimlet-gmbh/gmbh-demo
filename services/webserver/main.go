package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gmbh-micro/gmbh"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

var client *gmbh.Client

func main() {

	// set runtime options.
	// Note that because we will be running a webserver, it is fine for gmbh to not
	// run in blocking mode because the ListenAndServe will block the main thread
	// below
	runtime := gmbh.SetRuntime(gmbh.RuntimeOptions{Blocking: false, Verbose: true})

	// set the service options
	service := gmbh.SetService(gmbh.ServiceOptions{Name: "ws"})

	var err error
	// NewClient returns a new gmbh client configured for use with any of the options below
	client, err = gmbh.NewClient(runtime, service)
	if err != nil {
		panic(err)
	}

	// start the client
	client.Start()

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	// r.HandleFunc("/", handleIndex)
	r.HandleFunc("/gos/1", handleGos1)

	http.DefaultClient.Timeout = time.Minute
	log.Fatal(http.ListenAndServe(":2121", r))

}

// handleGos1 is an example of a function that will make a request via gmbh and write
// the result to the response writer
func handleGos1(w http.ResponseWriter, r *http.Request) {

	// create a new payload object
	payload := gmbh.NewPayload()

	// in this case we are going to assign a random ID string to the outgoing payload
	payload.AppendStringField("xid", xid.New().String())

	// the client will return a resulting payload and a potential error
	result, err := client.MakeRequest("gos", "testOne", payload)
	// this err is returned as non nil if gmbh has connectivity issues with the underlying
	// architecture.
	if err != nil {
		w.Write([]byte("could not contact; err=" + err.Error()))
		return
	}

	// This error is a result of a failure to resolve the request within gmbh
	if result.GetError() != "" {
		w.Write([]byte("Service error: " + result.GetError() + "\n"))
		return
	}

	// The happy path; write the result as set by the target to the writer
	w.Write([]byte("Gos1 -> " + result.GetPayload().GetStringField("result") + "\n"))

}
