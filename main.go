package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/joaocampari/postech-soat2-grupo16/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"log"
	"net/http"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/cliente"

	database2 "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/database"
	clienteHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/cliente"
	itemHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/item"
	pedidoHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/pedido"
	item "github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/item"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/pedido"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

//	@title			Fast Food API
//	@version		1.0
//	@description	Here you will find everything you need to have the best possible integration with our APIs.
//	@termsOfService	http://fastfood.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.fastfood.io/support
//	@contact.email	support@fastfood.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	dialector := database2.GetPostgresDialector()
	db := database2.NewORM(dialector)

	database2.DoMigration(db)

	r := chi.NewRouter()
	r.Use(commonMiddleware)
	MapRoutes(r, db)

	log.Println(http.ListenAndServe(":8000", r))
}

func MapRoutes(r *chi.Mux, orm *gorm.DB) {
	// Injections
	// Use cases
	itemUseCase := item.NewItemUseCase(orm)
	pedidoUseCase := pedido.NewPedidoUseCase(orm)
	clienteUseCase := cliente.NewClienteUseCase(orm)

	// Handler
	r.Get("/swagger/*", httpSwagger.Handler(
	//httpSwagger.URL("/docs/swagger.json"),
	))

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
