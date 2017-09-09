package main

import "testing"

func TestIsPrimeNumber(t *testing.T) {
	primeNumbers := []int{2, 3, 5}
	for _, v := range primeNumbers {
		result := isPrimeNumber(v)
		if result != true {
			t.Errorf("Got", result, "for number %v, expected true")
		}
	}
	result := isPrimeNumber(4)
	if result != false {
		t.Errorf("Got", result, "for number %v, expected false")
	}
}
