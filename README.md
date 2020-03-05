# flights-searcher-api

Тестовый REST API для получения информации по перелетам на основании ответов от gateway (моки в виде xml ответов).

```
go get github.com/soypita/flights-searcher-api
go run main.go
```

**Описание API**

* Получение всех вариантов перелетов.  
(GET) http://localhost:8080/flights/all?offset=&limit=  
Метод возвращает массив всех вариантов перелетов, которые вернул gateway. Метод работает с пагинацией данных.
    - limit - максимальное число ответов на странице</li>
    - offset - отдать ответы, начиная с значения offset  

Структура ответа:  
```json
{
   "limit":372,
   "offset":0,
   "total":372,
   "results":[
      {
         "flights":[
            {
               "carrier":{
                  "carrierId":"AI",
                  "carrierName":"AirIndia"
               },
               "flightNumber":996,
               "source":"DXB",
               "destination":"DEL",
               "departureTime":"2018-10-22T0005",
               "arrivalTime":"2018-10-22T0445",
               "flightClass":"G",
               "stops":0,
               "ticketType":"E"
            },
            {
               "carrier":{
                  "carrierId":"AI",
                  "carrierName":"AirIndia"
               },
               "flightNumber":332,
               "source":"DEL",
               "destination":"BKK",
               "departureTime":"2018-10-22T1350",
               "arrivalTime":"2018-10-22T1935",
               "flightClass":"G",
               "stops":0,
               "ticketType":"E"
            }
         ],
         "prices":[
            {
               "amount":546.8,
               "type":"SingleAdult",
               "chargeType":"TotalAmount"
            }
         ],
         "currency":"SGD"
      }
   ]
}
```
* Получение статистики по доступным вариантам перелетов  
(GET) http://localhost:8080/flights/stats  
Метод анализирует ответ от gateway и возвращает варианты самого дешевого/самого дорого полета,
самого быстрого/самого медленного перелета и оптимальный вариант полета по критериям стоимости и скорости.  
Структура ответа:  
```json
{
   "optimalFlight":{
      "flights":[],
      "prices":[],
      "currency":"string"
   },
   "cheapestFlight":{
      "flights":[],
      "prices":[],
      "currency":"string"
   },
   "fastFlight":{
      "flights":[],
      "prices":[],
      "currency":"string"
   },
   "slowFlight":{
      "flights":[],
      "prices":[],
      "currency":"string"
   },
   "expensiveFlight":{
      "flights":[],
      "prices":[],
      "currency":"string"
   }
}
```
Структура массива flights и prices аналогична структуре из ответа *Получение всех вариантов перелетов*.

