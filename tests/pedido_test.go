package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/pedido"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
)

func TestGetPedidos(t *testing.T) {
	t.Run("given_get_should_receive_a_list_of_pedidos", func(t *testing.T) {
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
			t.Fatalf("expected status OK; got %s", res.Status)
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
			t.Fatalf("expected status OK; got %s", res.Status)
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

	t.Run("given_existing_pedido_id_should_return_pedido_details", func(t *testing.T) {
		orderID := 1

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status OK; got %s", res.Status)
		}

		var response domain.Pedido
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if len(response.Items) != 2 {
			t.Fatalf("expected items length 2; got %d", len(response.Items))
		}
	})

	t.Run("given_nonexistent_pedido_id_should_return_404", func(t *testing.T) {
		orderID := 999

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status Not Found; got %s", res.Status)
		}
	})
}

func TestSavePedidos(t *testing.T) {
	t.Run("given_valid_pedido_should_create_new_pedido", func(t *testing.T) {
		newOrder := pedido.Pedido{
			Items:     []pedido.Item{{ItemID: 1, Quantity: 2}, {ItemID: 2, Quantity: 3}},
			Notes:     "Novo pedido",
			ClienteID: 1,
		}

		jsonOrder, err := json.Marshal(newOrder)
		if err != nil {
			t.Fatalf("could not marshal pedido: %v", err)
		}

		// Cria uma requisição POST com o JSON do novo pedido no corpo
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/pedidos", baseURL), bytes.NewBuffer(jsonOrder))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("expected status Created; got %s", res.Status)
		}
	})

	t.Run("given_existing_pedido_id_should_update_pedido", func(t *testing.T) {
		newNote := "Pedido atualizado"
		orderID := 1
		orderUpdated := pedido.Pedido{
			Items:     []pedido.Item{{ItemID: 1, Quantity: 5}, {ItemID: 2, Quantity: 3}},
			ClienteID: 1,
			Notes:     newNote,
		}

		jsonOrder, err := json.Marshal(orderUpdated)
		if err != nil {
			t.Fatalf("could not marshal pedido: %v", err)
		}

		req, err := http.NewRequest("PUT", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), bytes.NewBuffer(jsonOrder))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status OK; got %s", res.Status)
		}

		var response domain.Pedido
		log.Printf("%+v", response)
		err = json.NewDecoder(res.Body).Decode(&response)

		if response.Notes != newNote {
			t.Fatalf("expected notes to be %v, got %v", newNote, response.Notes)
		}
	})

	t.Run("given_nonexisting_pedido_id_should_return_404_when_updating", func(t *testing.T) {
		newNote := "Pedido atualizado"
		orderID := 999
		orderUpdated := pedido.Pedido{
			Items:     []pedido.Item{{ItemID: 1, Quantity: 5}, {ItemID: 2, Quantity: 3}},
			ClienteID: 1,
			Notes:     newNote,
		}

		jsonOrder, err := json.Marshal(orderUpdated)
		if err != nil {
			t.Fatalf("could not marshal pedido: %v", err)
		}

		req, err := http.NewRequest("PUT", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), bytes.NewBuffer(jsonOrder))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status not found; got %s", res.Status)
		}
	})

	t.Run("given_existing_pedido_id_should_delete_pedido", func(t *testing.T) {
		orderID := 1

		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNoContent {
			t.Fatalf("expected status No Content; got %s", res.Status)
		}

		req, err = http.NewRequest("GET", fmt.Sprintf("%s/pedidos/%d", baseURL, orderID), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err = http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status NOT FOUND; got %s", res.Status)
		}

	})
}
