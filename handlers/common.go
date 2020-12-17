package handlers

type historyService struct{}

type HistoryService interface {
	ReadFigureSingleProper(string, string) (interface{}, error)
	ReadFigureNodeProper(string) (interface{}, error)
	ReadEventSingleProper(string, string) (interface{}, error)
	ReadEventNodeProper(string) (interface{}, error)
	ReadLineRelation(string, string, string, string) (interface{}, error)
	ReadRelationNode(string, string) (interface{}, error)
	ReadAllRelationPath(string, string, string, string) (interface{}, error)
}
