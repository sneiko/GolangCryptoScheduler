package servicesb

import (
	"CryptoTest/internal/db/entity"
	"CryptoTest/internal/models"
	"CryptoTest/internal/repositories"
	"encoding/xml"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"log"
	"net/http"
)

type CbrService struct {
	Repos *repositories.AppRepositorues
}

const (
	GetFiatStatsUrl = "http://www.cbr.ru/scripts/XML_daily.asp"
)

func NewCbrService(repos *repositories.AppRepositorues) CbrService {
	return CbrService{
		Repos: repos,
	}
}

func (s *CbrService) getStats() ([]models.ParsedFiatModel, error) {
	resp, err := http.Get(GetFiatStatsUrl)
	if err != nil {
		return nil, err
	}

	decodedBody := xml.NewDecoder(resp.Body)
	decodedBody.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	var umResult models.ParsedRootFiatModel
	err = decodedBody.Decode(&umResult)
	if err != nil {
		return nil, err
	}
	return umResult.Data, nil
}

func (s *CbrService) UpdateSymbols() (bool, error) {
	result, err := s.getStats()
	if err != nil {
		return false, err
	}

	for _, item := range result {
		if _, err := s.Repos.Symbols.Add(entity.SymbolEntity{
			Name:   item.Name,
			Symbol: item.Code,
			IsFiat: true,
		}); err != nil {
			log.Println(err)
		}
	}
	return true, nil
}

func (s *CbrService) UpdateCourses() (bool, error) {
	result, err := s.getStats()
	if err != nil {
		return false, err
	}

	allSymbols, err := s.Repos.Symbols.GetAll()
	if err != nil {
		return false, err
	}

	for _, item := range result {
		for _, symbol := range allSymbols {
			var findedSymbol entity.SymbolEntity
			if symbol.Symbol == item.Code {
				findedSymbol = symbol

				if _, err := s.Repos.ExchangeRates.Add(entity.ExchangeRate{
					Value:      item.Value,
					CurrencyId: findedSymbol.Id,
				}); err != nil {
					log.Println(err)
				}
			}
		}
	}

	return true, nil
}

func (s *CbrService) GetAllByPair(pair string) ([]entity.ExchangeRate, error) {
	symbol, err := repositories.Repos.Symbols.GetByPair(pair, true, nil)
	if err != nil {
		panic(err)
	}

	history, err := repositories.Repos.ExchangeRates.GetAllBySymbol(symbol)
	return history, err
}

func (s *CbrService) GetAllByPairWithPage(pair string, page int) ([]entity.ExchangeRate, error) {
	symbol, err := repositories.Repos.Symbols.GetByPair(pair, true, &page)
	if err != nil {
		panic(err)
	}

	history, err := repositories.Repos.ExchangeRates.GetAllBySymbol(symbol)
	return history, err
}

func (s *CbrService) GetLast(pair string) (*entity.ExchangeRate, error) {
	symbol, err := repositories.Repos.Symbols.GetByPair(pair, true, nil)
	if err != nil {
		panic(err)
	}

	history, err := repositories.Repos.ExchangeRates.GetLast(symbol)
	return history, err
}
