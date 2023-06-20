package main

import (
	"log"
	"net/http"

	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver"
	_ "github.com/joaocampari/postech-soat2-grupo16/docs"
	_ "github.com/lib/pq"
)

//	@title			Fast Food API
//	@version		1.0
//	@description	Here you will find everything you need to have the best possible integration with our APIs.
//	@termsOfService	http://fastfood.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.fastfood.io/support
//	@contact.email	support@fastfood.io

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	db := driver.SetupDB()
	r := driver.SetupRouter(db)

	log.Println(http.ListenAndServe(":8000", r))
}
