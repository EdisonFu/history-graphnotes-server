package dao

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	. "history-graph-notes-server/model"
	"history-graph-notes-server/util"
	"log"
	"strings"
)

func GetFigureSingleProper(name string, proper string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return "", err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		cypher := `MATCH (a:Person) WHERE a.name=$name RETURN a.%s`
		log.Printf("cypher:%s,\n,properMap:%v\n", fmt.Sprintf(cypher, proper), map[string]interface{}{"name": name})

		result, err := tx.Run(fmt.Sprintf(cypher, proper), map[string]interface{}{"name": name})
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

func GetFigureNodeProper(name string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return nil, err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []*HistoryFigure

		result, err := tx.Run("MATCH (a:Person) WHERE a.name=$name RETURN "+
			"a.name as name, a.country as country, a.birthday as birthday, "+
			"a.homeland as homeland, a.occupation as occupation, a.achievements as achievements, a.works as works",
			map[string]interface{}{"name": name})
		if err != nil {
			log.Println("tx run err:", err)
			return nil, err
		}

		for result.Next() {
			figure := new(HistoryFigure)
			record := result.Record()
			name, _ := record.Get("name")
			country, _ := record.Get("country")
			birthday, _ := record.Get("birthday")
			homeland, _ := record.Get("homeland")
			occupation, _ := record.Get("occupation")
			achievements, _ := record.Get("achievements")
			works, _ := record.Get("works")

			if name != nil {
				figure.Name = name.(string)
			}
			if country != nil {
				figure.Country = country.(string)
			}
			if birthday != nil {
				figure.Birthday = birthday.(string)
			}
			if homeland != nil {
				figure.Homeland = homeland.(string)
			}
			if occupation != nil {
				figure.Occupation = occupation.(string)
			}
			if achievements != nil {
				figure.Achievements = achievements.(string)
			}
			if works != nil {
				figure.Works = works.(string)
			}

			if result.Record().GetByIndex(0) != nil {
				list = append(list, figure)
			}
		}

		if err = result.Err(); err != nil {
			log.Println("GetFigureNode result err:", err)
			return nil, err
		}

		if len(list) > 0 {
			return list[0], nil
		}
		return "", util.ErrEmpty
	})
	if err != nil {
		log.Println("GetFigureNode session ReadTransaction err:", err)
		return "", util.ErrEmpty
	}

	log.Println("GetFigureSingleProper from neo4j result", result)
	return result, nil
}
