package main

import (
	"history-graph-notes-server/dao"
	"history-graph-notes-server/handlers"
	"log"
	"net/http"
)

func main() {
	dao.InitNeo4jDB()
	defer dao.CloseNeo4j()

	handlers.Init()
	log.Println("router init ok!")
	http.ListenAndServe(":8080", nil)
}
