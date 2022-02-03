package route

import (
	"currencyConverter/controller"
	"currencyConverter/model"
	"github.com/gin-gonic/gin"
)

func CurrencyRouter(route *gin.Engine){
	exchangeTable := model.CreateExchangeTable()
	route.POST("/currencyConverter", controller.ConvertCurrencyExchange(exchangeTable.CurrencyTable))
}