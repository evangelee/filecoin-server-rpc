package models

type FilRecharge struct {
	ID          int     `gorm:"primary_key" json:"id"`
	Txid        string  `json:"txid"`
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	Amount      float64 `json:"amount"`
	NowBalance  float64 `json:"now_balance"`
	LastBalance float64 `json:"last_balance"`
	State       int     `json:"state"`
	CreateTime  int64   `json:"create_time"`
}
