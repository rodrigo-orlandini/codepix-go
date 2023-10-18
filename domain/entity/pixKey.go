package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
	Base      `valid:"require"`
	Kind      string `json:"kind" valid:"notnull"`
	Key       string `json:"key" valid:"notnull"`
	AccountID string `json:"account_id" valid:"notnull"`
	Account   *Account
	Status    string `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, validationError := govalidator.ValidateStruct(pixKey)

	if validationError != nil {
		return validationError
	}

	return nil
}

func NewPixKey(kind string, key string, account *Account) (*PixKey, error) {
	pixKey := PixKey{
		Kind:    kind,
		Key:     key,
		Account: account,
		Status:  "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	validationError := pixKey.isValid()

	if validationError != nil {
		return nil, validationError
	}

	return &pixKey, nil
}
