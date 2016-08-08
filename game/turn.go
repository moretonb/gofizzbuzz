package game

import "strconv"

func turn(number int) string {
	if number%15 == 0 {
		return "FizzBuzz"
	}
	if number%5 == 0 {
		return "Buzz"
	}
	if number%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(number)
}
