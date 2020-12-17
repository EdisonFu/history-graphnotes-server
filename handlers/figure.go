package handlers

import "history-graph-notes-server/dao"

func (historyService) ReadFigureNodeProper(name string) (interface{}, error) {
	result, err := dao.GetFigureNodeProper(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (historyService) ReadFigureSingleProper(name, proper string) (interface{}, error) {
	result, err := dao.GetFigureSingleProper(name, proper)
	if err != nil {
		return nil, err
	}
	return result, nil
}
