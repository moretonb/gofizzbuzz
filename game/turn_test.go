package game

import (
	"fmt"
	"testing"
)

var numberTable = []struct {
	in  int
	out string
}{
	{1, "1"},
	{2, "2"},
	{4, "4"},
	{7, "7"},
	{8, "8"},
}

func TestNumberWhenNotMatchingOtherRules(t *testing.T) {
	t.Parallel()

	for _, entry := range numberTable {
		result := turn(entry.in)

		if result != entry.out {
			t.Log("expected output ", entry.out)
			t.Log("actual output ", result)
			t.Fail()
		}
	}
}

var fizzTable = []struct {
	in int
}{
	{3},
	{6},
	{9},
	{12},
	{33},
}

func TestFizzWhenMultipleOfThree(t *testing.T) {
	t.Parallel()

	for _, entry := range fizzTable {
		result := turn(entry.in)
		expected := "Fizz"

		if result != expected {
			t.Log("expected output ", expected)
			t.Log("actual output ", result)
			t.Fail()
		}
	}
}

var buzzTable = []struct {
	in int
}{
	{5},
	{10},
	{20},
	{25},
	{95},
}

func TestBuzzWhenMultipleOfFive(t *testing.T) {
	t.Parallel()

	for _, entry := range buzzTable {
		result := turn(entry.in)
		expected := "Buzz"

		if result != expected {
			t.Log("expected output ", expected)
			t.Log("actual output ", result)
			t.Fail()
		}
	}
}

var fizzbuzzTable = []struct {
	in int
}{
	{15},
	{30},
	{45},
	{60},
	{1500},
}

func TestFizzBuzzWhenMultipleOfThreeAndFive(t *testing.T) {
	t.Parallel()

	for _, entry := range fizzbuzzTable {
		result := turn(entry.in)
		expected := "FizzBuzz"

		if result != expected {
			t.Log("expected output ", expected)
			t.Log("actual output ", result)
			t.Fail()
		}
	}
}

func BenchmarkOfTurn(b *testing.B) {
	for index := 0; index < b.N; index++ {
		turn(1)
	}
}

func ExampleTurn() {
	fmt.Println(turn(1))
	fmt.Println(turn(3))
	fmt.Println(turn(5))
	fmt.Println(turn(15))

	// Output:
	// 1
	// Fizz
	// Buzz
	// FizzBuzz
}
