package servicesb

import (
	"CryptoTest/internal/db/entity"
	"CryptoTest/internal/models"
	"CryptoTest/internal/repositories"
	"CryptoTest/pkg/helpers"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type KucoinService struct {
	Repos *repositories.AppRepositorues
}

const (
	GetSymbolsUrl     = "https://api.kucoin.com/api/v1/symbols"
	GetSymbolStatsUrl = "https://api.kucoin.com/api/v1/market/stats?symbol=%s"
)

func NewKucoinService(repos *repositories.AppRepositorues) KucoinService {
	return KucoinService{
		Repos: repos,
	}
}

func (s *KucoinService) UpdateSymbols() (bool, error) {
	resp, err := http.Get(GetSymbolsUrl)
	if err != nil {
		return false, err
	}

	body, err := helpers.GetBodyFromResp(resp)
	if err != nil {
		return false, err
	}

	var symbolsResult models.RootSymbol
	err = json.Unmarshal(body, &symbolsResult)
	if err != nil {
		return false, err
	}

	allSymbols, errAllSymbols := s.Repos.Symbols.GetAll()
	if errAllSymbols != nil {
		return false, errAllSymbols
	}

	for _, v := range symbolsResult.Items {

		result := entity.SymbolEntity{}
		helpers.MapFromTo(v, &result)

		isContain := false
		for _, item := range allSymbols {
			if item.Symbol == v.Symbol {
				isContain = true
			}
		}

		if !isContain {
			_, err := s.Repos.Symbols.Add(result)
			if err != nil {
				fmt.Printf("Cannot add: %s \n:: %s \n", v.Symbol, err)
				return false, err
			}
		}
	}

	return true, nil
}

func (s *KucoinService) UpdateStats(pair string) (*entity.ExchangeRate, error) {
	symbol, err := repositories.Repos.Symbols.GetByPair(pair, false, nil)

	url := fmt.Sprintf(GetSymbolStatsUrl, strings.ToUpper(symbol.Name))
	resp, err := http.Get(url)

	body, err := helpers.GetBodyFromResp(resp)

	var rate models.ParsedExchangeRateRoot
	err = json.Unmarshal(body, &rate)

	rate.Data.CurrencyId = symbol.Id
	rate.Data.SymbolName = symbol.Symbol

	var exchangeRateEntity entity.ExchangeRate
	helpers.MapFromTo(rate.Data, &exchangeRateEntity)

	addedRow, err := s.Repos.ExchangeRates.Add(exchangeRateEntity)

	return addedRow, err
}

func (s *KucoinService) GetAllByPair(pair string) ([]entity.ExchangeRate, error) {
	symbol, err := repositories.Repos.Symbols.GetByPair(pair, false, nil)
	if err != nil {
		panic(err)
	}

	history, err := repositories.Repos.ExchangeRates.GetAllBySymbol(symbol)
	return history, err
}

func (s *KucoinService) GetAllByPairWithPage(pair string, page int) ([]entity.ExchangeRate, error) {
	symbol, err := repositories.Repos.Symbols.GetByPair(pair, false, &page)
	if err != nil {
		panic(err)
	}

	history, err := repositories.Repos.ExchangeRates.GetAllBySymbol(symbol)
	return history, err
}

func (s *KucoinService) GetLast(pair string) (*entity.ExchangeRate, error) {
	symbol, err := repositories.Repos.Symbols.GetByPair(pair, false, nil)
	if err != nil {
		panic(err)
	}

	history, err := repositories.Repos.ExchangeRates.GetLast(symbol)
	return history, err
}
