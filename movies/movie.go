package main

import (
	"flag"
	"net/http"
	"fmt"
	"log"
	"errors"
	"encoding/json"
)

var (
	port int
	host string
)


const (
	defaultPort = 8087
	defaultHost = "0.0.0.0"
)

type request struct {
	StockNumer int
}

func handler(w http.ResponseWriter, r *http.Request) {
	//startTime := time.Now()
	query := r.FormValue("q")
	apiKey := r.FormValue("apikey")
	fmt.Printf("query="+query+"/ apiKey="+apiKey)
	if query == "" || apiKey == ""{
		renderError(w, errors.New("set query & apikey!"))
		return
	}

}

func renderOutput(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	output, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	w.Write(output)
}

func newErrorResult(err error) interface{} {
	return struct {
		Error string `json:"error"`
	}{
		err.Error(),
	}
}

func renderError(w http.ResponseWriter, err error) {
	result := newErrorResult(err)
	renderOutput(w, result)
}

func flags()  {
	// IntVar(p *int, name string, value int, usage string) {
	flag.IntVar(&port, "port", defaultPort, "port")
	flag.StringVar(&host, "host", defaultHost, "host")
	flag.Parse()
}

func main() {
	flags()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1", handler)
	fmt.Printf("Serving HTTP at: http://%s:%d\n", host, port)

	mux.HandleFunc("/movies", movies.Handle(movies.HandlerMovieInfo))



	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), mux); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}

}
