package models

type ParsedExchangeRateRoot struct {
	Data ParsedExchangeRate `json:"data"`
}

type ParsedExchangeRate struct {
	IsFiat     bool
	CurrencyId int32
	SymbolName string `json:"symbol"`
	Value      string `json:"last"`
}
