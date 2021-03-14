package exchange

import (
	"testing"
)

// Basic unit test for Calculator function
func TestCalculator(t *testing.T) {

	type fields struct {
		input   [2]string
		exRates ExchangeRate
	}

	tests := []struct {
		name     string
		fields   fields
		result   [3]string
		hasError bool
	}{
		{"test_BASIC_CALC_GPB",
			fields{
				[2]string{"1.00", ""}, ExchangeRate{
					Rates: []Rate{{
						Mid: 5.0,
					}},
				},
			},
			[3]string{"1.000000", "5.000000", "5.000000"},
			false,
		},
		{"test_BASIC_CALC_PLN",
			fields{
				[2]string{"", "5.00"}, ExchangeRate{
					Rates: []Rate{{
						Mid: 5.0,
					}},
				},
			},
			[3]string{"1.000000", "5.000000", "5.000000"},
			false,
		},
		{"test_BASIC_CALC_GPB_02",
			fields{
				[2]string{"1.00", "abc"}, ExchangeRate{
					Rates: []Rate{{
						Mid: 5.0,
					}},
				},
			},
			[3]string{"1.000000", "5.000000", "5.000000"},
			false,
		},
		{"test_BASIC_CALC_PLN_02",
			fields{
				[2]string{"abc", "5.0"}, ExchangeRate{
					Rates: []Rate{{
						Mid: 5.0,
					}},
				},
			},
			[3]string{"1.000000", "5.000000", "5.000000"},
			false,
		},
		{"test_BASIC_BAD_INPUT",
			fields{
				[2]string{"", ""}, ExchangeRate{
					Rates: []Rate{{
						Mid: 5.0,
					}},
				},
			},
			[3]string{},
			true,
		},
		{"test_BASIC_BAD_INPUT_02",
			fields{
				[2]string{"1.abc", "abc"}, ExchangeRate{
					Rates: []Rate{{
						Mid: 5.0,
					}},
				},
			},
			[3]string{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := RatesCalculator{exchangeRate: tt.fields.exRates}
			if res, err := c.Calculator(tt.fields.input[0], tt.fields.input[1]); res != tt.result || (err == nil && tt.hasError) {
				t.Errorf("assertion error: %s, %s", res, err)
			}
		})
	}
}
