package handlers

import "history-graph-notes-server/dao"

type relationService struct{}

func (relationService) ReadNodeProper(name string) (interface{}, error) {
	result, err := dao.GetRelationNodeProper(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (relationService) ReadSingleProper(name, proper string) (interface{}, error) {
	result, err := dao.GetRelationSingleProper(name, proper)
	if err != nil {
		return nil, err
	}
	return result, nil
}
