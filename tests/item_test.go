package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joaocampari/postech-soat2-grupo16/adapter/infrastructure/driver/item"
	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
	"log"
	"net/http"
	"testing"
)

func TestGetItems(t *testing.T) {
	t.Run("given_empty_params_should_receive_a_list_of_items", func(t *testing.T) {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/items", baseURL), nil)
		if err != nil {
			t.Fatalf("could not create request:  %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("OK status expected; got %v", res.Status)
		}

		var response []domain.Item
		log.Printf("%+v", response)
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if len(response) == 0 {
			t.Fatalf("a list of item expected; got 0")
		}
	})

	t.Run("given_a_nonexistent_category_param_should_receive_an_unprocessable_entity_status_code", func(t *testing.T) {
		category := "category_1"
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/items?category=%s", baseURL, category), nil)
		if err != nil {
			t.Fatalf("could not create request:  %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("Unprocessable Entity Status expected; got %v", res.Status)
		}
	})

	t.Run("given_a_existing_category_param_should_receive_a_list_of_items_by_category", func(t *testing.T) {
		category := "SOBREMESA"
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/items?category=%s", baseURL, category), nil)
		if err != nil {
			t.Fatalf("could not create request:  %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("OK status expected; got %v", res.Status)
		}

		var response []domain.Item
		log.Printf("%+v", response)
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if len(response) == 0 {
			t.Fatalf("a list of item expected; got 0")
		}
	})

	t.Run("given_an_existing_item_id_param_should_receive_the_specific_item", func(t *testing.T) {
		ID := 4
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/items/%d", baseURL, ID), nil)
		if err != nil {
			t.Fatalf("could not create request:  %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("OK status expected; got %v", res.Status)
		}

		var response domain.Item
		log.Printf("%+v", response)
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (domain.Item{}) {
			t.Fatalf("a list of item expected; got 0")
		}
	})

	t.Run("given_an_existing_item_id_param_but_deleted_in_db_should_receive_the_not_found_status", func(t *testing.T) {
		ID := 7
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/items/%d", baseURL, ID), nil)
		if err != nil {
			t.Fatalf("could not create request:  %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("Not Found Status expected; got %v", res.Status)
		}
	})

	t.Run("given_a_nonexistent_item_id_param_should_receive_the_not_found_status", func(t *testing.T) {
		ID := 7654765
		req, err := http.NewRequest("GET", fmt.Sprintf("%s/items/%d", baseURL, ID), nil)
		if err != nil {
			t.Fatalf("could not create request:  %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("Not Found Status expected; got %v", res.Status)
		}
	})
}

func TestCreateItem(t *testing.T) {
	t.Run("given_a_body_should_create_an_item", func(t *testing.T) {
		newItem := item.Item{
			Name:        "teste_create_item",
			Category:    "bebida",
			Description: "teste create",
			Price:       19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/items", baseURL), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Fatalf("Created Status expected; got %d", res.StatusCode)
		}

		var response domain.Item
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (domain.Item{}) {
			t.Fatal("expected an item; got 0")
		}
	})

	t.Run("given_a_body_with_a_nonexistent_category_should_receive_an_unprocessable_entity_status", func(t *testing.T) {
		newItem := item.Item{
			Name:        "teste_create_item2",
			Category:    "categoria_nao_existinge",
			Description: "teste create",
			Price:       19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/items", baseURL), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("Unprocessable Entity Status expected; got %d", res.StatusCode)
		}
	})

	t.Run("given_a_body_with_a_nonexistent_category_should_receive_an_unprocessable_entity_status", func(t *testing.T) {
		newItem := item.Item{
			Name:        "teste_create_item2",
			Category:    "categoria_nao_existinge",
			Description: "teste create",
			Price:       19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/items", baseURL), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("Unprocessable Entity Status expected; got %d", res.StatusCode)
		}
	})

	t.Run("given_a_body_with_blank_name_field_should_receive_an_unprocessable_entity_status", func(t *testing.T) {
		newItem := item.Item{
			Name:        "",
			Category:    "ACOMPANHAMENTO",
			Description: "teste create",
			Price:       19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/items", baseURL), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("Unprocessable Entity Status expected; got %d", res.StatusCode)
		}
	})

	t.Run("given_a_body_with_invalid_price_field_should_receive_an_unprocessable_entity_status", func(t *testing.T) {
		newItem := item.Item{
			Name:        "name_invalid_price",
			Category:    "BEBIDA",
			Description: "teste create",
			Price:       -19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/items", baseURL), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("Unprocessable Entity Status expected; got %d", res.StatusCode)
		}
	})
}

func TestUpdateItem(t *testing.T) {
	t.Run("given_a_body_and_an_existing_item_id_should_receive_an_ok_status", func(t *testing.T) {
		ID := 3
		newItem := item.Item{
			Name:        "teste_update_item",
			Category:    "bebida",
			Description: "teste update",
			Price:       19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/items/%d", baseURL, ID), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			t.Fatalf("OK status expected; got %d", res.StatusCode)
		}

		var response domain.Item
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("could not parse response: %v", err)
		}

		if response == (domain.Item{}) {
			t.Fatal("expected an item; got 0")
		}
	})

	t.Run("given_a_body_with_a_nonexistent_category_should_receive_an_unprocessable_entity_status", func(t *testing.T) {
		ID := 3
		newItem := item.Item{
			Name:        "teste_update_item2",
			Category:    "categoria_nao_existente",
			Description: "teste update",
			Price:       19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/items/%d", baseURL, ID), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("Unprocessable Entity Status expected; got %d", res.StatusCode)
		}
	})

	t.Run("given_a_body_with_empty_name_field_should_receive_an_unprocessable_entity_status", func(t *testing.T) {
		ID := 3
		newItem := item.Item{
			Name:        "",
			Category:    "bebida",
			Description: "teste update",
			Price:       19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/items/%d", baseURL, ID), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("Unprocessable Entity Status expected; got %d", res.StatusCode)
		}
	})

	t.Run("given_a_body_with_invalid_price_field_should_receive_an_unprocessable_entity_status", func(t *testing.T) {
		ID := 3
		newItem := item.Item{
			Name:        "",
			Category:    "bebida",
			Description: "teste update",
			Price:       -19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/items/%d", baseURL, ID), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusUnprocessableEntity {
			t.Fatalf("Unprocessable Entity Status expected; got %d", res.StatusCode)
		}
	})

	t.Run("given_a_body_with_a_nonexistent_item_id_should_receive_a_not_found_status", func(t *testing.T) {
		ID := 36444
		newItem := item.Item{
			Name:        "teste_update_item3",
			Category:    "bebida",
			Description: "teste update",
			Price:       19.43,
		}

		body := new(bytes.Buffer)
		err := json.NewEncoder(body).Encode(newItem)
		if err != nil {
			t.Fatalf("could not convert to json: %v", err)
		}

		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/items/%d", baseURL, ID), body)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("Unprocessable Entity Status expected; got %d", res.StatusCode)
		}
	})
}

func TestDeleteItem(t *testing.T) {
	t.Run("given_an_existing_item_id_should_receive_no_content_status", func(t *testing.T) {
		ID := 6
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/items/%d", baseURL, ID), nil)
		if err != nil {
			t.Fatalf("could not create request:  %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNoContent {
			t.Fatalf("No Content Status expected; got %v", res.Status)
		}
	})

	t.Run("given_an_nonexistent_item_id_should_receive_not_found_status", func(t *testing.T) {
		ID := 65645
		req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/items/%d", baseURL, ID), nil)
		if err != nil {
			t.Fatalf("could not create request:  %v", err)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatalf("could not send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusNotFound {
			t.Fatalf("OK status expected; got %v", res.Status)
		}
	})
}
