package main

import (
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/cliente"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	database2 "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/database"
	clienteHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/cliente"
	itemHandler "github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/item"
	item "github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases/item"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func main() {
	dialector := database2.GetPostgresDialector()
	db := database2.NewORM(dialector)

	database2.DoMigration(db)

	r := chi.NewRouter()
	r.Use(commonMiddleware)
	MapRoutes(r, db)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func MapRoutes(r *chi.Mux, orm *gorm.DB) {
	// Injections
	// Use cases
	itemUseCase := item.NewItemUseCase(orm)
	clienteUseCase := cliente.NewClienteUseCase(orm)
	// Handler
	_ = itemHandler.NewHandler(itemUseCase, r)
	_ = clienteHandler.NewHandler(clienteUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
