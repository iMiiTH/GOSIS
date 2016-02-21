package controllers

import (
	"io"
	"net/http"
)

func Cpu(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cpu response.")
}
