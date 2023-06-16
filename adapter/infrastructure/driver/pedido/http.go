package pedido

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joaocampari/postech-soat2-grupo16/internal/util"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
)

type Handler struct {
	useCase ports.PedidoUseCase
}

func NewHandler(useCase ports.PedidoUseCase, r *chi.Mux) *Handler {
	handler := Handler{useCase: useCase}
	r.Route("/pedidos", func(r chi.Router) {
		r.Get("/", handler.GetAll())
		r.Post("/", handler.Create())
		r.Get("/{id}", handler.GetById())
		r.Put("/{id}", handler.Update())
		r.Delete("/{id}", handler.Delete())
	})
	return &handler
}

func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pedidos, err := h.useCase.List()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(pedidos)
	}
}

func (h *Handler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pedido, err := h.useCase.GetById(uint32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if pedido == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(pedido)
	}
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Pedido
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pedido, err := h.useCase.Create(p.ToDomain())
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(pedido)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p Pedido
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		p.ID = uint32(id)
		pedido, err := h.useCase.Update(p.ToDomain())
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(pedido)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = h.useCase.Delete(uint32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
