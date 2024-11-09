package server

import (
	"strings"
	"fmt"
	"net/http"
)

func PlayerServer(res http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	// if player == "Pepper" {
	// 	fmt.Fprint(res, "20")
	// } else if player == "Floyd" {
	// 	fmt.Fprint(res, "10")
	// }
	fmt.Fprint(res, GetPlayerScore(player))
}

func GetPlayerScore(name string) string {
	switch name {
	case "Pepper":
		return "20"
	case "Floyd":
		return "10"
	default:
		return ""
	}
}