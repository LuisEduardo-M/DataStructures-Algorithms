package FizzBuzz

import "fmt"

func fizzBuzz(n int) []string {
	r := []string{}

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			r = append(r, "FizzBuzz")
			//fmt.Print("Fizz Buzz")
		case i%3 == 0:
			r = append(r, "Fizz")
			//fmt.Print("Fizz")
		case i%5 == 0:
			r = append(r, "Buzz")
			//fmt.Print("Buzz")
		default:
			r = append(r, fmt.Sprint(i))
			//fmt.Print(i)
		}
	}
	return r
}
