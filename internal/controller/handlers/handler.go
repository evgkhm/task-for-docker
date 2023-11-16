package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"task-for-docker/internal/usecase"
)

const (
	createTime = "/api/time"
	getTime    = "/api/last_time"
)

type Handler struct {
	useCases *usecase.UseCase
}

func (h *Handler) InitRoutes() (router *chi.Mux) {
	r := chi.NewRouter()
	//turn on debug for chi
	r.Use(middleware.Logger)

	r.Post(createTime, h.createTime)
	r.Get(getTime, h.getTime)

	return r
}

func (h *Handler) createTime(writer http.ResponseWriter, request *http.Request) {
	err := h.useCases.CreateTime(request.Context())
	if err != nil {
		h.Respond(writer, nil, http.StatusInternalServerError)
		return
	}
	h.Respond(writer, nil, http.StatusCreated)
}

func (h *Handler) getTime(writer http.ResponseWriter, request *http.Request) {
	//data := []byte("Ok")
	//t := &entity.Data{}
	t, err := h.useCases.GetTime(request.Context())
	if err != nil {
		h.Respond(writer, nil, http.StatusInternalServerError)
		return
	}

	h.Respond(writer, []byte(t.Time.String()), http.StatusOK)
}

func New(useCases *usecase.UseCase) *Handler {
	return &Handler{
		useCases: useCases,
	}
}

func (h *Handler) Respond(w http.ResponseWriter, data []byte, header int) {
	w.WriteHeader(header)

	if data != nil {
		_, err := w.Write(data)
		if err != nil {
			panic(err)
		}
	}
}
