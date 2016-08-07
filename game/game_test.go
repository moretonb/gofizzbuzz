package game

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Game tests are about to run")
	result := m.Run()
	fmt.Println("Game tests are done executing")

	os.Exit(result)
}

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
		result := Play(entry.in)

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
		result := Play(entry.in)
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
		result := Play(entry.in)
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
		result := Play(entry.in)
		expected := "FizzBuzz"

		if result != expected {
			t.Log("expected output ", expected)
			t.Log("actual output ", result)
			t.Fail()
		}
	}
}

func BenchmarkOfPlay(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Play(1)
	}
}

func ExamplePlay() {
	fmt.Println(Play(1))
	fmt.Println(Play(3))
	fmt.Println(Play(5))
	fmt.Println(Play(15))

	// Output:
	// 1
	// Fizz
	// Buzz
	// FizzBuzz
}
