package app

import (
	"CryptoTest/internal/db"
	"CryptoTest/internal/repositories"
	service "CryptoTest/internal/services"
	"fmt"
)

func loadRepositories() {
	symbolsTable := db.NewDbTable(db.SymbolsTableNameScheme)
	exchangeRatesTable := db.NewDbTable(db.ExchangeRatesNameScheme)
	repositories.Repos = repositories.NewAppRepositories(
		symbolsTable,
		exchangeRatesTable,
	)
}

func preloadDataCheck() {
	kucoinService := service.NewKucoinService(&repositories.Repos)
	_, err := kucoinService.UpdateSymbols()
	if err != nil {
		fmt.Println(err)
	}
}
