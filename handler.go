package main

import (
	"io"
	"net/http"
)

type handlerWrapper struct{}

func (*handlerWrapper) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if handleFunction, ok := mux[request.URL.String()]; ok {
		handleFunction(writer, request)
		return
	}
	io.WriteString(writer, "Unknown path.")
}
