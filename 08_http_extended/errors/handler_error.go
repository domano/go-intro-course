package main

import (
	"fmt"
	"net/http"
	"strings"
	"errors"
)

type handler string

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Accept"), "text/plain") {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	err := somethingGoesWrong()
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%v\n", h)
}
func somethingGoesWrong() error {
	return errors.New("Something went wrong!")
}

func main() {
	panic(http.ListenAndServe(":8080", handler("Hello World")))
}
