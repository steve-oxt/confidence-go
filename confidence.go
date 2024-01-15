package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Confidence struct {
	rate float64
	start float64
	previous float64
	current float64
	interval float64
	end float64
	tests int
	positive []int
	percentage []float64
	pretty_percentage []string
	validated bool
	Results string
}

func (c *Confidence) Validate() {
	if c.rate == 0.00 { c.rate = 0.67 }
	if c.start == 0.00 { c.start = 1000.00 }
	if c.previous == 0.00 { c.previous = 999.75 }
	if c.current == 0.00 { c.current = 1000.00 }
	if c.interval == 0.00 { c.interval = 0.25 }
	if c.end == 0.00 { c.end = 1000.00 }
	if c.tests == 0 { c.tests = 1000 }
	if c.positive == nil { c.positive = []int{0, 0, 0, 0, 0, 0, 0} }
	if c.percentage == nil { c.percentage = []float64{0.01, 0.025, 0.05, 0.10, 1.00, 2.00, 4.00} }
	if c.pretty_percentage == nil { c.pretty_percentage = []string{"one_percent", "two_percent", "five_percent", "ten_percent", "one_hundred_percent", "two_hundred_percent", "four_hundred_percent"} }
	if c.Results == "" { c.Results = fmt.Sprintf("{\"time\": %d", time.Now().Unix()) }
	c.validated = true
}

func (c Confidence) Calculate() {
	if ! c.validated {c.Validate()}
	for i := 0; i < 1000; i++ {
		for j := 0; j < c.tests; j++ {
			y := rand.Float64()
			if y > c.rate {
				if c.current > c.previous {
					c.current, c.previous = c.current+c.interval, c.current
				} else if c.current < c.previous {
					c.current, c.previous = c.current-c.interval, c.current
				}
			} else {
				c.current, c.previous = c.previous, c.current
			}
		}
		change := math.Abs(c.current-c.start) / c.start
		for j := 0; j < 7; j++ {
			if change > c.percentage[j] {
				c.positive[j]++
			}
		}
		c.current = c.start
		c.previous = c.start - c.interval
	}
	for i := 0; i < 7; i++ {
		success := float64(c.positive[i]) / 1000.00 * 100.00
		c.Results += fmt.Sprintf(",\"%s\": %.2f", c.pretty_percentage[i], success)
	}
	fmt.Printf("%s,\"current_price\": %.2f}\n", c.Results, c.end)
}

func main() {
	c := Confidence{}
	c.Calculate()
}