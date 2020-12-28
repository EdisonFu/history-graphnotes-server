package handlers

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func Init() {
	historysvc := historyService{}

	figureSingleProperHandler := httptransport.NewServer(
		MakeFigureSingleProperEndpoint(historysvc),
		decodeFigureSingleProperRequest,
		encodeResponse,
	)

	figureNodeProperHandler := httptransport.NewServer(
		MakeFigureNodeProperEndpoint(historysvc),
		decodeFigureNodeProperRequest,
		encodeResponse,
	)

	eventSingleProperHandler := httptransport.NewServer(
		MakeEventSingleProperEndpoint(historysvc),
		decodeEventSingleProperRequest,
		encodeResponse,
	)

	eventNodeProperHandler := httptransport.NewServer(
		MakeEventNodeProperEndpoint(historysvc),
		decodeEventNodeProperRequest,
		encodeResponse,
	)

	relationLineHandler := httptransport.NewServer(
		MakeRelationLineEndpoint(historysvc),
		decodeRelationLineRequest,
		encodeResponse,
	)

	relationNodeHandler := httptransport.NewServer(
		MakeRelationNodeEndpoint(historysvc),
		decodeRelationNodeRequest,
		encodeResponse,
	)

	relationPathHandler := httptransport.NewServer(
		MakeRelationPathEndpoint(historysvc),
		decodeRelationPathRequest,
		encodeResponse,
	)

	addNodeHandler := httptransport.NewServer(
		MakeAddNodeEndpoint(historysvc),
		decodeAddNodeRequest,
		encodeResponse,
	)

	addRelationHandler := httptransport.NewServer(
		MakeAddNodeRelationEndpoint(historysvc),
		decodeAddNodeRelationRequest,
		encodeResponse,
	)

	setContributeHandler := httptransport.NewServer(
		MakeSetContributeEndpoint(historysvc),
		decodeSetContributeRequest,
		encodeResponse,
	)

	getContributeHandler := httptransport.NewServer(
		MakeGetContributeEndpoint(historysvc),
		decodeGetContributeRequest,
		encodeResponse,
	)

	readNodeListHandler := httptransport.NewServer(
		MakeReadNodeListEndpoint(historysvc),
		decodeReadNodeListRequest,
		encodeResponse,
	)

	readFitNodeHandler := httptransport.NewServer(
		MakeReadFitNodeEndpoint(historysvc),
		decodeReadFitNodeRequest,
		encodeResponse,
	)

	http.Handle("/figure/proper", figureSingleProperHandler)
	http.Handle("/figure/node", figureNodeProperHandler)
	http.Handle("/event/proper", eventSingleProperHandler)
	http.Handle("/event/node", eventNodeProperHandler)
	http.Handle("/relation/line", relationLineHandler)
	http.Handle("/relation/node", relationNodeHandler)
	http.Handle("/relation/path", relationPathHandler)
	http.Handle("/add/node", addNodeHandler)
	http.Handle("/add/relation", addRelationHandler)
	http.Handle("/contribute/set", setContributeHandler)
	http.Handle("/contribute/get", getContributeHandler)
	http.Handle("/node/list", readNodeListHandler)
	http.Handle("/node/fit", readFitNodeHandler)
}
