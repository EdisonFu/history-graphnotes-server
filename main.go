package main

import (
	"github.com/prometheus/common/log"
	"history-graph-notes-server/handlers"
	"net/http"
)

func main() {
	handlers.Init()
	log.Info("router init ok!")
	http.ListenAndServe(":8080", nil)
}
