package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	Calculate()
}

func Calculate() {
	rate := 0.67
	start := 0.10
	previous := 0.09
	current := 0.10
	interval := 0.01
	end := 0.10
	tests := 1000
	positive := [7]int{0, 0, 0, 0, 0, 0, 0}
	percentage := [7]float64{0.01, 0.025, 0.05, 0.10, 1.00, 2.00, 4.00}
	pretty_percentage := [7]string{"one_percent", "two_percent", "five_percent", "ten_percent", "one_hundred_percent", "two_hundred_percent", "four_hundred_percent"}
	results := fmt.Sprintf("{\"time\": %d", time.Now().Unix())

	for i := 0; i < 1000; i++ {
		for j := 0; j < tests; j++ {
			y := rand.Float64()
			if y > rate {
				if current > previous {
					current, previous = current+interval, current
				} else if current < previous {
					current, previous = current-interval, current
				}
			} else {
				current, previous = previous, current
			}
		}
		change := math.Abs(current-start) / start
		for j := 0; j < 7; j++ {
			if change > percentage[j] {
				positive[j]++
			}
		}
		current = start
		previous = start - interval
	}
	for i := 0; i < 7; i++ {
		success := float64(positive[i]) / 1000.00 * 100.00
		results += fmt.Sprintf(",\"%s\": %.2f", pretty_percentage[i], success)
	}
	fmt.Printf("%s,\"current_price\": %.2f}\n", results, end)
}

