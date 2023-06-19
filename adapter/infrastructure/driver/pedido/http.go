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

//	@Summary	Get all orders
//
//	@Tags		Orders
//
//	@ID			get-all-orders
//	@Produce	json
//	@Success	200	{object}	Pedido
//	@Failure	500
//	@Router		/pedidos [get]
func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pedidos, err := h.useCase.List()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(pedidos)
	}
}

//	@Summary	Get a order by ID
//
//	@Tags		Orders
//
//	@ID			get-order-by-id
//	@Produce	json
//	@Param		id	path		string	true	"Order ID"
//	@Success	200	{object}	Pedido
//	@Failure	404
//	@Router		/pedidos/{id} [get]
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

//	@Summary	New order
//
//	@Tags		Orders
//
//	@ID			create-order
//	@Produce	json
//	@Param		data	body		Pedido	true	"Order data"
//	@Success	200		{object}	Pedido
//	@Failure	400
//	@Router		/pedidos [post]
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

//	@Summary	Update a order
//
//	@Tags		Orders
//
//	@ID			update-order
//	@Produce	json
//	@Param		id		path		string	true	"Order ID"
//	@Param		data	body		Pedido	true	"Order data"
//	@Success	200		{object}	Pedido
//	@Failure	404
//	@Failure	400
//	@Router		/pedidos/{id} [put]
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

//	@Summary	Delete a order by ID
//
//	@Tags		Orders
//
//	@ID			delete-order-by-id
//	@Produce	json
//	@Param		id	path	string	true	"Order ID"
//	@Success	204
//	@Failure	500
//	@Router		/pedidos/{id} [delete]
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
