package driver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/database"
	clienteHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/cliente"
	itemHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/item"
	pedidoHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/pedido"
	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/repositories/api"
	clienterepo "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/repositories/db/cliente"
	itemrepo "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/repositories/db/item"
	pedidorepo "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/repositories/db/pedido"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/cliente"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/item"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/pedido"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

const (
	// This is only a test token, not a real one and will be remove in the future replacing by a secret service.
	authToken = "TEST-8788837371574102-082018-c29a1c5da797dbf70a8c99b842da2850-144255706"
)

func SetupDB() *gorm.DB {
	dialector := database.GetPostgresDialector()
	db := database.NewORM(dialector)

	return db
}

func SetupRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db)

	return r
}

func mapRoutes(r *chi.Mux, orm *gorm.DB) {
	// Handler
	r.Get("/swagger/*", httpSwagger.Handler())

	// Injections
	// Repositories
	pedidoRepository := pedidorepo.NewRepository(orm)
	clienteRepository := clienterepo.NewRepository(orm)
	itemRepository := itemrepo.NewRepository(orm)
	mercadoPagoRepository := api.NewMercadoPagoAPIRepository(authToken)
	// Use cases
	itemUseCase := item.NewUseCase(itemRepository)
	pedidoUseCase := pedido.NewUseCase(pedidoRepository, mercadoPagoRepository)
	clienteUseCase := cliente.NewUseCase(clienteRepository)
	// Handlers
	_ = itemHandler.NewHandler(itemUseCase, r)
	_ = pedidoHandler.NewHandler(pedidoUseCase, r)
	_ = clienteHandler.NewHandler(clienteUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
