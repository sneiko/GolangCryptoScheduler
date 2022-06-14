package entity

type SymbolEntity struct {
	*ModelId
	Symbol string `db:"symbol"`
	Name   string `db:"name"`
	IsFiat bool   `db:"isfiat"`
}
