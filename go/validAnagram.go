/*
	Valid Anagram
	https://leetcode.com/problems/valid-anagram/description/

	Given two strings s and t, return true if t is an anagram of s, and false otherwise.

	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

	Example 1:
	Input: s = "anagram", t = "nagaram"
	Output: true

	Example 2:
	Input: s = "rat", t = "car"
	Output: false

	Constraints:
	1 <= s.length, t.length <= 5 * 104
	s and t consist of lowercase English letters.

	Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, in := range os.Args[1:] {
		inputs := strings.Split(in, " ")
		fmt.Println(inputs, "->", isAnagram(inputs[0], inputs[1]))
	}
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			hash[s[i]]++
		} else {
			hash[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if _, ok := hash[t[i]]; ok {
			hash[t[i]]--
		} else {
			hash[t[i]] = 1
		}
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}
