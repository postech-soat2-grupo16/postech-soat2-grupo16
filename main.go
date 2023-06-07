package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joaocampari/postech-soat2-grupo16/infrastructure/database"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/usecases"
	"github.com/joaocampari/postech-soat2-grupo16/internal/handlers/product"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func main() {
	dialector := database.GetPostgresDialector()
	db := database.NewORM(dialector)

	database.DoMigration(db)

	r := chi.NewRouter()
	r.Use(commonMiddleware)
	MapRoutes(r, db)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func MapRoutes(r *chi.Mux, orm *gorm.DB) {
	// Injections
	// Use cases
	productUseCase := usecases.NewProductUseCase(orm)
	// Handler
	_ = product.NewHandler(productUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
