package handlers

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request)  {
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if (err != nil) {
		http.Error(rw, "bad request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "hello %s\n", d)
}