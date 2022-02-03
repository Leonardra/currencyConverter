package controller

import (
	"currencyConverter/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ConvertCurrencyExchange(exchangeTable map[string]map[string]float64) gin.HandlerFunc{
	currencyExchangeRequest := new(model.CurrencyExchangeRequestDto)
	return func(context *gin.Context){
		err := context.Bind(currencyExchangeRequest)
		err1, convertedAmount := convertToCurrency(exchangeTable, currencyExchangeRequest)

		currencyResponseDto := model.CurrencyExchangeResponseDto{
			From: currencyExchangeRequest.From,
			To: currencyExchangeRequest.To,
			AmountToConvert: currencyExchangeRequest.AmountToConvert,
			ConvertedAmount: convertedAmount,
		}
		context.JSON(http.StatusOK, currencyResponseDto)
		if err != nil || err1 != nil{
			return
		}
	}
}

func  convertToCurrency(exchangeTable map[string]map[string]float64,
	currencyExchangeRequest *model.CurrencyExchangeRequestDto) (error, float64){
	_, ok1 := exchangeTable[strings.ToUpper(currencyExchangeRequest.From)]
	_, ok2 := exchangeTable[strings.ToUpper(currencyExchangeRequest.To)]

	if !ok1 || !ok2{
		return errors.New("currency does not exist"), 0.0
	}
	result := currencyExchangeRequest.AmountToConvert * exchangeTable[currencyExchangeRequest.From][currencyExchangeRequest.To]
		return nil, result
}

