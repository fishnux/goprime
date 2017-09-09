package main

import "testing"

func TestIsPrimeNumber(t *testing.T) {
	result := isNumberPrime(3)
	if result != true {
		t.Error("Got", result, "expected true")
	}
}
