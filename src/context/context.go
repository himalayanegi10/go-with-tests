package mycontext

import (
	"net/http"
	"fmt"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
}

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}

