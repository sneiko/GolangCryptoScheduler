package controllers

import (
	"CryptoTest/internal/models"
	"CryptoTest/internal/repositories"
	s "CryptoTest/internal/services"
	"CryptoTest/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
)

// CryptoLast godoc
// @ID crypto_last
// @Accept json
// @Produce json
// @Success 200 models.CryptoExchangeRate Response
// @Router /api/cryptoLast [get]
func CryptoLast(c *fiber.Ctx) error {
	pair := c.Params("pair", "btc-usdt")
	symbols := s.NewKucoinService(&repositories.Repos)

	result, err := symbols.GetLast(pair)
	if err != nil {
		panic(err)
	}

	exchangeRate := models.CryptoExchangeRate{
		OnCreateDateTime: result.OnCreateDateTime,
		Value:            result.Value,
	}

	return c.JSON(exchangeRate)
}

// CryptoHistory godoc
// @ID crypto_history
// @Accept json
// @Produce json
// @Success 200 models.CryptoHistory Response
// @Router /api/cryptoHistory [post]
func CryptoHistory(c *fiber.Ctx) error {
	pair := c.Params("pair", "btc-usdt")
	page, err := strconv.Atoi(c.Get("page", "1"))
	if err != nil {
		panic(err)
	}

	symbols := s.NewKucoinService(&repositories.Repos)

	result, err := symbols.GetAllByPairWithPage(pair, page)
	if err != nil {
		panic(err)
	}

	var history models.CryptoHistory
	for _, v := range result {
		var er models.CryptoExchangeRate
		helpers.MapFromTo(v, &er)
		history.History = append(history.History, er)
	}
	history.Total = len(history.History)

	return c.JSON(history)
}

// FiatLast godoc
// @ID fiat_last
// @Accept json
// @Produce json
// @Success 200 models.FiatExchangeRate Response
// @Router /api/fiatLast [get]
func FiatLast(c *fiber.Ctx) error {
	pair := c.Params("pair", "USD")
	symbols := s.NewCbrService(&repositories.Repos)

	result, err := symbols.GetLast(pair)
	if err != nil {
		panic(err)
	}

	symbolValue := strings.ReplaceAll(result.Value, ",", ".")
	val, err := strconv.ParseFloat(symbolValue, 32)
	if err != nil {
		return err
	}

	exchangeRate := models.FiatExchangeRate{
		OnCreateDateTime: result.OnCreateDateTime,
		Value:            val,
	}

	return c.JSON(exchangeRate)
}

// FiatHistory godoc
// @ID fiat_history
// @Accept json
// @Produce json
// @Success 200 models.FiatHistory Response
// @Router /api/fiatHistory [post]
func FiatHistory(c *fiber.Ctx) error {
	pair := c.Params("pair", "USD")
	page, err := strconv.Atoi(c.Get("page", "1"))
	if err != nil {
		panic(err)
	}

	symbols := s.NewCbrService(&repositories.Repos)

	result, err := symbols.GetAllByPairWithPage(pair, page)
	if err != nil {
		panic(err)
	}

	var history models.CryptoHistory
	for _, v := range result {
		var er models.CryptoExchangeRate
		helpers.MapFromTo(v, &er)
		history.History = append(history.History, er)
	}
	history.Total = len(history.History)

	return c.JSON(history)
}
