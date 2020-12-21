package model

const (
	Uri      = "bolt://52.3.242.73:7687"
	Username = "neo4j"
	Password = "web123456"
)

//figure
type FigureSingleProperReq struct {
	Name   string
	Proper string
}

type FigureSingleProperResp struct {
	Code int
	Msg  string
	Data string
}

type FigureNodeProperReq struct {
	Name string
}

type FigureNodeProperResp struct {
	Code int
	Msg  string
	Data *HistoryFigure
}

type HistoryFigure struct {
	Name         string
	Country      string
	Birthday     string
	Homeland     string
	Occupation   string
	Achievements []string
	Works        []string
}

//event
type EventSingleProperReq struct {
	Name   string
	Proper string
}

type EventSingleProperResp struct {
	Code int
	Msg  string
	Data string
}

type EventNodeProperReq struct {
	Name string
}

type EventNodeProperResp struct {
	Code int
	Msg  string
	Data *HistoryEvent
}

type HistoryEvent struct {
	Name                   string
	AliasName              string
	OcurrTime              string
	Location               string
	Description            string
	Meaning                string
	ParticipantGroups      []string
	MainParticipantFigures []string
}

//relation
type RelationLineReq struct {
	NameA  string
	LabelA string
	NameB  string
	LabelB string
}

type RelationLineResp struct {
	Code int
	Msg  string
	Data string
}

type NodeRelationReq struct {
	Name  string
	Label string
}

type NodeRelationResp struct {
	Code int
	Msg  string
	Data map[string]string
}

type RelationPathReq struct {
	NameA  string
	LabelA string
	NameB  string
	LabelB string
}

type RelationPathResp struct {
	Code int
	Msg  string
	Data []string
}

//add
type AddNodeReq struct {
	Label  string
	Proper map[string]string
}

type AddNodeResp struct {
	Code int
	Msg  string
}

type AddRelationReq struct {
	NameA        string
	LabelA       string
	NameB        string
	LabelB       string
	RelationType string
	Proper       map[string]string
}

type AddRelationResp struct {
	Code int
	Msg  string
}

//contribution
type SetContributeReq struct {
	UserId  string
	Content string
}

type SetContributeResp struct {
	Code int
	Msg  string
}

type GetContributeReq struct {
	UserId string
}

type GetContributeResp struct {
	Code    int
	Msg     string
	Content []string
}
