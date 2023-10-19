package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rodrigo-orlandini/codepix-go/domain/entity"
)

type PixKeyRepositoryDatabase struct {
	Database *gorm.DB
}

func (repository PixKeyRepositoryDatabase) AddAccount(account *entity.Account) error {
	databaseError := repository.Database.Create(account).Error

	if databaseError != nil {
		return databaseError
	}

	return nil
}

func (repository PixKeyRepositoryDatabase) AddBank(bank *entity.Bank) error {
	databaseError := repository.Database.Create(bank).Error

	if databaseError != nil {
		return databaseError
	}

	return nil
}

func (repository PixKeyRepositoryDatabase) FindAccount(id string) (*entity.Account, error) {
	var account entity.Account

	repository.Database.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, errors.New("Account was not found.")
	}

	return &account, nil
}

func (repository PixKeyRepositoryDatabase) FindBank(id string) (*entity.Bank, error) {
	var bank entity.Bank

	repository.Database.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, errors.New("Bank was not found.")
	}

	return &bank, nil
}

func (repository PixKeyRepositoryDatabase) FindKeyByKind(key string, kind string) (*entity.PixKey, error) {
	var pixKey entity.PixKey

	repository.Database.Preload("Account.Bank").First(&pixKey, "key = ? AND kind = ?", key, kind)

	if pixKey.ID == "" {
		return nil, errors.New("Pix Key was not found.")
	}

	return &pixKey, nil
}

func (repository PixKeyRepositoryDatabase) RegisterKey(pixKey *entity.PixKey) (*entity.PixKey, error) {
	databaseError := repository.Database.Create(pixKey).Error

	if databaseError != nil {
		return nil, databaseError
	}

	return pixKey, nil
}
