package main

import (
	"fmt"
	"flag"
	"net/http"
	"github.com/gorilla/mux"
)

var (
	listenPort int
	routePath  string
	respMsg    string
)

func initFlags() {
	flag.IntVar(&listenPort, "port", 8080, "listening port")
	flag.StringVar(&routePath, "routepath", "/", "route path to receive the http requests")
	flag.StringVar(&respMsg, "respmsg", "Hello world!", "the response message to be sent back to clients")

	flag.Parse()
}

func main() {
	fmt.Println("starting...")
	initFlags()

	r := mux.NewRouter()
	r.HandleFunc(routePath, func (w http.ResponseWriter, r *http.Request) {
		fmt.Println(respMsg)
		w.Write([]byte(respMsg))
	})

	addr := fmt.Sprintf(":%d", listenPort)
	http.ListenAndServe(addr, r)
	fmt.Println("exiting...")
}


