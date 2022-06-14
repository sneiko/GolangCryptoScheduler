package repositories

import "CryptoTest/internal/db"

type AppRepositorues struct {
	Symbols       SymbolsRepository
	ExchangeRates ExchangeRatesRepository
}

var Repos AppRepositorues

func NewAppRepositories(symbols db.DbTable, exchangeRates db.DbTable) AppRepositorues {
	return AppRepositorues{
		Symbols:       NewSymbolRepository(symbols),
		ExchangeRates: NewExchangeRatesRepository(exchangeRates),
	}
}
