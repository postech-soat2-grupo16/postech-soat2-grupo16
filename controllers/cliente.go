package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	cliente2 "github.com/joaocampari/postech-soat2-grupo16/adapters/cliente"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"github.com/joaocampari/postech-soat2-grupo16/util"
)

type ClienteController struct {
	useCase interfaces.ClienteUseCase
}

func NewClienteController(useCase interfaces.ClienteUseCase, r *chi.Mux) *ClienteController {
	controller := ClienteController{useCase: useCase}
	r.Route("/clientes", func(r chi.Router) {
		r.Get("/", controller.GetAll())
		r.Post("/", controller.Create())
		r.Get("/{id}", controller.GetByID())
		r.Put("/{id}", controller.Update())
		r.Delete("/{id}", controller.Delete())
	})
	return &controller
}

// @Summary	Get all clients
// @Tags		Clients
// @ID			get-all-clients
// @Produce	json
// @Success	200	{object}	cliente2.Cliente
// @Param       cpf  query       string  false   "Optional Filter by CPF"
// @Failure	500
// @Router		/clientes [get]
func (c *ClienteController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CPF := r.URL.Query().Get("cpf")
		clientes, err := c.useCase.List(CPF)
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
// @Success	200	{object}	cliente2.Cliente
// @Failure	404
// @Router		/clientes/{id} [get]
func (c *ClienteController) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cliente, err := c.useCase.GetByID(uint32(id))
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
// @Param		data	body		cliente2.Cliente	true	"Client data"
// @Success	200		{object}	cliente2.Cliente
// @Failure	400
// @Router		/clientes [post]
func (c *ClienteController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i cliente2.Cliente
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		cliente, err := c.useCase.Create(i.Email, i.CPF, i.Nome)
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
// @Param		data	body		cliente2.Cliente	true	"Client data"
// @Success	200		{object}	cliente2.Cliente
// @Failure	404
// @Failure	400
// @Router		/clientes/{id} [put]
func (c *ClienteController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i cliente2.Cliente
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

		existingCliente, err := c.useCase.GetByID(uint32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if existingCliente == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		cliente, err := c.useCase.Update(uint32(id), i.Email, i.CPF, i.Nome)
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
func (c *ClienteController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
