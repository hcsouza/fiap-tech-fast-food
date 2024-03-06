package dto

type CustomerCreateDTO struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
	Cpf   string `json:"cpf" binding:"required,IsCpfValid" `
}
