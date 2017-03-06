package trial

import "testing"

func TestFizzbuzz(t *testing.T) {
	/*
		testData := map[int]string{
			1:  "Nothing",
			2:  "Nothing",
			3:  "Fizz",
			4:  "Nothing",
			5:  "Buzz",
			15: "FizzBuzz",
		}
	*/
	testData := map[int]string{
		3:  "Fizz",
		5:  "Buzz",
		15: "FizzBuzz",
	}

	for key, value := range testData {
		if res := Fizzbuzz(key); value != res {
			t.Errorf("not matching:%s:%s", res, value)
		}
	}
}
