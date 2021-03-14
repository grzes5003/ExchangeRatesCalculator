package main

import (
	"exchangeCalculator/pkg/exchange"
	"fmt"
)

const ApiAddr = "http://api.nbp.pl/api/exchangerates/rates/a/gbp/?format=json"

// entry function with example usage
func main() {
	calc, err := exchange.NewCalculator(ApiAddr)
	if err != nil {
		println(err)
	}
	ans, err := calc.CalculateRates("13.3", "")
	fmt.Printf("%s, %s, %s", ans[0], ans[1], ans[2])
}
