package model

import (
	"github.com/asaskevich/govalidator"
	"time"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base `valid:"required"`
	Code string `json:"code" valid:"notnull"`
	Name string `json:"name" valid:"notnull"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)
	if (err != nil) {
		return err
	}
	return nil
}

func newBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name
	}

	bank.ID = uuid.NewV4().String()
	bank.createdAt = time.now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}