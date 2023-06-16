package cliente

type Cliente struct {
	CPF   string `json:"cpf"`
	Email string `json:"email"`
	Nome  string `json:"nome"`
}
