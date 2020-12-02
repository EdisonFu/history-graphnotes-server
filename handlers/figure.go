package handlers

import "history-graph-notes-server/dao"

type figureService struct{}

func (figureService) ReadNodeProper(name string) (interface{}, error) {
	result, err := dao.GetFigureNodeProper(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (figureService) ReadSingleProper(name, proper string) (interface{}, error) {
	result, err := dao.GetFigureSingleProper(name, proper)
	if err != nil {
		return nil, err
	}
	return result, nil
}
