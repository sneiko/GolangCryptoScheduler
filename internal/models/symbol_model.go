package models

type RootSymbol struct {
	Items []Symbol `json:"Data"`
}

type Symbol struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}
