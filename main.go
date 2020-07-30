package main

import (
	"log"
	"net/http"
)

var cmd Cmd
var srv http.Server

func StartServer(bind string, param string) {
	log.Printf("Listening on %s, query proxy header to %s", bind, param)
	h := &handle{param: param}
	srv.Addr = bind
	srv.Handler = h
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func StopServer() {
	if err := srv.Shutdown(nil); err != nil {
		log.Println(err)
	}
}

func main() {
	cmd = parseCmd()
	StartServer(cmd.bind, cmd.param)
}
