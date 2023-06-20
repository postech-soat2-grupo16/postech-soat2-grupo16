package driver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/database"
	clienteHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/cliente"
	itemHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/item"
	pedidoHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/pedido"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/cliente"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/item"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/pedido"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
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
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/docs/swagger.json"),
	))

	// Injections
	// Use cases
	itemUseCase := item.NewItemUseCase(orm)
	pedidoUseCase := pedido.NewPedidoUseCase(orm)
	clienteUseCase := cliente.NewClienteUseCase(orm)
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
