package game

import "testing"

func TestNumberWhenOne(t *testing.T) {
	result := Play(1)
	expected := "1"

	if result != expected {
		t.Log("expected output ", expected)
		t.Log("actual output ", result)
		t.Fail()
	}
}

func TestNumberWhenTwo(t *testing.T) {
	result := Play(2)
	expected := "2"

	if result != expected {
		t.Log("expected output ", expected)
		t.Log("actual output ", result)
		t.Fail()
	}
}

func TestFizzWhenThree(t *testing.T) {
	result := Play(3)
	expected := "Fizz"

	if result != expected {
		t.Log("expected output ", expected)
		t.Log("actual output ", result)
		t.Fail()
	}
}

func TestFizzWhenSix(t *testing.T) {
	result := Play(6)
	expected := "Fizz"

	if result != expected {
		t.Log("expected output ", expected)
		t.Log("actual output ", result)
		t.Fail()
	}
}

func TestBuzzWhenFive(t *testing.T) {
	result := Play(5)
	expected := "Buzz"

	if result != expected {
		t.Log("expected output ", expected)
		t.Log("actual output ", result)
		t.Fail()
	}
}

func TestBuzzWhenTen(t *testing.T) {
	result := Play(10)
	expected := "Buzz"

	if result != expected {
		t.Log("expected output ", expected)
		t.Log("actual output ", result)
		t.Fail()
	}
}

func TestFizzBuzzWhenFifteen(t *testing.T) {
	result := Play(15)
	expected := "FizzBuzz"

	if result != expected {
		t.Log("expected output ", expected)
		t.Log("actual output ", result)
		t.Fail()
	}
}

func TestFizzBuzzWhenThirty(t *testing.T) {
	result := Play(30)
	expected := "FizzBuzz"

	if result != expected {
		t.Log("expected output ", expected)
		t.Log("actual output ", result)
		t.Fail()
	}
}

func BenchmarkOfPlay(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Play(1)
	}
}
