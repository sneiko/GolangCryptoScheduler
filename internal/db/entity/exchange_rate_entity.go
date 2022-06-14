package entity

import "time"

type ExchangeRate struct {
	*ModelId
	OnCreateDateTime time.Time `db:"oncreatedatetime"`
	Value            string    `db:"value"`
	CurrencyId       int32     `db:"currencyid"`
}
