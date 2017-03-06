package trial

func Fizzbuzz(i int) (result string) {
	if i%3 == 0 && i%5 == 0 {
		result = "FizzBuzz"
	} else if i%3 == 0 {
		result = "Fizz"
	} else if i%5 == 0 {
		result = "Buzz"
	} else {
		result = "Nothing"
	}
	return result
}

func Burden() bool {

	return true
}
