package main

import (
	"net/http"
	"errors"
	"encoding/json"
)

func personService(w http.ResponseWriter, r *http.Request)  {

	query := r.FormValue("q")
 
	if query == "" {
		renderError(w, errors.New("NO Query!!"))
		return
	}
}

func renderError(w http.ResponseWriter, err error) {
	result := newErrorResult(err)
	renderOutput(w, result)

}
func newErrorResult(err error) interface{}{
	return struct {
		Error string `json:"error"`
	}{
		err.Error(),
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