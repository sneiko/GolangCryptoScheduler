package repositories

import (
	"CryptoTest/internal/db"
	"CryptoTest/internal/db/entity"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
)

type ExchangeRatesRepository struct {
	*db.DbTable
}

var (
	addExchangeRateQuery    = "INSERT INTO %s (CurrencyId, Value) VALUES (%d, '%s') RETURNING *"
	getAllByCurrencyIdQuery = "SELECT * FROM %s WHERE CurrencyId=%d ORDER BY OnCreateDateTime DESC LIMIT 1000"
	getLastQuery            = "SELECT * FROM %s WHERE CurrencyId=%d ORDER BY OnCreateDateTime DESC LIMIT 1"
)

func NewExchangeRatesRepository(c db.DbTable) ExchangeRatesRepository {
	return ExchangeRatesRepository{&c}
}

func (t *ExchangeRatesRepository) Add(obj entity.ExchangeRate) (*entity.ExchangeRate, error) {
	sql := fmt.Sprintf(addExchangeRateQuery,
		t.Table,
		obj.CurrencyId,
		obj.Value)

	rows, err := db.Db.Query(context.Background(), sql)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var e entity.ExchangeRate
	if err := pgxscan.ScanOne(&e, rows); err != nil {
		return nil, err
	}

	return &e, err
}

func (t *ExchangeRatesRepository) GetAllBySymbol(s *entity.SymbolEntity) ([]entity.ExchangeRate, error) {
	sql := fmt.Sprintf(getAllByCurrencyIdQuery, t.Table, s.Id)
	rows, err := db.Db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entity.ExchangeRate
	if err := pgxscan.ScanAll(&result, rows); err != nil {
		return nil, err
	}

	return result, nil
}

func (t *ExchangeRatesRepository) GetLast(s *entity.SymbolEntity) (*entity.ExchangeRate, error) {
	sql := fmt.Sprintf(getLastQuery, t.Table, s.Id)
	rows, err := db.Db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result entity.ExchangeRate
	if err := pgxscan.ScanOne(&result, rows); err != nil {
		return nil, err
	}

	return &result, nil
}
