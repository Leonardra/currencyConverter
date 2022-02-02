package controller

import (
	"currencyConverter/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func convertCurrencyExchange(exchangeTable map[string]map[string]float64) gin.HandlerFunc{
	currencyExchangeRequest := new(model.CurrencyExchangeRequestDto)
	return func(context *gin.Context){
		err := context.Bind(currencyExchangeRequest)
		convertToCurrency(exchangeTable, currencyExchangeRequest)
		if err != nil {
			return
		}
	}
}

func  convertToCurrency(exchangeTable map[string]map[string]float64,
	currencyExchangeRequest *model.CurrencyExchangeRequestDto) float64{

	var result = 0.0
	for key, value := range exchangeTable {
		if currencyExchangeRequest.From == key {
			result = currencyExchangeRequest.AmountToConvert * value[currencyExchangeRequest.To]
			fmt.Println("Success")
		}
	}
		return result
}

