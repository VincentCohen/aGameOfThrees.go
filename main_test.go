package main

import "testing"

func TestIsDividable(t *testing.T) {
	result := isDividable(3)
	if result == false {
		t.Errorf("Can not divide 3", result)
	}
}

func TestDivide(t *testing.T) {
	result := divide(3)

	if result != 1 {
		t.Errorf("Can not divide 3", result)
	}
}

func TestMinus(t *testing.T) {
	result := minus(4)

	if result != 3 {
		t.Errorf("Minus failed", result)
	}
}

func TestAdd(t *testing.T) {
	result := add(4)

	if result != 5 {
		t.Errorf("Add failed", result)
	}
}
