package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	pedido2 "github.com/joaocampari/postech-soat2-grupo16/adapters/pedido"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"github.com/joaocampari/postech-soat2-grupo16/util"

	"github.com/go-chi/chi/v5"
)

type PedidoController struct {
	useCase interfaces.PedidoUseCase
}

func NewPedidoController(useCase interfaces.PedidoUseCase, r *chi.Mux) *PedidoController {
	controller := PedidoController{useCase: useCase}
	r.Route("/pedidos", func(r chi.Router) {
		r.Get("/", controller.GetAll)
		r.Get("/{id}/pagamentos/status", controller.GetPaymentStatusByOrderID)
		r.Post("/", controller.Create)
		r.Get("/{id}", controller.GetByID)
		r.Get("/{id}/qr-code", controller.GetQRCodeByPedidoID)
		r.Put("/{id}", controller.Update)
		r.Delete("/{id}", controller.Delete)
		r.Post("/mp-webhook", controller.PaymentWebhookCreate)
	})
	return &controller
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
func (c *PedidoController) GetAll(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	pedidos, err := c.useCase.List(status)
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
func (c *PedidoController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pedido, err := c.useCase.GetByID(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if pedido == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pedido)
}

// @Summary	Get QR Code pedido
//
// @Tags		Orders
//
// @ID			get-qr-code-by-id
// @Produce	json
// @Param		id	path		string	true	"Order ID"
// @Success	200	{object}	Pedido
// @Failure	404
// @Router		/pedidos/{id}/qr-code [get]
func (c *PedidoController) GetQRCodeByPedidoID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	qrCodeStr, err := c.useCase.CreateQRCode(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if qrCodeStr == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	qrCode := pedido2.QRCode{
		QRCode: *qrCodeStr,
	}
	json.NewEncoder(w).Encode(qrCode)
}

// @Summary	Get payment status by order ID
//
// @Tags		Orders
//
// @ID			get-payment-by-order-id
// @Produce	json
// @Param		id	path		string	true	"Order ID"
// @Success	200	{object}	Pagamento
// @Failure	404
// @Router		/pedidos/{id}/pagamentos/status [get]
func (c *PedidoController) GetPaymentStatusByOrderID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pagamento, err := c.useCase.GetLastPaymentStatus(uint32(id))
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
func (c *PedidoController) Create(w http.ResponseWriter, r *http.Request) {
	var p pedido2.Pedido
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pedido, err := c.useCase.Create(p.ToDomain())
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

// @Summary	Receive payment callback from MercadoPago
//
// @Tags		Orders
//
// @ID			receive-callback
// @Produce	json
// @Param		data	body		PaymentCallback	true	"Order data"
// @Success	200		{object}	Pedido
// @Failure	400
// @Router		/pagamentos/mp-webhook [post]
func (c *PedidoController) PaymentWebhookCreate(w http.ResponseWriter, r *http.Request) {
	var payment pedido2.PaymentCallback
	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pagamento, err := c.useCase.UpdatePaymentStatusByPaymentID(payment.Data.ID)
	if err != nil {
		if util.IsDomainError(err) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(err)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if pagamento == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pagamento)
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
func (c *PedidoController) Update(w http.ResponseWriter, r *http.Request) {
	var p pedido2.Pedido
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
	pedido, err := c.useCase.Update(uint32(id), p.ToDomain())
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
func (c *PedidoController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.useCase.Delete(uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
