package usecase

import "github.com/rodrigo-orlandini/codepix-go/domain/entity"

type TransactionUseCase struct {
	TransactionRepository entity.ITransactionRepository
	PixKeyRepository      entity.IPixKeyRepository
}

func (usecase *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKind string, description string) (*entity.Transaction, error) {
	account, err := usecase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := usecase.PixKeyRepository.FindKeyByKind(pixKeyTo, pixKeyKind)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (usecase *TransactionUseCase) Confirm(transactionId string) (*entity.Transaction, error) {
	transaction, err := usecase.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = entity.TransactionConfirmed

	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (usecase *TransactionUseCase) Complete(transactionId string) (*entity.Transaction, error) {
	transaction, err := usecase.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = entity.TransactionCompleted

	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (usecase *TransactionUseCase) Error(transactionId string, reason string) (*entity.Transaction, error) {
	transaction, err := usecase.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = entity.TransactionError
	transaction.CancelDescription = reason

	err = usecase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
