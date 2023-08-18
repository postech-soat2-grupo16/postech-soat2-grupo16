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
		r.Get("/", handler.GetAll)
		r.Get("/{id}/pagamentos/status", handler.GetPaymentStatusByOrderID)
		r.Post("/", handler.Create)
		r.Get("/{id}", handler.GetByID)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
	return &handler
}

// @Summary	Get all orders
//
// @Tags		Orders
//
// @ID			get-all-orders
// @Produce	json
// @Param       status  query       string  false   "Optional Filter by Status"
// @Success	200	{object}	Pedido
// @Failure	500
// @Router		/pedidos [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	pedidos, err := h.useCase.List(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(pedidos)
}

// @Summary	Get a order by ID
//
// @Tags		Orders
//
// @ID			get-order-by-id
// @Produce	json
// @Param		id	path		string	true	"Order ID"
// @Success	200	{object}	Pedido
// @Failure	404
// @Router		/pedidos/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pedido, err := h.useCase.GetByID(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if pedido == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pedido)
}

// @Summary	Get payment status by order ID
//
// @Tags		Orders
//
// @ID			get-order-by-id
// @Produce	json
// @Param		id	path		string	true	"Order ID"
// @Success	200	{object}	Pedido
// @Failure	404
// @Router		/pedidos/{id} [get]
func (h *Handler) GetPaymentStatusByOrderID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pagamento, err := h.useCase.GetLastPaymentStatus(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if pagamento == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pagamento)
}

// @Summary	New order
//
// @Tags		Orders
//
// @ID			create-order
// @Produce	json
// @Param		data	body		Pedido	true	"Order data"
// @Success	200		{object}	Pedido
// @Failure	400
// @Router		/pedidos [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
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

// @Summary	Update a order
//
// @Tags		Orders
//
// @ID			update-order
// @Produce	json
// @Param		id		path		string	true	"Order ID"
// @Param		data	body		Pedido	true	"Order data"
// @Success	200		{object}	Pedido
// @Failure	404
// @Failure	400
// @Router		/pedidos/{id} [put]
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
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
	pedido, err := h.useCase.Update(uint32(id), p.ToDomain())
	if err != nil {
		if util.IsDomainError(err) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if pedido == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pedido)
}

// @Summary	Delete a order by ID
//
// @Tags		Orders
//
// @ID			delete-order-by-id
// @Produce	json
// @Param		id	path	string	true	"Order ID"
// @Success	204
// @Failure	500
// @Router		/pedidos/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
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
