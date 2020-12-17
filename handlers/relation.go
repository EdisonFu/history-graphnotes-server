package handlers

import "history-graph-notes-server/dao"

func (historyService) ReadLineRelation(nameA, lableA, nameB, lableB string) (interface{}, error) {
	result, err := dao.GetRelation(nameA, lableA, nameB, lableB)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (historyService) ReadRelationNode(name, lable string) (interface{}, error) {
	result, err := dao.GetRelationAndNode(name, lable)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (historyService) ReadAllRelationPath(nameA, lableA, nameB, lableB string) (interface{}, error) {
	result, err := dao.GetAllRelationPath(nameA, lableA, nameB, lableB)
	if err != nil {
		return nil, err
	}
	return result, nil
}
