package dao

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

func AddNode(label string, proper map[string]string) (err error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return err
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		//CREATE (a:Person) SET a.name = "柏拉图", a.age = 200
		cypher := `CREATE (a:%s) SET `
		properMap := make(map[string]interface{})

		for proper, value := range proper {
			cypher = cypher + fmt.Sprintf("a.%s = $%s,", proper, proper)
			properMap[proper] = value
		}

		cypher = fmt.Sprintf(cypher, label)
		log.Printf("cypher:%s,\n,properMap:%v\n", cypher[:len(cypher)-1], properMap)

		_, err := tx.Run(cypher[:len(cypher)-1], properMap)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		log.Println("AddNode session WriteTransaction err:", err)
		return err
	}

	log.Println("AddNode to neo4j success:", proper)
	return
}

func AddNodeRelation(nameA, labelA, nameB, labelB, relationType string, proper map[string]string) (err error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return err
	}
	defer session.Close()

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		//cypher := `match (a:Person{name:"亚里士多德"}),(b:Person{name:"柏拉图"}) create (a) -[r:STUDY{year:666}]->(b)`
		cypher := `match (a:%s{name:$nameA}),(b:%s{name:$nameB}) create (a) -[r:%s{`
		properMap := make(map[string]interface{})
		properMap["nameA"] = nameA
		properMap["nameB"] = nameB

		for proper, value := range proper {
			cypher = cypher + fmt.Sprintf("%s:$%s,", proper, proper)
			properMap[proper] = value
		}
		cypher = fmt.Sprintf(cypher, labelA, labelB, relationType)
		cypher = cypher[:len(cypher)-1] + "}]->(b)"
		log.Printf("cypher:%s\n,properMap:%v\n", cypher, properMap)

		_, err := tx.Run(cypher, properMap)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})

	if err != nil {
		log.Println("AddNodeRelation session WriteTransaction err:", err)
		return err
	}

	log.Println("AddNodeRelation to neo4j success:", proper)
	return
}
