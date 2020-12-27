package dao

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"history-graph-notes-server/util"
	"log"
	"strings"
)

func GetNodeList(label string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return "", err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		cypher := `MATCH (a:%s) WHERE a.name=$name RETURN a.name`
		log.Printf("cypher:%s", fmt.Sprintf(cypher, label))

		result, err := tx.Run(fmt.Sprintf(cypher, label), nil)
		if err != nil {
			return nil, err
		}

		for result.Next() {
			if result.Record().GetByIndex(0) != nil {
				switch result.Record().GetByIndex(0).(type) {
				case []interface{}:
					strlist := util.ToStringSlice(result.Record().GetByIndex(0).([]interface{}))
					list = append(list, strings.Join(strlist, ";"))
				case interface{}:
					list = append(list, result.Record().GetByIndex(0).(string))
				default:
					continue
				}
			}
		}

		if err = result.Err(); err != nil {
			log.Println("result err:", err)
			return nil, err
		}

		if len(list) > 0 {
			return list[0], nil
		}
		return "", util.ErrEmpty
	})

	if err != nil {
		log.Println("session ReadTransaction err:", err)
		return "", util.ErrEmpty
	}

	log.Println("GetFigureNodeProper from neo4j result", result)
	return result, nil
}
