package dto

import (
	"encoding/json"
	"fmt"

	"github.com/asaskevich/govalidator"
)

type TransactionDTO struct {
	ID           string  `json:"id" validate:"required,uuid4"`
	AccountID    string  `json:"accountId" validate:"required,uuid4"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"description"`
}

func NewTransactionDTO() *TransactionDTO {
	return &TransactionDTO{}
}

func (dto *TransactionDTO) isValid() error {
	_, validationError := govalidator.ValidateStruct(dto)

	if validationError != nil {
		fmt.Errorf("Transaction DTO validation was failed: %s", validationError.Error())
		return validationError
	}

	return nil
}

func (dto *TransactionDTO) ParseJson(data []byte) error {
	parseError := json.Unmarshal(data, dto)

	if parseError != nil {
		return parseError
	}

	_, validationError := govalidator.ValidateStruct(dto)

	if validationError != nil {
		fmt.Errorf("Transaction DTO validation was failed: %s", validationError.Error())
		return validationError
	}

	return nil
}

func (dto *TransactionDTO) ToJSON() ([]byte, error) {
	err := dto.isValid()
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return data, nil
}
