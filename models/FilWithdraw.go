package models

import (
	"github.com/jinzhu/gorm"
)

type FilWithdraw struct {
	ID          int     `gorm:"primary_key" json:"id"`
	BusinessNo  string  `json:"business_no"`
	Txid        string  `json:"txid"`
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	Amount      float64 `json:"amount"`
	State       int     `json:"state"`
	Direct      string  `json:"direct"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}

func (t *FilWithdraw) AddWithdraw(db *gorm.DB) error {

	err := db.Where("business_no = ?", t.BusinessNo).FirstOrCreate(&t)
	if err != nil {
		return err.Error
	}
	return nil
}
