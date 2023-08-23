package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/joaocampari/postech-soat2-grupo16/adapters/cliente"
	"github.com/joaocampari/postech-soat2-grupo16/entities"
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

		var response []entities.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if len(response) == 0 {
			t.Fatal("expected a list of clientes; got 0")
		}
	})

	t.Run("given_get_with_param_cpf_should_receive_a_cliente", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/clientes?cpf=12312312312", baseURL), nil)
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

		var response []entities.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if len(response) > 1 {
			t.Fatal("expected a list of clientes; got 0")
		}
		expectedName := "cliente teste 1"
		if response[0].Name != expectedName {
			t.Fatalf("expected name %s; got %s", expectedName, response[0].Name)
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

		var response entities.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (entities.Cliente{}) {
			t.Fatal("expected a cliente; got 0")
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

		var response entities.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (entities.Cliente{}) {
			t.Fatal("expected a cliente; got 0")
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

		var response entities.Cliente
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (entities.Cliente{}) {
			t.Fatal("expected a cliente; got 0")
		}

		if response.CPF != updatedCliente.CPF {
			t.Fatalf("expected a cliente CPF %s; got %s", updatedCliente.CPF, response.CPF)
		}

		if response.Email != updatedCliente.Email {
			t.Fatalf("expected a cliente Email %s; got %s", updatedCliente.Email, response.Email)
		}

		if response.Name != updatedCliente.Nome {
			t.Fatalf("expected a cliente Name %s; got %s", updatedCliente.Nome, response.Name)
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
