#Koa Go Engineer Challenge
This is my solution to the challenge for the Go Engineer role.
***
#GitHub Repository
https://github.com/Leonardra/currencyConverter
***


#Tools
* Testify: assertion library for unit testing
* Gin-Gonic : Web service framework.

#Rest API Usage
***

##Create Inventory Item
> #### P0ST http://localhost:8080/currencyConverter

#####Parameter
 ```json
  {
  "From": "NGN",
  "To": "GHS",
  "AmountToConvert":45000.00
 }
```

##Response
####200 OK on successful request

```json
  {
  "From": "NGN",
  "To": "GHS",
  "AmountToConvert": 45000,
  "ConvertedAmount": 765
 }
```



