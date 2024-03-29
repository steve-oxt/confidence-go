package confidence

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Confidence struct {
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
	validated         bool
	Results           string
}

func (c *Confidence) validate() {
	if c.rate == 0.00 {
		c.rate = 0.67
	}
	if c.start == 0.00 {
		c.start = 1000.00
	}
	if c.previous == 0.00 {
		c.previous = 999.75
	}
	if c.current == 0.00 {
		c.current = 1000.00
	}
	if c.interval == 0.00 {
		c.interval = 0.25
	}
	if c.end == 0.00 {
		c.end = 1000.00
	}
	if c.tests == 0 {
		c.tests = 1000
	}
	if c.positive == nil {
		c.positive = []int{0, 0, 0, 0, 0, 0, 0}
	}
	if c.percentage == nil {
		c.percentage = []float64{0.01, 0.025, 0.05, 0.10, 1.00, 2.00, 4.00}
	}
	if c.pretty_percentage == nil {
		c.pretty_percentage = []string{"one_percent", "two_percent", "five_percent", "ten_percent", "one_hundred_percent", "two_hundred_percent", "four_hundred_percent"}
	}
	if c.Results == "" {
		c.Results = fmt.Sprintf("{\"time\": %d", time.Now().Unix())
	}
	c.validated = true
}

func (c *Confidence) setArrays() {
	c.positive = []int{0, 0, 0, 0, 0, 0, 0}
	c.percentage = []float64{0.01, 0.025, 0.05, 0.10, 1.00, 2.00, 4.00}
	c.pretty_percentage = []string{"one_percent", "two_percent", "five_percent", "ten_percent", "one_hundred_percent", "two_hundred_percent", "four_hundred_percent"}
}

func (c *Confidence) New(ticks float64, seconds float64, start_time float64, end_time float64, rate float64, start float64, previous float64, interval float64, end float64) {
	t := start_time
    if start_time == 0 {
        t = float64(time.Now().Unix())
    }
	et := end_time
    if end_time == 0 {
        et = float64(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 20, 0, 0, 0, time.UTC).Unix())
    }
	c.rate = rate
	c.start = start
	c.previous = previous
	c.current = start
	c.interval = interval
	c.end = end
	c.tests = int((float64(et) - float64(t)) / seconds * ticks)
}

func (c *Confidence) Calculate() {
	if !c.validated {
		c.validate()
	} else {
		c.setArrays()
	}
	c.current = c.start
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
