package model

import (
	"github.com/asaskevich/govalidator"
	"time"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Base 			`valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank 			*Bank `valid:"-"`
	Number 		string	`json:"number" valid:"notnull"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if (err != nil) {
		return err
	}
	return nil
}

func newAccount(bank *Bank, number string, ownerName string) (*Account, error) {
	account := Account{
		OwnerName: ownerName
		Bank: bank,
		Number: number
	}

	account.ID = uuid.NewV4().String()
	account.createdAt = time.now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}