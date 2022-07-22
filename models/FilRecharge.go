package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type FilRecharge struct {
	ID          int     `gorm:"primary_key" json:"id"`
	Txid        string  `json:"txid"`
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	Amount      float64 `json:"amount"`
	NowBalance  float64 `json:"now_balance"`
	LastBalance float64 `json:"last_balance"`
	State       int     `json:"state"`
	Direct      string  `json:"direct"`
	CreateTime  int64   `json:"create_time"`
	UpdateTime  int64   `json:"update_time"`
}

func (t *FilRecharge) UpdateRecharge(db *gorm.DB) error {

	db.Transaction(func(tx *gorm.DB) error {

		var filwallet = &FilWallet{}
		tx.Where("address = ?", t.ToAddress).First(&filwallet)
		if filwallet != nil {
			if err := tx.Model(&FilRecharge{}).Where("txid = ?", t.Txid).UpdateColumns(map[string]interface{}{"state": 1, "now_balance": filwallet.Balance, "last_balance": t.Amount + filwallet.Balance, "update_time": time.Now().Unix()}).Error; err != nil {
				return err
			}

			if err := tx.Model(&FilWallet{}).Where("address = ?", t.ToAddress).UpdateColumns(map[string]interface{}{"balance": gorm.Expr("balance+?", t.Amount), "update_time": time.Now().Unix()}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return nil
}
