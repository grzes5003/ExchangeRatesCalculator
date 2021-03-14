# Exchange rate calculator

### Project structure
File `cmd/exchange-service/main.go` contains entry function with example usage.
Directory `pkg/exchange` consists of two files:
- `exchange.go`: exchange rate calculator code
- `exchange_test.go`: unit tests for calculator methods

### Assumptions

As an input calculator gets two strings (respectively GBP and PLN values), only one must
 be valid float (other can be empty or contain nonnumerical value).
It returns three strings as parsed float32: GBP, PLN and used exchange rate. 
If error occurs, calculator returns it.

It is up to frontend how data will be rounded/presented and possible errors handled. 

### Running 
- Main function:
```shell script
cd cmd/exchange-service
go run main.go
```

- Testing exchange package
```shell script
cd pkg/exchange
go test -v
```