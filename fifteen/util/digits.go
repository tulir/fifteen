package util

// Digits returns the number of digits in the given number.
//
// Zero is special cased to return zero.
func Digits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
