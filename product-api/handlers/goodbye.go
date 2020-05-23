package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
 l *log.Logger
}

func NewGoodbye(l *log.Logger)  *Goodbye {
	return & Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, w * http.Request)  {
	g.l.Println("goodbye")
	rw.Write([] byte("byeee\n"))
}