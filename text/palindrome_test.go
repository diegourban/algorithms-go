package text

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "single character", s: "a", want: true},
		{name: "even length", s: "abba", want: true},
		{name: "odd length", s: "racecar", want: true},
		{name: "not a palindrome", s: "abc", want: false},
		// TODO: decide and add the edge cases: the empty string (algorithms-java returns false),
		// case ("Abba"), and multi-byte runes ("áéá").
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.s); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
