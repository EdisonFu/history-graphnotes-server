package handlers

import "history-graph-notes-server/dao"

func (historyService) AddNode(label string, proper map[string]string) error {
	err := dao.AddNode(label, proper)
	if err != nil {
		return err
	}
	return nil
}

func (historyService) AddNodeRelation(nameA, labelA, nameB, labelB, relationType string, proper map[string]string) error {
	err := dao.AddNodeRelation(nameA, labelA, nameB, labelB, relationType, proper)
	if err != nil {
		return err
	}
	return nil
}
