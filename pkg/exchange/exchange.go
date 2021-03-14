package exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// contains API address used to get exchange rate
// and ExchangeRate instance containing currently used rate
type RatesCalculator struct {
	apiAddr      string
	exchangeRate ExchangeRate
}

// structures below are used to parse nbp api's json response
type Rate struct {
	No            string  `json:"no"`
	EffectiveDate string  `json:"effectiveDate"`
	Mid           float32 `json:"mid"`
}

type ExchangeRate struct {
	Table    string `json:"table"`
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Rates    []Rate `json:"rates"`
}

// Constructor for RatesCalculator
// Gets from addr current exchange rate as initial value
func NewCalculator(addr string) (*RatesCalculator, error) {
	c := RatesCalculator{
		apiAddr: addr,
	}
	if err := c.getExchangeRate(&c.exchangeRate); err != nil {
		return &RatesCalculator{}, err
	}
	return &c, nil
}

// Gets from apiAddr URL rate, parses it and modifies exchangeRate
// inputs:
//		ExchangeRate structure instance
// returns:
// 		error, if one occurs
func (c RatesCalculator) getExchangeRate(exchangeRate *ExchangeRate) error {
	res, err := http.Get(c.apiAddr)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad request response: %s", res.Status)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}

	if jsonErr := json.Unmarshal(body, &exchangeRate); jsonErr != nil {
		return readErr
	}
	return nil
}

// calculates the other number based on exchange rate and user input
// 		note: gbp value has priority over pln.
// returns: [3]string{}, error
//		[0]string -> gbp value
//		[1]string -> pln value
//		[2]string -> used exchange rate
//
//		calculated values are based on first valid float,
//		if nor gbp, nor pln value is valid returns parse error
func (c RatesCalculator) Calculator(gbpStr string, plnStr string) ([3]string, error) {
	rate := c.exchangeRate.Rates[0].Mid
	resultStr := [3]string{gbpStr, plnStr, fmt.Sprintf("%f", rate)}

	gbp, gbpErr := strconv.ParseFloat(gbpStr, 32)
	pln, plnErr := strconv.ParseFloat(plnStr, 32)

	if gbpErr == nil {
		resultStr[0] = fmt.Sprintf("%f", float32(gbp))
		resultStr[1] = fmt.Sprintf("%f", float32(gbp)*rate)
	} else if plnErr == nil {
		resultStr[0] = fmt.Sprintf("%f", float32(pln)/rate)
		resultStr[1] = fmt.Sprintf("%f", float32(pln))
	} else {
		return [3]string{}, fmt.Errorf("invalid both inputs: %s; %s", gbpErr, plnErr)
	}

	return resultStr, nil
}

// wrapper for Calculator method, updates exchange rate and calls c.Calculator
func (c *RatesCalculator) CalculateRates(gbpStr string, plnStr string) ([3]string, error) {
	if err := c.getExchangeRate(&c.exchangeRate); err != nil {
		return [3]string{}, err
	}

	return c.Calculator(gbpStr, plnStr)
}
