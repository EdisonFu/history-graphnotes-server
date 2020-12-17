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
