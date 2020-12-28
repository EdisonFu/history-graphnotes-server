package dao

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"history-graph-notes-server/util"
	"log"
	"strings"
)

func GetNodeList() (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return "", err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		cypher := `MATCH (a) RETURN labels(a) as label, a.name as name`
		log.Printf("cypher:%s", cypher)

		result, err := tx.Run(fmt.Sprintf(cypher), nil)
		if err != nil {
			return nil, err
		}

		var label, name string
		for result.Next() {
			record := result.Record()
			l, _ := record.Get("label")
			n, _ := record.Get("name")

			if l == nil || n == nil {
				continue
			}

			label = util.ToStringSlice(l.([]interface{}))[0]
			switch n.(type) {
			case string:
				name = n.(string)
			case []interface{}:
				continue
			default:
				continue
			}

			if result.Record().GetByIndex(0) != nil {
				list = append(list, label+":"+name)
			}
		}

		if err = result.Err(); err != nil {
			log.Println("GetNodeList result err:", err)
			return nil, err
		}

		if len(list) > 0 {
			return list, nil
		}
		return "", util.ErrEmpty
	})

	if err != nil {
		log.Println("GetNodeList session ReadTransaction err:", err)
		return "", util.ErrEmpty
	}

	log.Println("GetNodeList from neo4j result", result)
	return result, nil
}

func GetFitNode(label, proper, key string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return "", err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		cypher := `MATCH (a:%s) WHERE a.%s=~'.*%s.*'  RETURN a.name`
		log.Printf("cypher:%s", fmt.Sprintf(cypher, label, proper, key))

		result, err := tx.Run(fmt.Sprintf(cypher, label, proper, key), nil)
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
			return list, nil
		}
		return "", util.ErrEmpty
	})

	if err != nil {
		log.Println("GetFitNode session ReadTransaction err:", err)
		return "", util.ErrEmpty
	}

	log.Println("GetFitNode from neo4j result", result)
	return result, nil
}
