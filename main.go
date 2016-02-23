package main

import (
	"gitlab.com/iMiiTH/GOSIS/controllers"
	"net/http"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &handlerWrapper{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/cpu"] = controllers.Cpu
	mux["/mem"] = controllers.Mem

	server.ListenAndServe()
}
