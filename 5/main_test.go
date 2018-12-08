package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	demos := []string{"aaaa", "1abcdcba2", "1abbcbba"}
	values := []string{"aaaa", "abcdcba", "abbcbba"}
	length := len(demos)
	var value string
	for i := 0; i < length; i++ {
		value = longestPalindrome(demos[i])
		if value == values[i] {
			t.Logf("equal: %s=%s", value, values[i])
		} else {
			t.Errorf("not equal: %s!=%s", value, values[i])
		}
	}
}
