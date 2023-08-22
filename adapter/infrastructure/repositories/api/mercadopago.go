package api

import "strconv"

type MercadoPagoAPIRepository struct{}

func NewMercadoPagoAPIRepository() *MercadoPagoAPIRepository {
	return &MercadoPagoAPIRepository{}
}

// TODO: This will be implemented in the future, to help the tests.
func (m *MercadoPagoAPIRepository) GetPedidoIDByPaymentID(paymentID string) (uint32, error) {
	pedidoID, err := strconv.ParseInt(paymentID, 10, 32)
	return uint32(pedidoID), err
}
