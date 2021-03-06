package controller

import (
	"currencyConverter/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_convertToCurrency(t *testing.T) {
	exchangeTable := model.CreateExchangeTable()
	currencyExchange := model.CurrencyExchangeRequestDto{From: "NGN", To: "GHS", AmountToConvert: 45000}
	_, result := convertToCurrency(exchangeTable.CurrencyTable, &currencyExchange)
	assert.Equal(t, 765.00, result)
}
