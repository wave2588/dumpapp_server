package handler

import "net/http"

type CheckHealthHandler struct{}

func NewCheckHealthHandler() *CheckHealthHandler {
	return &CheckHealthHandler{}
}

func (h *CheckHealthHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi get"))
}

func (h *CheckHealthHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi post"))
}

func (h *CheckHealthHandler) Put(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi put"))
}

func (h *CheckHealthHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi delete"))
}
