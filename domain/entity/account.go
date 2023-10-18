package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base      `valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank      *Bank  `valid:"-"`
	Number    string `json:"number" valid:"notnull"`
}

func (account *Account) isValid() error {
	_, validationError := govalidator.ValidateStruct(account)

	if validationError != nil {
		return validationError
	}

	return nil
}

func NewAccount(bank *Bank, ownerName string, number string) (*Account, error) {
	account := Account{
		Bank:      bank,
		OwnerName: ownerName,
		Number:    number,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	validationError := account.isValid()

	if validationError != nil {
		return nil, validationError
	}

	return account, nil
}
