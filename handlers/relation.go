package handlers

import "history-graph-notes-server/dao"

func (historyService) ReadLineRelation(nameA, labelA, nameB, labelB string) (interface{}, error) {
	result, err := dao.GetRelation(nameA, labelA, nameB, labelB)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (historyService) ReadRelationNode(name, label string) (interface{}, error) {
	result, err := dao.GetRelationAndNode(name, label)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (historyService) ReadAllRelationPath(nameA, labelA, nameB, labelB string) (interface{}, error) {
	result, err := dao.GetAllRelationPath(nameA, labelA, nameB, labelB)
	if err != nil {
		return nil, err
	}
	return result, nil
}
