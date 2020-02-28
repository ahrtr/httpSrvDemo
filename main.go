package main

import (
	"fmt"
	"flag"
	"net/http"
	"github.com/gorilla/mux"
)
const (
	XRealIP       = "X-Real-IP"
	XForwardedFor = "X-Forwarded-For"
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
		printRemoteIP(r)
		fmt.Println(respMsg)
		w.Write([]byte(respMsg))
	})

	addr := fmt.Sprintf(":%d", listenPort)
	http.ListenAndServe(addr, r)
	fmt.Println("exiting...")
}

func printRemoteIP(req *http.Request) {
	fmt.Printf("req.RemoteAddr = %s\n", req.RemoteAddr)
	if realIP := req.Header.Get(XRealIP); realIP != "" {
		fmt.Printf("req.Header.Get(%s) = %s\n", XRealIP, realIP)
	} else {
		fmt.Printf("req.Header.Get(%s) not found\n", XRealIP)
	}
	
	if forwardIP := req.Header.Get(XForwardedFor); forwardIP != "" {
		fmt.Printf("req.Header.Get(%s) = %s\n", XForwardedFor, forwardIP)
	} else {
		fmt.Printf("req.Header.Get(%s) not found\n", XForwardedFor)
	}
}

