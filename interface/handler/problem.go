package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type problemJSON struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
}

//HTTPProblem writes a Problem Json response (rfc7807), or fallbacks to http.Error if something goes wrong
func HTTPProblem(w http.ResponseWriter, status int) {
	problem, err := json.Marshal(problemJSON{
		Title:  http.StatusText(status),
		Status: status,
	})

	if err != nil {
		http.Error(w, http.StatusText(status), status)
		return
	}

	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.WriteHeader(status)
	fmt.Fprint(w, string(problem))
}
