package handlers

import (
	"context"
	"encoding/json"
	. "history-graph-notes-server/model"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeFigureSingleProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FigureSingleProperReq)
		v, err := svc.ReadSingleProper(req.Name, req.Proper)
		if err != nil || v == nil {
			return FigureSingleProperResp{-1, err.Error(), ""}, nil
		}
		return FigureSingleProperResp{0, "", v.(string)}, nil
	}
}

func MakeFigureNodeProperEndpoint(svc HistoryService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(FigureNodeProperReq)
		v, err := svc.ReadNodeProper(req.Name)
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

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
