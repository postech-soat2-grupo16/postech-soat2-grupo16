package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/cliente"
	"net/http"
	"testing"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
)

func TestGetClientes(t *testing.T) {
	t.Run("given_get_without_param_should_receive_a_list_of_clientes", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/clientes", baseURL), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status OK; got %d", res.StatusCode)
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

	t.Run("given_get_with_param_id_should_receive_a_cliente", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/clientes/1", baseURL), nil)
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

		var response domain.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (domain.Cliente{}) {
			t.Fatalf("expected a cliente; got 0")
		}
	})

	t.Run("given_get_with_invalid_id_should_receive_not_found", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/clientes/100", baseURL), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status not found; got %d", res.StatusCode)
		}
	})

	t.Run("given_post_with_body_should_create_a_cliente", func(t *testing.T) {
		newCliente := cliente.Cliente{
			CPF:   "951.254.400-86",
			Email: "test@fastfood.io",
			Nome:  "User Test",
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newCliente)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/clientes", baseURL), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("expected status created; got %d", res.StatusCode)
		}

		var response domain.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (domain.Cliente{}) {
			t.Fatalf("expected a cliente; got 0")
		}
	})

	t.Run("given_put_with_body_and_valid_id_should_update_a_cliente", func(t *testing.T) {
		updatedCliente := cliente.Cliente{
			CPF:   "951.254.400-86",
			Email: "test@fastfood.io",
			Nome:  "User Test",
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(updatedCliente)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/clientes/1", baseURL), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("expected status OK; got %d", res.StatusCode)
		}

		var response domain.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (domain.Cliente{}) {
			t.Fatalf("expected a cliente; got 0")
		}

		if response.CPF != updatedCliente.CPF && response.Email != updatedCliente.Email && response.Name != updatedCliente.Nome {
			t.Fatalf("expected a cliente updated; got 0")
		}
	})

	t.Run("given_put_with_body_and_invalid_id_should_receive_not_found", func(t *testing.T) {
		updatedCliente := cliente.Cliente{
			CPF:   "000.000.000-00",
			Email: "user@fastfood.io",
			Nome:  "User Test updated",
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(updatedCliente)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/clientes/100", baseURL), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("expected status not found; got %d", res.StatusCode)
		}
	})

	t.Run("given_delete_with_param_id_should_delete_a_cliente", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/clientes/2", baseURL), nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNoContent {
			t.Fatalf("expected status no content; got %d", res.StatusCode)
		}
	})
}
