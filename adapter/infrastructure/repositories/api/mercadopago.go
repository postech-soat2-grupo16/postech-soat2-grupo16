package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joaocampari/postech-soat2-grupo16/internal/core/domain"
)

const (
	createQRCodeURL = "https://api.mercadopago.com/instore/orders/qr/seller/collectors/144255706/pos/FIAP/qrs"
	callbackURL     = "http://projetofuturo.com/mp/callback"
	createTitle     = "Order created"
	currencyID      = "BRL"
	unitMeasure     = "UNIT"
)

type MercadoPagoAPIRepository struct {
	AuthToken string
}

type Item struct {
	Title       string `json:"title"`
	CurrencyID  string `json:"currency_id"`
	UnitPrice   int    `json:"unit_price"`
	Quantity    int    `json:"quantity"`
	UnitMeasure string `json:"unit_measure"`
	TotalAmount int    `json:"total_amount"`
}

type RequestBody struct {
	ExternalReference string  `json:"external_reference"`
	NotificationURL   string  `json:"notification_url"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	TotalAmount       float64 `json:"total_amount"`
	Items             []Item  `json:"items"`
}

type Response struct {
	InStoreOrderID string `json:"in_store_order_id"`
	QRData         string `json:"qr_data"`
}

func NewMercadoPagoAPIRepository(authToken string) *MercadoPagoAPIRepository {
	return &MercadoPagoAPIRepository{
		authToken,
	}
}

// TODO: This will be implemented in the future, to help the tests.
func (m *MercadoPagoAPIRepository) GetPedidoIDByPaymentID(paymentID string) (uint32, error) {
	pedidoID, err := strconv.ParseInt(paymentID, 10, 32)
	return uint32(pedidoID), err
}

func (m *MercadoPagoAPIRepository) CreateQRCodeForPedido(pedido domain.Pedido) (string, error) {
	url := createQRCodeURL

	var items []Item
	for _, item := range pedido.Items {
		items = append(items, Item{
			Title:       item.Item.Name,
			CurrencyID:  currencyID,
			UnitPrice:   int(item.Item.Price),
			Quantity:    item.Quantity,
			UnitMeasure: unitMeasure,
			TotalAmount: int(item.Item.Price) * item.Quantity,
		})
	}

	requestBody := RequestBody{
		ExternalReference: strconv.Itoa(int(pedido.ID)),
		NotificationURL:   callbackURL,
		Title:             createTitle,
		Description:       createTitle,
		TotalAmount:       pedido.GetAmount(),
		Items:             items,
	}

	payload, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "application/json")
	req.Header.Set("Authorization", "Bearer "+m.AuthToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "", err
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.QRData, nil
}
