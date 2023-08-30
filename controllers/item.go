package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	item2 "github.com/joaocampari/postech-soat2-grupo16/adapters/item"
	"github.com/joaocampari/postech-soat2-grupo16/interfaces"
	"github.com/joaocampari/postech-soat2-grupo16/util"

	"github.com/go-chi/chi/v5"
)

type ItemController struct {
	useCase interfaces.ItemUseCase
}

func NewItemController(useCase interfaces.ItemUseCase, r *chi.Mux) *ItemController {
	controller := ItemController{useCase: useCase}
	r.Route("/items", func(r chi.Router) {
		r.Get("/", controller.GetAll())
		r.Post("/", controller.Create())
		r.Get("/{id}", controller.GetByID())
		r.Put("/{id}", controller.Update())
		r.Delete("/{id}", controller.Delete())
	})
	return &controller
}

//	@Summary	Get all items
//
//	@Tags		Items
//
//	@ID			get-all-items
//
// @Param        category    query     string  false  "category search by category"
//
//	@Produce	json
//
// @Success	200	{object}	item2.Item
// @Failure	500
// @Router		/items [get]
func (c *ItemController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result interface{}
		var err error
		if r.URL.Query().Get("category") != "" {
			result, err = c.useCase.GetByCategory(r.URL.Query().Get("category"))
		} else {
			result, err = c.useCase.List()
		}
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(result)
	}
}

// @Summary	Get a item by ID
//
// @Tags		Items
//
// @ID			get-item-by-id
// @Produce	json
// @Param		id	path		string	true	"Item ID"
// @Success	200	{object}	item2.Item
// @Failure	404
// @Router		/items/{id} [get]
func (c *ItemController) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item, err := c.useCase.GetByID(uint32(id))
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

// @Summary	New item
//
// @Tags		Items
//
// @ID			create-item
// @Produce	json
// @Param		data	body		item2.Item	true	"Item data"
// @Success	200		{object}	item2.Item
// @Failure	400
// @Router		/items [post]
func (c *ItemController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i item2.Item
		err := json.NewDecoder(r.Body).Decode(&i)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item, err := c.useCase.Create(i.Name, i.Category, i.Description, i.Price)
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

// @Summary	Update a item
//
// @Tags		Items
//
// @ID			update-item
// @Produce	json
// @Param		id		path		string	true	"Item ID"
// @Param		data	body		item2.Item	true	"Item data"
// @Success	200		{object}	item2.Item
// @Failure	404
// @Failure	400
// @Router		/items/{id} [put]
func (c *ItemController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i item2.Item
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
		item, err := c.useCase.Update(uint32(id), i.Name, i.Category, i.Description, i.Price)
		if err != nil {
			if util.IsDomainError(err) {
				w.WriteHeader(http.StatusUnprocessableEntity)
				json.NewEncoder(w).Encode(err)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if item == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(item)
	}
}

// @Summary	Delete a item by ID
//
// @Tags		Items
//
// @ID			delete-item-by-id
// @Produce	json
// @Param		id	path	string	true	"Item ID"
// @Success	204
// @Failure	500
// @Router		/items/{id} [delete]
func (c *ItemController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item, err := c.useCase.Delete(uint32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if item == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
