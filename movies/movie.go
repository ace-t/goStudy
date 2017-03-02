package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"io/ioutil"
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
	fmt.Printf("query=" + query + "/ apiKey=" + apiKey)
	if query == "" || apiKey == "" {
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

func flags() {
	// IntVar(p *int, name string, value int, usage string) {
	flag.IntVar(&port, "port", defaultPort, "port")
	flag.StringVar(&host, "host", defaultHost, "host")
	flag.Parse()
}


func WrapHandle(handlerFunc http.HandlerFunc) http.HandlerFunc {
	var i int
	i = 10  // 클로저
	return func(w http.ResponseWriter, r *http.Request) {
		i += 10
		//w.Write([]byte("Start Movie Handler!"))
		handlerFunc(w, r)
		//w.Write([]byte("after handler"))
		fmt.Println(i)
	}
}


type MovieInfoResult struct {
	Channel struct{
		Q string `json:"q"`
		Item [] struct{
			Genre [] struct{
				Content string `json:"content"`
			} `json:"genre"`
			Story [] struct{
				Content string `json:"content"`
			} `json:"story"`
			Open_info [] struct{
				Content string `json:"content"`
			}`json:"open_info"`
		} `json:"item"`
	}`json:"channel"`
}

// personResult, err := PersonRequest(personName)

func HandlerMovieInfo(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("movie info"))

	movieResult, err := getMovieInfo(r.FormValue("q"))
	if err != nil {
		panic(err)
	}
	renderOutput(w, movieResult)
}

func getMovieInfo(query string) (*MovieInfoResult, error)  {
	baseUrl := "https://apis.daum.net/contents/movie?apikey=29c7e60b0deec823061babb82691aeb7&output=json"
	u, err := url.Parse(baseUrl) // u :  https://apis.daum.net/contents/movie?apikey=29c7e60b0deec823061babb82691aeb7&output=json
	fmt.Println(u)
	if err != nil {
		panic(err) // panic()함수는 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴한다.
	}

	q := u.Query()  // q => map[apikey:[29c7e60b0deec823061babb82691aeb7] output:[json]]
	q.Set("q", query) // map[apikey:[29c7e60b0deec823061babb82691aeb7] output:[json] query:[미생]]

	u.RawQuery = q.Encode() // https://apis.daum.net/contents/movie?apikey=29c7e60b0deec823061babb82691aeb7&output=json&q=%!E(MISSING)B%!A(MISSING)F%!B(MISSING)8%!E(MISSING)C%9D

	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	fmt.Println(u.String())
	defer resp.Body.Close() // defer : defer를 호출하는 함수가 리턴하기 직전에 호출(java finally 구문 같은?

	body, err := ioutil.ReadAll(resp.Body) // Body : io.ReadCloser

	var result MovieInfoResult
	if err := json.Unmarshal(body, &result); err != nil {
		panic(err)
	}

	return &result, nil
}

func init() {
	fmt.Println("main 이전에 호출 된다!")

}

func main() {
	flags()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1", handler)
	fmt.Printf("Serving HTTP at: http://%s:%d\n", host, port)

	mux.HandleFunc("/movies", WrapHandle(HandlerMovieInfo))

	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), mux); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}

}
