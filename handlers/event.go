package handlers

import (
	"history-graph-notes-server/dao"
)

type eventService struct{}

func (eventService) ReadNodeProper(name string) (interface{}, error) {
	result, err := dao.GetEventNodeProper(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (eventService) ReadSingleProper(name, proper string) (interface{}, error) {
	result, err := dao.GetEventSingleProper(name, proper)
	if err != nil {
		return nil, err
	}
	return result, nil
}
