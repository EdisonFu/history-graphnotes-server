package dao

import (
	"fmt"
	"testing"
)

func TestGetRelationAndNode(t *testing.T) {
	InitNeo4jDB()
	defer CloseNeo4j()
	result, err := GetRelationAndNode("艾萨克·牛顿", "Person")
	if err != nil {
		fmt.Printf("GetRelationAndNode err:%v", err)
	}
	fmt.Printf("GetRelationAndNode result:%v", result)
}

func TestGetShortestPath(t *testing.T) {
	InitNeo4jDB()
	defer CloseNeo4j()
	result, err := GetAllRelationPath("艾萨克·牛顿", "Person", "高斯", "Person")
	if err != nil {
		fmt.Printf("TestGetShortestPath err:%v", err)
	}
	fmt.Printf("TestGetShortestPath result:%v", result)
}

func TestAddNode(t *testing.T) {
	InitNeo4jDB()
	defer CloseNeo4j()
	err := AddNode("Person", map[string]string{"name": "杜甫", "country": "中国"})
	if err != nil {
		fmt.Printf("TestAddNode err:%v", err)
	}
	fmt.Println("TestAddNode ok")
}

func TestAddNodeRelation(t *testing.T) {
	InitNeo4jDB()
	defer CloseNeo4j()
	err := AddNodeRelation("杜甫", "Person", "李白", "Person", "KNOWS", map[string]string{"year": "1000"})
	if err != nil {
		fmt.Printf("TestGetShortestPath err:%v", err)
	}
	fmt.Println("TestAddNodeRelation ok")
}

func TestGetNodeList(t *testing.T) {
	InitNeo4jDB()
	defer CloseNeo4j()
	result, err := GetNodeList()
	if err != nil {
		fmt.Printf("TestGetNodeList err:%v", err)
	}
	fmt.Printf("TestGetNodeList result:%v", result)
}

func TestGetFitNode(t *testing.T) {
	InitNeo4jDB()
	defer CloseNeo4j()
	result, err := GetFitNode("Aerson", "name", "A")
	if err != nil {
		fmt.Printf("GetFitNode err:%v", err)
	}
	fmt.Printf("GetFitNode result:%v", result)
}
