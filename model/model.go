package model

const(
	Uri = "bolt://52.3.242.73:7687"
	Username = "neo4j"
	Password = "web123456"
)

type HistoryFigure struct{
	Name string
	Country string
	Birthday string
	Homeland string
	Occupation string
	Achievements []string
	Works []string
}

type HistoryEvent struct{
	Name string
	AliasName string
	OcurrTime string
	Location string
	Description string
	Meaning string
	ParticipantGroups []string
	MainParticipantFigures []string
}