package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Payment struct {
	Id                 int             `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Client             Client          `json:"client" gorm:"column:client;not null"`
	PaySum             decimal.Decimal `json:"pay_sum" gorm:"type:decimal(19,4);not null"`
	PayDate            time.Time       `json:"pay_date" gorm:"type:timestamp;not null"`
	OrderTransactionId int             `json:"orderTransactionId"`
	TransactionId      string          `json:"transactionId"`
	Cancelled          bool            `json:"cancelled" gorm:"default:false"`
}
