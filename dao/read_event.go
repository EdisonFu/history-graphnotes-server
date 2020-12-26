package dao

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	. "history-graph-notes-server/model"
	"history-graph-notes-server/util"
	"log"
	"strings"
)

func GetEventSingleProper(name string, proper string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return "", err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		cypher := `MATCH (a:Event) WHERE a.name=$name RETURN a.%s`
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

func GetEventNodeProper(name string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return nil, err
	}
	defer session.Close()

	events, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []*HistoryEvent

		result, err := tx.Run("MATCH (a:Event) WHERE a.name=$name RETURN "+
			"a.name as name, a.aliasName as aliasName, a.ocurrTime as ocurrTime, "+
			"a.location as location, a.participantGroups as participantGroups,"+
			"a.mainParticipantFigures as mainParticipantFigures, a.description as description, a.meaning as meaning",
			map[string]interface{}{"name": name})
		if err != nil {
			log.Println("GetEventNode tx run err:", err)
			return nil, err
		}

		for result.Next() {
			event := new(HistoryEvent)
			record := result.Record()
			name, _ := record.Get("name")
			aliasName, _ := record.Get("aliasName")
			ocurrTime, _ := record.Get("ocurrTime")
			location, _ := record.Get("location")
			participantGroups, _ := record.Get("participantGroups")
			mainParticipantFigures, _ := record.Get("mainParticipantFigures")
			description, _ := record.Get("description")
			meaning, _ := record.Get("meaning")

			if name != nil {
				event.Name = name.(string)
			}
			if aliasName != nil {
				event.AliasName = aliasName.(string)
			}
			if ocurrTime != nil {
				event.OcurrTime = ocurrTime.(string)
			}
			if location != nil {
				event.Location = location.(string)
			}
			if description != nil {
				event.Description = description.(string)
			}
			if meaning != nil {
				event.Meaning = meaning.(string)
			}
			if participantGroups != nil {
				event.ParticipantGroups = util.ToStringSlice(participantGroups.([]interface{}))
			}
			if mainParticipantFigures != nil {
				event.MainParticipantFigures = util.ToStringSlice(mainParticipantFigures.([]interface{}))
			}

			if result.Record().GetByIndex(0) != nil {
				list = append(list, event)
			}
		}

		if err = result.Err(); err != nil {
			log.Println("GetEventNode result err:", err)
			return nil, err
		}

		if len(list) > 0 {
			return list[0], nil
		}
		return "", util.ErrEmpty
	})
	if err != nil {
		log.Println("GetEventNode session ReadTransaction err:", err)
		return "", util.ErrEmpty
	}

	log.Println("result", events)

	return events, nil
}
