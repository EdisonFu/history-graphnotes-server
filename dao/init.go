package dao

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	. "history-graph-notes-server/model"
	"log"
)

type neo4jConnPool struct {
	Driver neo4j.Driver
}

var neo4jConn *neo4jConnPool

func InitNeo4jDB() {
	driver, err := neo4j.NewDriver(Uri, neo4j.BasicAuth(Username, Password, ""), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		log.Println("error connecting to neo4j:", err)
		panic("init DB err!")
	}

	neo4jConn = &neo4jConnPool{Driver: driver}
}

func GetNeo4jConn() neo4j.Driver {
	return neo4jConn.Driver
}

func CloseNeo4j() {
	neo4jConn.Driver.Close()
}
