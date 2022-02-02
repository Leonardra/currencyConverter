package model

type CurrencyExchangeRequestDto struct{
	From string
	To string
	AmountToConvert float64
}
