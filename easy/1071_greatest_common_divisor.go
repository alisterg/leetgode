package easy

// For two strings s and t, we say "t divides s" if and only if s = t + t + t +
// ... + t + t (i.e., t is concatenated with itself one or more times).
//
// Given two strings str1 and str2, which are uppercase words, return the
// largest string x such that x divides both str1 and str2.
//
// Example: str1 = "ABCABC", str2 = "ABC"
// Output: "ABC"
//
// Example: str1 = "ABABAB", str2 = "ABAB"
// Output: "AB"
//
// Example: str1 = "LEET", str2 = "CODE"
// Output: ""
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	result := gcd(len(str1), len(str2))
	return str1[:result]
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
