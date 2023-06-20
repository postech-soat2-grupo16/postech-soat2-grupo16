package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
)

func TestGetClientes(t *testing.T) {
	t.Run("given_get_without_param_should_receive_a_list_of_clientes", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/clientes", baseURL), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		var response []domain.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if len(response) == 0 {
			t.Fatalf("expected a list of clientes; got 0")
		}
	})
}
