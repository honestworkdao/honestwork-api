package controller

import (
	"github.com/takez0o/honestwork-api/api/repository"
)

type TransactionController struct {
	Hash string
}

func NewTransactionController(hash string) *TransactionController {
	return &TransactionController{
		Hash: hash,
	}
}

func (t *TransactionController) GetTransaction() (string, error) {
	tx, err := repository.StringRead("tx:" + t.Hash)
	if err != nil {
		return "", err
	}
	return tx, nil
}

func (t *TransactionController) AddTransaction(hash string) error {
	record_id := "tx:" + t.Hash
	err := repository.StringWrite(record_id, hash, 0)
	if err != nil {
		return err
	}
	return nil
}
