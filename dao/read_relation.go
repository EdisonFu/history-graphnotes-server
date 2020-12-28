package dao

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"history-graph-notes-server/util"
	"log"
	"strings"
)

func GetRelation(nameA, labelA, nameB, labelB string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return "", err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		cypher := `match( p1: %s {name:$nameA} )-[rel]->(p2: %s {name:$nameB} ) return rel.type`
		log.Printf("cypher:%s,\n,properMap:%v\n", fmt.Sprintf(cypher, labelA, labelB), map[string]interface{}{"nameA": nameA, "nameB": nameB})

		result, err := tx.Run(fmt.Sprintf(cypher, labelA, labelB), map[string]interface{}{"nameA": nameA, "nameB": nameB})
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

func GetRelationAndNode(name, label string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return "", err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		relaNodeMap := make(map[string]string)

		cypher := `MATCH (p: %s {name:$name} )-[r]- (a) RETURN type(r) as relation, r.year as year, a.name as name`
		log.Printf("cypher:%s,\n,properMap:%v\n", fmt.Sprintf(cypher, label), map[string]interface{}{"name": name})

		result, err := tx.Run(fmt.Sprintf(cypher, label), map[string]interface{}{"name": name})
		if err != nil {
			return nil, err
		}

		for result.Next() {
			var relation, year, name string
			record := result.Record()
			n, _ := record.Get("name")
			r, _ := record.Get("relation")
			y, _ := record.Get("year")
			if n != nil {
				switch n.(type) {
				case interface{}:
					if util.CheckInterfaceIsString(n) {
						name = n.(string)
					}
				case []interface{}:
					name = strings.Join(util.ToStringSlice(n.([]interface{})), ",")
				}
			}
			if r != nil {
				switch r.(type) {
				case interface{}:
					relation = r.(string)
				case []interface{}:
					relation = strings.Join(util.ToStringSlice(r.([]interface{})), ",")
				}
			}
			if y != nil {
				switch y.(type) {
				case interface{}:
					year = y.(string)
				case []interface{}:
					year = strings.Join(util.ToStringSlice(y.([]interface{})), ",")
				}
			}

			if _, ok := relaNodeMap[relation+":"+year]; ok {
				relaNodeMap[relation+":"+year] = relaNodeMap[relation+":"+year] + "," + name
			} else {
				relaNodeMap[relation+":"+year] = name
			}

		}

		if err = result.Err(); err != nil {
			log.Println("result err:", err)
			return nil, err
		}

		if len(relaNodeMap) > 0 {
			return relaNodeMap, nil
		}
		return "", util.ErrEmpty
	})

	if err != nil {
		log.Println("session ReadTransaction err:", err)
		return "", util.ErrEmpty
	}

	log.Println("GetRelationAndNode from neo4j result", result)
	return result, nil
}

func GetAllRelationPath(nameA, labelA, nameB, labelB string) (interface{}, error) {
	session, err := GetNeo4jConn().Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return "", err
	}
	defer session.Close()

	result, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		//cypher := `MATCH n=shortestPath((a:%s{name:$nameA})-[*]-(b:%s{name:$nameB})) return n`
		cypher := `MATCH n=(a:%s{name:$nameA})-[*]-(b:%s{name:$nameB}) return n`
		log.Printf("cypher:%s,\n,properMap:%v\n", fmt.Sprintf(cypher, labelA, labelB), map[string]interface{}{"nameA": nameA, "nameB": nameB})

		result, err := tx.Run(fmt.Sprintf(cypher, labelA, labelB), map[string]interface{}{"nameA": nameA, "nameB": nameB})
		if err != nil {
			return nil, err
		}

		var nodeRelationPath []string
		for result.Next() {
			var nodeIdMap = make(map[int64]interface{})
			var nodeRelation string

			if result.Record().GetByIndex(0) != nil {
				for _, node := range result.Record().GetByIndex(0).(neo4j.Path).Nodes() {
					nodeIdMap[node.Id()] = node.Props()["name"]
				}
				for _, relation := range result.Record().GetByIndex(0).(neo4j.Path).Relationships() {
					nodeRelation = nodeRelation + fmt.Sprintf("%v%s%v; ", nodeIdMap[relation.StartId()], relation.Type(), nodeIdMap[relation.EndId()])
				}
			}

			nodeRelationPath = append(nodeRelationPath, nodeRelation[:len(nodeRelation)-1])
		}

		if err = result.Err(); err != nil {
			log.Println("result err:", err)
			return nil, err
		}

		return nodeRelationPath, nil
	})

	if err != nil {
		log.Println("session ReadTransaction err:", err)
		return "", util.ErrEmpty
	}

	log.Println("GetFigureNodeProper from neo4j result", result)
	return result, nil
}
