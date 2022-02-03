package main

import (
	"currencyConverter/route"
	"github.com/gin-gonic/gin"
)

func main() {
	router:=gin.Default()
	route.CurrencyRouter(router)
	error := router.Run()
	if error != nil {
		return
	}
}
