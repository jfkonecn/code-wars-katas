package kata

import "slices"

func Solution(word string) string {
	runes := []rune(word)
	slices.Reverse(runes)
	return string(runes)
}
