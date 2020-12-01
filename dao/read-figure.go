package dao

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	."history-graph-notes-server/model"
	"history-graph-notes-server/util"
	"log"
)

func GetFigureProper(){
	driver, err := neo4j.NewDriver(Uri, neo4j.BasicAuth(Username, Password, ""), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		log.Println("error connecting to neo4j:", err)
		return
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return
	}
	defer session.Close()

	people, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []string

		result, err := tx.Run("MATCH (a:Person) RETURN a.name", nil)
		if err != nil {
			return nil, err
		}

		for result.Next() {
			if result.Record().GetByIndex(0) != nil{
				list = append(list, result.Record().GetByIndex(0).(string))
			}
		}

		if err = result.Err(); err != nil {
			log.Println("result err:", err)
			return nil, err
		}

		return list, nil
	})
	if err != nil {
		log.Println("session ReadTransaction err:", err)
	}

	log.Println("result", people)
}

func GetFigureNode(name string){
	driver, err := neo4j.NewDriver(Uri, neo4j.BasicAuth(Username, Password, ""), func(c *neo4j.Config) {
		c.Encrypted = false
	})
	if err != nil {
		log.Println("error connecting to neo4j:", err)
		return
	}
	defer driver.Close()

	session, err := driver.Session(neo4j.AccessModeRead)
	if err != nil {
		log.Println("driver session err:", err)
		return
	}
	defer session.Close()

	figures, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var list []*HistoryFigure

		result, err := tx.Run("MATCH (a:Person) WHERE a.name=$name RETURN " +
			"a.name as name, a.country as country, a.birthday as birthday, " +
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

			if name !=nil {
				figure.Name = name.(string)
			}
			if country !=nil {
				figure.Country = country.(string)
			}
			if birthday !=nil {
				figure.Birthday = birthday.(string)
			}
			if homeland !=nil {
				figure.Homeland = homeland.(string)
			}
			if occupation !=nil {
				figure.Occupation = occupation.(string)
			}
			if achievements !=nil {
				figure.Achievements = util.ToStringSlice(achievements.([]interface{}))
			}
			if works !=nil {
				figure.Works = util.ToStringSlice(works.([]interface{}))
			}

			if result.Record().GetByIndex(0) != nil{
				list = append(list, figure)
			}
		}

		if err = result.Err(); err != nil {
			log.Println("GetFigureNode result err:", err)
			return nil, err
		}

		return list, nil
	})
	if err != nil {
		log.Println("GetFigureNode session ReadTransaction err:", err)
	}

	log.Println("result", figures)
}
