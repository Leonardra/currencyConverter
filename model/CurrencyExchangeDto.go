package model

type CurrencyExchangeRequestDto struct{
	From string
	To string
	AmountToConvert float64
}

type CurrencyExchangeResponseDto struct{
	From string
	To string
	AmountToConvert float64
	ConvertedAmount float64
}
