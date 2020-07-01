package string

import "fmt"

// ReverseString reverses the array of bytes
func ReverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	var low, hi = 0, len(s)-1

	for low < hi {
		s[low], s[hi] = s[hi], s[low]
		low++
		hi--
	}

	fmt.Printf("Solution: %s\n", s)
}