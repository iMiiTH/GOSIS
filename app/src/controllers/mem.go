package controllers

import (
	"io"
	"net/http"
)

func Mem(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Mem response.")
}
