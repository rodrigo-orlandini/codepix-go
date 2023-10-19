package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" valid:"notnull"`
	Name     string     `json:"name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (bank *Bank) isValid() error {
	_, validationError := govalidator.ValidateStruct(bank)

	if validationError != nil {
		return validationError
	}

	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	validationError := bank.isValid()

	if validationError != nil {
		return nil, validationError
	}

	return &bank, nil
}
