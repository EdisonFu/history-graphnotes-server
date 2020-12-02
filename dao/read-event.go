package dao

import (
	. "history-graph-notes-server/model"
)

func GetEventSingleProper(name string, proper string) (string, error) {
	return "开发中", nil
}

func GetEventNodeProper(name string) (*HistoryEvent, error) {
	//driver, err := neo4j.NewDriver(Uri, neo4j.BasicAuth(Username, Password, ""), func(c *neo4j.Config) {
	//	c.Encrypted = false
	//})
	//if err != nil {
	//	log.Println("error connecting to neo4j:", err)
	//	return
	//}
	//defer driver.Close()
	//
	//session, err := driver.Session(neo4j.AccessModeRead)
	//if err != nil {
	//	log.Println("driver session err:", err)
	//	return
	//}
	//defer session.Close()
	//
	//events, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
	//	var list []*HistoryEvent
	//
	//	result, err := tx.Run("MATCH (a:Event) WHERE a.name=$name RETURN " +
	//		"a.name as name, a.aliasName as aliasName, a.ocurrTime as ocurrTime, " +
	//		"a.location as location, a.participantGroups as participantGroups," +
	//		"a.mainParticipantFigures as mainParticipantFigures, a.description as description, a.meaning as meaning",
	//		map[string]interface{}{"name": name})
	//	if err != nil {
	//		log.Println("GetEventNode tx run err:", err)
	//		return nil, err
	//	}
	//
	//	for result.Next() {
	//		event := new(HistoryEvent)
	//		record := result.Record()
	//		name, _ := record.Get("name")
	//		aliasName, _ := record.Get("aliasName")
	//		ocurrTime, _ := record.Get("ocurrTime")
	//		location, _ := record.Get("location")
	//		participantGroups, _ := record.Get("participantGroups")
	//		mainParticipantFigures, _ := record.Get("mainParticipantFigures")
	//		description, _ := record.Get("description")
	//		meaning, _ := record.Get("meaning")
	//
	//		if name !=nil {
	//			event.Name = name.(string)
	//		}
	//		if aliasName !=nil {
	//			event.AliasName = aliasName.(string)
	//		}
	//		if ocurrTime !=nil {
	//			event.OcurrTime = ocurrTime.(string)
	//		}
	//		if location !=nil {
	//			event.Location = location.(string)
	//		}
	//		if description !=nil {
	//			event.Description = description.(string)
	//		}
	//		if meaning !=nil {
	//			event.Meaning = meaning.(string)
	//		}
	//		if participantGroups !=nil {
	//			event.ParticipantGroups = util.ToStringSlice(participantGroups.([]interface{}))
	//		}
	//		if mainParticipantFigures !=nil {
	//			event.MainParticipantFigures = util.ToStringSlice(mainParticipantFigures.([]interface{}))
	//		}
	//
	//		if result.Record().GetByIndex(0) != nil{
	//			list = append(list, event)
	//		}
	//	}
	//
	//	if err = result.Err(); err != nil {
	//		log.Println("GetEventNode result err:", err)
	//		return nil, err
	//	}
	//
	//	return list, nil
	//})
	//if err != nil {
	//	log.Println("GetEventNode session ReadTransaction err:", err)
	//}
	//
	//log.Println("result", events)

	return nil, nil
}
