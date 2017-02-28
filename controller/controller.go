package main

import (
	"flag"
	"net/http"
	"fmt"
	"log"
)

var (
	port int
	host string
)

const (
	defaultPort = 8088
	defaultHost = "0.0.0.0"
)


func flags()  {
	// IntVar(p *int, name string, value int, usage string) {
	flag.IntVar(&port, "port", defaultPort, "port")
	flag.StringVar(&host, "host", defaultHost, "host")
	flag.Parse()
}

func main()  {
	flags()
	mux := http.NewServeMux()
	mux.HandleFunc("/api/person", personService)
	fmt.Printf("Serving HTTP at: http://%s:%d\n", host, port)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), mux); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}

}