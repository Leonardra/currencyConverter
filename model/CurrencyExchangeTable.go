package model

type ExchangeTable struct{
	CurrencyTable map[string]map[string]float64
}


func CreateExchangeTable() *ExchangeTable{
	exchangeTable :=new(ExchangeTable)
	exchangeTable.CurrencyTable = make(map[string]map[string]float64)
	exchangeTable.CurrencyTable["NGN"] = map[string]float64{"CED": 0.017, "SHI":0.27}
	exchangeTable.CurrencyTable["CED"] = map[string]float64{"NGN": 60.27, "SHI":16.46}
	exchangeTable.CurrencyTable["SHI"] = map[string]float64{"NGN": 3.66, "CED":0.061}

	return exchangeTable
}