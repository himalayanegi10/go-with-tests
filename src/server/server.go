package server

import (
	"fmt"
	"net/http"
)

func PlayerServer(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "20")
}