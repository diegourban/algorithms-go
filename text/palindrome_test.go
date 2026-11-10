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
		{name: "long", s: "socorrammesubinoonibusemmarrocos", want: true},
		{name: "not a palindrome", s: "ab", want: false},
		{name: "empty is not a palindrome", s: "", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.s); got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
