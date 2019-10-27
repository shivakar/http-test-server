package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"os"

	gor "github.com/gorilla/handlers"
	"github.com/shivakar/http-test-server/pkg/handlers"
)

var (
	addr = flag.String("addr", "0.0.0.0:8080", "bind address")

	localIPs = []string{}
)

func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range addrs {
		localIPs = append(localIPs, v.String())
	}
}

func main() {
	r := handlers.NewRouter()
	r.Use(gor.CompressHandler)
	srv := gor.CombinedLoggingHandler(os.Stdout, r)

	log.Printf("Starting server on '%v'", *addr)
	log.Println(http.ListenAndServe(*addr, srv))
}
