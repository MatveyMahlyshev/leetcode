package tests

import (
	"fmt"
	"leetcode/functions"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name        string
		inputString string
		expectedInt int
	}{
		{"042", "042", 42},
		{"0-42", "0-42", 0},
		{"-042", "-042", -42},
		{"-420", "-420", -420},
		{"420", "420", 420},
		{"0", "0", 0},
		{"4201", "4201", 4201},
		{"abcd-421020ab", "abcd-421020ab", 0},
		{"abcd-021-345", "abcd-021-345", 0},
		{"-abcd123123", "-abcd123123", 0},
		{"42", "42", 42},
		{"1337c0d3", "1337c0d3", 1337},
		{"0-1", "0-1", 0},
		{"words and 987", "words and 987", 0},
		{"20000000000000000000", "20000000000000000000", 2147483647},
		{"  +  413", "  +  413", 0},
		{"3.14159", "3.14159", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.MyAtoi(tt.inputString)
			if result != tt.expectedInt {
				t.Errorf("myAtoi(%s)\thas result: %d\texpected: %d\n", tt.inputString, result, tt.expectedInt)
			} else {
				fmt.Printf("PASS!\tmyAtoi(%s)\thas result: %d\texpected: %d\n", tt.inputString, result, tt.expectedInt)
			}
		})
	}

}

func TestReverse(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedOutput int
	}{
		{"123", 123000, 321},
		{"-123", -123, -321},
		{"120", 120, 21},
		{"1534236469", 1534236469, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.Reverse(tt.input)
			if result != tt.expectedOutput {
				t.Errorf("reverse(%d)\toutput: %d\t expected: %d\n", tt.input, result, tt.expectedOutput)
			} else {
				fmt.Printf("PASS!\treverse(%d)\toutput: %d\t expected: %d\n", tt.input, result, tt.expectedOutput)
			}
		})
	}

}

func TestPrefixCount(t *testing.T) {
	tests := []struct {
		name           string
		words          []string
		pref           string
		expectedOutput int
	}{
		{"at", []string{"pay", "attention", "practice", "attend"}, "at", 2},
		{"code", []string{"leetcode", "win", "loops", "success"}, "at", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.PrefixCount(tt.words, tt.pref)
			if result != tt.expectedOutput {
				t.Errorf("FAILED!\tprefixCount(%s, %s)\toutput: %d\texpected: %d\n", tt.words, tt.pref, result, tt.expectedOutput)
			} else {
				fmt.Printf("PASS!\tprefixCount(%s, '%s')\toutput: %d\texpected: %d\n", tt.words, tt.pref, result, tt.expectedOutput)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{
		{"abcabcbb", "abcabcbb", 3},
		{"bbbbb", "bbbbb", 1},
		{"pwwkew", "pwwkew", 3},
		{"", "", 0},
		{"abcdefg", "abcdefg", 7},
		{"abcdabcdef", "abcdabcdef", 6},
		{"dvdf", "dvdf", 3},
		{"anviaj", "anviaj", 5},
		{"jbpnbwwd", "jbpnbwwd", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.LengthOfLongestSubstring(tt.input)
			if result != tt.expectedOutput {
				t.Errorf("lengthOfLongestSubstring(%s)\toutput: %d\texpected: %d\n", tt.input, result, tt.expectedOutput)
			} else {
				fmt.Printf("PASS!\tlengthOfLongestSubstring(%s)\toutput: %d\texpected: %d\n", tt.input, result, tt.expectedOutput)
			}
		})
	}
}

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name           string
		words          []string
		k              int
		expectedOutput []string
	}{
		{"i love", []string{"i", "love", "leetcode", "i", "love", "coding"}, 2, []string{"i", "love"}},
		{"the is sunny day", []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4, []string{"the", "is", "sunny", "day"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.TopKFrequent(tt.words, tt.k)
			if len(result) != len(tt.expectedOutput) {
				t.Errorf("topkFrequent(%s, %d)\toutput: %s\texpected: %s\n", tt.words, tt.k, result, tt.expectedOutput)
			} else {
				isCorrect := true
				for i := 0; i < len(result); i++ {
					if result[i] != tt.expectedOutput[i] {
						t.Errorf("topkFrequent(%s, %d)\toutput: %s\texpected: %s\n", tt.words, tt.k, result, tt.expectedOutput)
						isCorrect = false
						break
					}
				}
				if isCorrect {
					fmt.Printf("PASS!\ttopkFrequent(%s, %d)\toutput: %s\texpected: %s\n", tt.words, tt.k, result, tt.expectedOutput)
				}
			}
		})
	}
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name           string
		s              string
		expectedOutput string
	}{
		{"hello world", "  hello  world  ", "world hello"},
		{"the sky is blue", "the sky is blue", "blue is sky the"},
		{"a good   example", "a good   example", "example good a"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.ReverseWords(tt.s)
			if result != tt.expectedOutput {
				t.Errorf("reverseWords(%s)\toutput: %s\texpected: %s\n", tt.s, result, tt.expectedOutput)
			} else {
				fmt.Printf("PASS!\treverseWords(%s)\toutput: %s\texpected: %s\n", tt.s, result, tt.expectedOutput)
			}
		})
	}
}

func TestCountPrefixSuffixPairs(t *testing.T) {
	tests := []struct {
		name           string
		words          []string
		expectedOutput int
	}{
		{"a", []string{"a", "aba", "ababa", "aa"}, 4},
		{"pa", []string{"pa", "papa", "ma", "mama"}, 2},
		{"abab", []string{"abab", "ab"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.CountPrefixSuffixPairs(tt.words)
			if result != tt.expectedOutput {
				t.Errorf("countPrefixSuffixPairs(%s)\toutput: %d\texpected: %d\n", tt.words, result, tt.expectedOutput)
			} else {
				fmt.Printf("countPrefixSuffixPairs(%s)\toutput: %d\texpected: %d\n", tt.words, result, tt.expectedOutput)
			}
		})
	}
}
