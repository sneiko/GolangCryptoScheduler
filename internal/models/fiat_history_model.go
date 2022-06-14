package models

import "time"

type FiatHistory struct {
	Total   int                `json:"total"`
	History []FiatExchangeRate `json:"history"`
}

type FiatExchangeRate struct {
	OnCreateDateTime time.Time `json:"timestamp"`
	Value            float64   `json:"value"`
}
