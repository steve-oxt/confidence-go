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
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{name: "test1", fields: fields{ tests: 100000 }},
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
			}
			c.Calculate()
		})
	}
}
