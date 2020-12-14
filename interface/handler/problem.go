package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type problemJSON struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

//HTTPProblem writes a Problem Json response (rfc7807), or fallbacks to http.Error if something goes wrong
func HTTPProblem(w http.ResponseWriter, status int, detail string) {
	problem := problemJSON{
		Title:  http.StatusText(status),
		Status: status,
	}

	var body []byte
	var err error
	if detail != "" {
		var problemWithDetail struct {
			problemJSON
			Detail string `json:"detail"`
		}
		problemWithDetail.problemJSON = problem
		problemWithDetail.Detail = detail
		body, err = json.Marshal(problemWithDetail)
	} else {
		body, err = json.Marshal(problem)
	}

	if err != nil {
		http.Error(w, http.StatusText(status), status)
		return
	}

	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	w.WriteHeader(status)
	fmt.Fprint(w, string(body))
}
