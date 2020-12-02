package handlers

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func Init() {
	figuresvc := figureService{}

	figureSingleProperHandler := httptransport.NewServer(
		MakeFigureSingleProperEndpoint(figuresvc),
		decodeFigureSingleProperRequest,
		encodeResponse,
	)

	figureNodeProperHandler := httptransport.NewServer(
		MakeFigureNodeProperEndpoint(figuresvc),
		decodeFigureNodeProperRequest,
		encodeResponse,
	)

	http.Handle("/figure/proper", figureSingleProperHandler)
	http.Handle("/figure/node", figureNodeProperHandler)
}
