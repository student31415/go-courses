package bank

import (
	"fmt"

	"github.com/CodingSquire/go-courses/lection_6/hw/bank/pkg/money"
)

type bank struct {
	data map[int]money.Money
}

func (b bank) CreateAcc(id int) (err error) {
	if _, err := b.GetBalance(id); err != nil {
		err = fmt.Errorf("account reserved")
	} else {
		b.data[id] = 0
	}
	return
}

func (b bank) AddToBell(id int, data money.Money) (err error) {
	balance, err := b.GetBalance(id)
	if err != nil {
		return err
	}
	b.SetBalance(id, balance+data)
	return
}

func (b bank) DeductFromBell(id int, data money.Money) (err error) {
	balance, err := b.GetBalance(id)
	if err != nil {
		return err
	}
	if err = money.Validate(balance - data); err != nil {
		return err
	}
	b.SetBalance(id, balance-data)
	return
}

func (b bank) GetBalance(id int) (data money.Money, err error) {
	if v, ok := b.data[id]; ok {
		return v, nil
	} else {
		err = fmt.Errorf("no account data")
		return
	}
}

func (b bank) SetBalance(id int, data money.Money) {
	b.data[id] = data
}

type Bank interface {
	GetBalance(id int) (data money.Money, err error)
	SetBalance(id int, data money.Money)
	AddToBell(id int, data money.Money) (err error)
	DeductFromBell(id int, data money.Money) (err error)
	CreateAcc(id int) error
}

func NewBank() Bank {
	return &bank{
		data: make(map[int]money.Money),
	}
}
