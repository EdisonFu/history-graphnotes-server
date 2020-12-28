package handlers

import "history-graph-notes-server/dao"

func (historyService) ReadNodeList() (interface{}, error) {
	result, err := dao.GetNodeList()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (historyService) ReadFitNode(label, proper, key string) (interface{}, error) {
	result, err := dao.GetFitNode(label, proper, key)
	if err != nil {
		return nil, err
	}
	return result, nil
}
