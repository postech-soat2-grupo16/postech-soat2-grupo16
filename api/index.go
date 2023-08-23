package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/controllers"
	"github.com/joaocampari/postech-soat2-grupo16/external"
	"github.com/joaocampari/postech-soat2-grupo16/gateways/api"
	clientegateway "github.com/joaocampari/postech-soat2-grupo16/gateways/db/cliente"
	itemgateway "github.com/joaocampari/postech-soat2-grupo16/gateways/db/item"
	pedidogateway "github.com/joaocampari/postech-soat2-grupo16/gateways/db/pedido"
	"github.com/joaocampari/postech-soat2-grupo16/usecases/cliente"
	"github.com/joaocampari/postech-soat2-grupo16/usecases/item"
	"github.com/joaocampari/postech-soat2-grupo16/usecases/pedido"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

const (
	// This is only a test token, not a real one and will be removed in the future replacing by a secret service.
	authToken = "TEST-8788837371574102-082018-c29a1c5da797dbf70a8c99b842da2850-144255706"
)

func SetupDB() *gorm.DB {
	dialector := external.GetPostgresDialector()
	db := external.NewORM(dialector)

	return db
}

func SetupRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db)

	return r
}

func mapRoutes(r *chi.Mux, orm *gorm.DB) {
	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler())

	// Injections
	// Gateways
	pedidoGateway := pedidogateway.NewGateway(orm)
	clienteGateway := clientegateway.NewGateway(orm)
	itemGateway := itemgateway.NewGateway(orm)
	mercadoPagoGateway := api.NewGateway(authToken)
	// Use cases
	itemUseCase := item.NewUseCase(itemGateway)
	pedidoUseCase := pedido.NewUseCase(pedidoGateway, mercadoPagoGateway)
	clienteUseCase := cliente.NewUseCase(clienteGateway)
	// Handlers
	_ = controllers.NewItemController(itemUseCase, r)
	_ = controllers.NewPedidoController(pedidoUseCase, r)
	_ = controllers.NewClienteController(clienteUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
