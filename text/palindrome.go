package text

// IsPalindrome checks if a string reads the same forwards and backwards.
// The comparison is exact, character by character: case, spaces and punctuation all count, so {@code "Aba"} is not a palindrome.
// Runs in O(n) time.
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return true
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
