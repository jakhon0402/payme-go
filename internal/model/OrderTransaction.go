package model

import "time"

type OrderTransaction struct {
	Id                      int       `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	TransactionId           string    `json:"transaction_id" gorm:"column:transaction_id"`
	TransactionCreationTime time.Time `json:"transaction_creation_time" gorm:"type:timestamp;not null"`
	PerformTime             time.Time `json:"perform_time" gorm:"type:timestamp;not null"`
	CancelTime              time.Time `json:"cancel_time" gorm:"type:timestamp;not null"`
	Reason                  int       `json:"reason"`
	State                   int       `json:"state"`
}
