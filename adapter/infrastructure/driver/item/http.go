package item

import (
	"encoding/json"
	"github.com/joaocampari/postech-soat2-grupo16/internal/util"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
)

type Handler struct {
	useCase ports.ItemUseCase
}

func NewHandler(useCase ports.ItemUseCase, r *chi.Mux) *Handler {
	handler := Handler{useCase: useCase}
	r.Route("/items", func(r chi.Router) {
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
		items, err := h.useCase.List()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(items)
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
		item, err := h.useCase.GetById(uint32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if item == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(item)
	}
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item, err := h.useCase.Create(i.Name, i.Category, i.Description, i.Price)
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
		json.NewEncoder(w).Encode(item)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		err := json.NewDecoder(r.Body).Decode(&i)
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
		item, err := h.useCase.Update(uint32(id), i.Name, i.Category, i.Description, i.Price)
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
		json.NewEncoder(w).Encode(item)
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
