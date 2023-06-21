package tests

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
)

func TestGetPedidos(t *testing.T) {
	t.Run("given_get_without_param_should_receive_a_list_of_pedidos", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/pedidos", baseURL), nil)
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

		var response []domain.Pedido
		log.Printf("%+v", response)
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if len(response) == 0 {
			t.Fatalf("expected a list of clientes; got 0")
		}
	})

	t.Run("given_get_with_status_query_param_should_receive_only_pedidos_with_same_status_value", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/pedidos", baseURL), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		q := req.URL.Query()
		q.Add("status", "AGUARDANDO_PAGAMENTO")
		req.URL.RawQuery = q.Encode()

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status OK; got %v", res.Status)
		}

		var response []domain.Pedido
		log.Printf("%+v", response)
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if len(response) == 0 {
			t.Fatalf("expected a list of pedidos; got 0")
		}

		for _, pedido := range response {
			if pedido.Status != "AGUARDANDO_PAGAMENTO" {
				t.Fatalf("pedido %d, expected status %s; got status %s", pedido.ID, "AGUARDANDO_PAGAMENTO", pedido.Status)
			}
		}
	})
}
