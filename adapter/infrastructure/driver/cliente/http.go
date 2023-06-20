package cliente

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/ports"
	"github.com/joaocampari/postech-soat2-grupo16/internal/util"
)

type Handler struct {
	useCase ports.ClienteUseCase
}

func NewHandler(useCase ports.ClienteUseCase, r *chi.Mux) *Handler {
	handler := Handler{useCase: useCase}
	r.Route("/clientes", func(r chi.Router) {
		r.Get("/", handler.GetAll())
		r.Post("/", handler.Create())
		r.Get("/{id}", handler.GetById())
		r.Put("/{id}", handler.Update())
		r.Delete("/{id}", handler.Delete())
	})
	return &handler
}

// @Summary	Get all clients
// @Tags		Clients
// @ID			get-all-clients
// @Produce	json
// @Success	200	{object}	Cliente
// @Failure	500
// @Router		/clientes [get]
func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientes, err := h.useCase.List()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(clientes)
	}
}

// @Summary	Get a client by ID
//
// @Tags		Clients
//
// @ID			get-client-by-id
// @Produce	json
// @Param		id	path		string	true	"Client ID"
// @Success	200	{object}	Cliente
// @Failure	404
// @Router		/clientes/{id} [get]
func (h *Handler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cliente, err := h.useCase.GetByID(uint32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if cliente == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(cliente)
	}
}

// @Summary	New client
//
// @Tags		Clients
//
// @ID			create-client
// @Produce	json
// @Param		data	body		Cliente	true	"Client data"
// @Success	200		{object}	Cliente
// @Failure	400
// @Router		/clientes [post]
func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Cliente
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cliente, err := h.useCase.Create(i.Email, i.CPF, i.Nome)
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
		json.NewEncoder(w).Encode(cliente)
	}
}

// @Summary	Update a client
//
// @Tags		Clients
//
// @ID			update-client
// @Produce	json
// @Param		id		path		string	true	"Client ID"
// @Param		data	body		Cliente	true	"Client data"
// @Success	200		{object}	Cliente
// @Failure	404
// @Failure	400
// @Router		/clientes/{id} [put]
func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Cliente
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
		cliente, err := h.useCase.Update(uint32(id), i.Email, i.CPF, i.Nome)
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
		json.NewEncoder(w).Encode(cliente)
	}
}

// @Summary	Delete a client by ID
//
// @Tags		Clients
//
// @ID			delete-client-by-id
// @Produce	json
// @Param		id	path	string	true	"Client ID"
// @Success	204
// @Failure	500
// @Router		/clientes/{id} [delete]
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
