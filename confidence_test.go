package main

import (
	"testing"
	//"fmt"
	//"time"
)

func TestConfidence_Calculate(t *testing.T) {
	type fields struct {
		rate              float64
		start             float64
		previous          float64
		current           float64
		interval          float64
		end               float64
		tests             int
		positive          []int
		percentage        []float64
		pretty_percentage []string
		Results           string
		validated         bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "no field", fields: fields{ }},
		{name: "default", fields: fields{ rate: 0.67, start: 1000.00, previous: 999.75, current: 1000.00, interval: 0.25, end: 1000.00, tests: 10000}},
		{name: "small interval", fields: fields{ rate: 0.67, start: 100.00, previous: 99.99, current: 100.00, interval: 0.01, end: 100.00, tests: 10000 }},
		{name: "option_size", fields: fields{ rate: 0.67, start: 0.10, previous: 0.09, current: 0.10, interval: 0.01, end: 0.10, tests: 1000 }},
		{name: "set_arrays", fields: fields{ rate: 0.67, start: 100.00, previous: 99.99, current: 100.00, interval: 0.01, end: 100.00, tests: 100000, validated: true }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Confidence{
				rate:              tt.fields.rate,
				start:             tt.fields.start,
				previous:          tt.fields.previous,
				current:           tt.fields.current,
				interval:          tt.fields.interval,
				end:               tt.fields.end,
				tests:             tt.fields.tests,
				positive:          tt.fields.positive,
				percentage:        tt.fields.percentage,
				pretty_percentage: tt.fields.pretty_percentage,
				Results:           tt.fields.Results,
				validated:         tt.fields.validated,
			}
			c.Calculate()
		})
	}
}
