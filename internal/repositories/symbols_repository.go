package repositories

import (
	"CryptoTest/internal/db"
	"CryptoTest/internal/db/entity"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"strings"
)

type SymbolsRepository struct {
	*db.DbTable
}

var (
	addSymbolQuery         = "INSERT INTO %s (symbol, name, isfiat) VALUES ('%s', '%s', %t)"
	getByPairQuery         = "SELECT id, symbol, name FROM %s WHERE symbol='%s' AND IsFiat=%t"
	getByPairWithPageQuery = "SELECT id, symbol, name FROM %s WHERE symbol='%s' AND IsFiat=%t LIMIT 10 OFFSET %d"
)

func NewSymbolRepository(c db.DbTable) SymbolsRepository {
	return SymbolsRepository{&c}
}

func (t *SymbolsRepository) Add(obj entity.SymbolEntity) (entity.SymbolEntity, error) {
	sql := fmt.Sprintf(addSymbolQuery,
		t.Table,
		obj.Symbol,
		obj.Name,
		obj.IsFiat)
	_, err := db.Db.Exec(context.Background(), sql)

	return obj, err
}

func (t *SymbolsRepository) GetAll() ([]entity.SymbolEntity, error) {
	sql := fmt.Sprintf("SELECT * FROM %s", t.Table)
	rows, err := db.Db.Query(context.Background(), sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var symbols []entity.SymbolEntity
	err = pgxscan.ScanAll(&symbols, rows)
	return symbols, err
}

func (t *SymbolsRepository) GetByPair(pair string, isFiat bool, page *int) (*entity.SymbolEntity, error) {
	var sql string
	if page == nil || *page > 0 {
		sql = fmt.Sprintf(getByPairQuery, t.Table, strings.ToUpper(pair), isFiat)
	} else {
		pageCountItems := *page * 10
		sql = fmt.Sprintf(getByPairWithPageQuery, t.Table, strings.ToUpper(pair), isFiat, pageCountItems)
	}

	row, err := db.Db.Query(context.Background(), sql)
	defer row.Close()

	if err != nil {
		return nil, err
	}

	var symbol entity.SymbolEntity
	if err := pgxscan.ScanOne(&symbol, row); err != nil {
		return nil, err
	}

	return &symbol, nil
}
