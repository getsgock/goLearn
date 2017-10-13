// myhttp project myhttp.go
package myhttp

import (
	"net/http"
)

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func ListenAndServe(address string, h Handler) error {
	err := http.ListenAndServe(address, h)
	return err
}

//type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) { f(w, r) }
