package main

import "testing"


func TestAddition(t *testing.T) {

	got := Addition(3, 5)

	if got != 8 {

		t.Errorf("Addition(3, 5) got: %d want: 8", got)
	}
}

func TestSubtraction(t *testing.T) {

	got := Subtraction(20, 5)

	if got != 15 {

		t.Errorf("Subtraction(20, 5) got: %d want: 15", got)
	}
}

func TestMultiplication(t *testing.T) {

	got := Multiplication(5, 5)

	if got != 25{

		t.Errorf("Multiplication(5, 5) got: %d want: 25", got)
	}
}

func TestDivision(t *testing.T) {

	got, message, err := Division(4, 2)

	if got != 2 && err {

		t.Errorf("Division(4, 2) got: %d want: 2 %s", got, message)
	}
}