package handlers

import (
	"history-graph-notes-server/dao"
)

func (historyService) ReadEventNodeProper(name string) (interface{}, error) {
	result, err := dao.GetEventNodeProper(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (historyService) ReadEventSingleProper(name, proper string) (interface{}, error) {
	result, err := dao.GetEventSingleProper(name, proper)
	if err != nil {
		return nil, err
	}
	return result, nil
}
