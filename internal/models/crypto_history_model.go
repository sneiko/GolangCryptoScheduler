package models

import "time"

type CryptoHistory struct {
	Total   int                  `json:"total"`
	History []CryptoExchangeRate `json:"history"`
}

type CryptoExchangeRate struct {
	OnCreateDateTime time.Time `json:"timestamp"`
	Value            string    `json:"last"`
}
