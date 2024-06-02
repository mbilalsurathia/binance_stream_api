package model

import (
	"fmt"
	"gorm.io/gorm"
)

type CurrencyPairDbModel struct {
	db *gorm.DB
}

func NewCurrencyPairModel(db *gorm.DB) CurrencyPair {
	return &CurrencyPairDbModel{db: db}
}

type (
	CurrencyPairManagement struct {
		gorm.Model
		ID           int64 `gorm:"primaryKey"`
		CurrencyPair string
		EpochTime    int64
		ExpireTime   int64
		IsActive     bool
	}

	CurrencyPair interface {
		GetAllCurrencyPairs() (*[]CurrencyPairManagement, error)
	}
)

func (u *CurrencyPairDbModel) GetAllCurrencyPairs() (*[]CurrencyPairManagement, error) {
	var currencyPairs []CurrencyPairManagement

	err := u.db.Model(&CurrencyPairManagement{}).
		Select("*").
		Find(&currencyPairs).Error

	if err != nil {
		return nil, fmt.Errorf("unable to find user balances: %w", err)
	}
	return &currencyPairs, nil
}
