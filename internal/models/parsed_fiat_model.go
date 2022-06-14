package models

import "encoding/xml"

type ParsedRootFiatModel struct {
	XMLName xml.Name          `xml:"ValCurs"`
	Data    []ParsedFiatModel `xml:"Valute"`
}

type ParsedFiatModel struct {
	XMLName xml.Name `xml:"Valute"`
	Id      string   `xml:"ID,attr"`
	Name    string   `xml:"Name"`
	Code    string   `xml:"CharCode"`
	Nominal int      `xml:"Nominal"`
	Value   string   `xml:"Value"`
}
