package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	. "history-graph-notes-server/model"
	"net/http"
)

//figure handler
func MakeFigureSingleProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FigureSingleProperReq)
		v, err := svc.ReadFigureSingleProper(req.Name, req.Proper)
		if err != nil || v == nil {
			return FigureSingleProperResp{-1, err.Error(), ""}, nil
		}
		return FigureSingleProperResp{0, "", v.(string)}, nil
	}
}

func MakeFigureNodeProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FigureNodeProperReq)
		v, err := svc.ReadFigureNodeProper(req.Name)
		if err != nil || v == nil {
			return FigureNodeProperResp{-1, err.Error(), nil}, nil
		}
		return FigureNodeProperResp{0, "", v.(*HistoryFigure)}, nil
	}
}

func decodeFigureSingleProperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FigureSingleProperReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeFigureNodeProperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FigureNodeProperReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//event handler
func MakeEventSingleProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(EventSingleProperReq)
		v, err := svc.ReadEventSingleProper(req.Name, req.Proper)
		if err != nil || v == nil {
			return EventSingleProperResp{-1, err.Error(), ""}, nil
		}
		return EventSingleProperResp{0, "", v.(string)}, nil
	}
}

func MakeEventNodeProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(EventNodeProperReq)
		v, err := svc.ReadEventNodeProper(req.Name)
		if err != nil || v == nil {
			return EventNodeProperResp{-1, err.Error(), nil}, nil
		}
		return EventNodeProperResp{0, "", v.(*HistoryEvent)}, nil
	}
}

func decodeEventSingleProperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request EventSingleProperReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeEventNodeProperRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request EventNodeProperReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//relation handler
func MakeRelationLineEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(RelationLineReq)
		v, err := svc.ReadLineRelation(req.NameA, req.LabelA, req.NameB, req.LabelB)
		if err != nil || v == nil {
			return RelationLineResp{-1, err.Error(), ""}, nil
		}
		return RelationLineResp{0, "", v.(string)}, nil
	}
}

func decodeRelationLineRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request RelationLineReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeRelationNodeEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(NodeRelationReq)
		v, err := svc.ReadRelationNode(req.Name, req.Label)
		if err != nil || v == nil {
			return NodeRelationResp{-1, err.Error(), nil}, nil
		}
		return NodeRelationResp{0, "", v.(map[string]string)}, nil
	}
}

func decodeRelationNodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request NodeRelationReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeRelationPathEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(RelationPathReq)
		v, err := svc.ReadAllRelationPath(req.NameA, req.LabelA, req.NameB, req.LabelB)
		if err != nil || v == nil {
			return RelationPathResp{-1, err.Error(), nil}, nil
		}
		return RelationPathResp{0, "", v.([]string)}, nil
	}
}

func decodeRelationPathRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request RelationPathReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//add handler
func MakeAddNodeEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNodeReq)
		err := svc.AddNode(req.Label, req.Proper)
		if err != nil {
			return AddNodeResp{-1, err.Error()}, nil
		}
		return AddNodeResp{0, ""}, nil
	}
}

func decodeAddNodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AddNodeReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeAddNodeRelationEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(AddRelationReq)
		err := svc.AddNodeRelation(req.NameA, req.LabelA, req.NameB, req.LabelB, req.RelationType, req.Proper)
		if err != nil {
			return AddRelationResp{-1, err.Error()}, nil
		}
		return AddRelationResp{0, ""}, nil
	}
}

func decodeAddNodeRelationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AddRelationReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//contribute handler
func MakeSetContributeEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(SetContributeReq)
		svc.SetUserContribute(req.UserId, req.Content)
		return SetContributeResp{0, ""}, nil
	}
}

func decodeSetContributeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SetContributeReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func MakeGetContributeEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetContributeReq)
		data := svc.GetUserContribute(req.UserId)
		if data == nil {
			return GetContributeResp{0, "data is null!", nil}, nil
		}
		return GetContributeResp{0, "", data.([]*ContributeTime)}, nil
	}
}

func decodeGetContributeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetContributeReq
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return json.NewEncoder(w).Encode(response)
}
