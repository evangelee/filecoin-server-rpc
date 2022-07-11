package models

type FilWallet struct {
	ID         int     `gorm:"primary_key" json:"id"`
	Address    string  `json:"address"`
	PrivateKey string  `json:"private_key"`
	Balance    float64 `json:"balance"`
	CreateTime int64   `json:"create_time"`
	UpdateTime int64   `json:"update_time"`
}
