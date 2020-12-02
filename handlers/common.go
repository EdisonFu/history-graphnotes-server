package handlers

type HistoryService interface {
	ReadSingleProper(string, string) (interface{}, error)
	ReadNodeProper(string) (interface{}, error)
}
