package tests

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver"
	"github.com/joaocampari/postech-soat2-grupo16/tests/tutils"
)

var baseURL string

func TestMain(m *testing.M) {
	server := setup()
	defer server.Close()

	fmt.Println(m.Run())
}

func setup() *http.Server {
	os.Setenv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=fastfood_db sslmode=disable TimeZone=UTC")
	db := driver.SetupDB()
	r := driver.SetupRouter(db)

	server := http.Server{
		Handler: r,
	}
	serverAddress := tutils.StartNewTestServer(&server)
	baseURL = fmt.Sprintf("http://%s", serverAddress)

	return &server
}
