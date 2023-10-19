package usecase

import "github.com/rodrigo-orlandini/codepix-go/domain/entity"

type PixUseCase struct {
	PixKeyRepository entity.IPixKeyRepository
}

func (usecase *PixUseCase) RegisterKey(kind string, key string, accountId string) (*entity.PixKey, error) {
	account, err := usecase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := entity.NewPixKey(kind, key, account)
	if err != nil {
		return nil, err
	}

	_, err = usecase.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, err
	}

	return pixKey, nil
}

func (usecase *PixUseCase) FindKey(key string, kind string) (*entity.PixKey, error) {
	pixKey, err := usecase.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
