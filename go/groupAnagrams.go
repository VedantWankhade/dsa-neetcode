/*
	Group Anagrams

	https://leetcode.com/problems/group-anagrams/description/

	Given an array of strings strs, group the anagrams together. You can return the answer in any order.

	An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

	Example 1:
	Input: strs = ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	Example 2:
	Input: strs = [""]
	Output: [[""]]

	Example 3:
	Input: strs = ["a"]
	Output: [["a"]]

	Constraints:
	1 <= strs.length <= 104
	0 <= strs[i].length <= 100
	strs[i] consists of lowercase English letters.
*/

package main

import (
	"fmt"
	"slices"
)

func main() {
	// fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagramsFast([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	hash := make(map[string][]string)
	for _, str := range strs {
		s := []rune(str)
		// fmt.Println(s)
		slices.Sort(s)
		// fmt.Println(s)
		arr, _ := hash[string(s)]
		arr = append(arr, str)
		hash[string(s)] = arr
	}
	// fmt.Println(hash)
	var ans [][]string
	for _, vals := range hash {
		ans = append(ans, vals)
	}
	return ans
}

func groupAnagramsFast(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}
	hash := make(map[[26]int][]string)
	for _, str := range strs {
		arr := [26]int{}
		for _, r := range str {
			arr[r-97] = arr[r-97] + 1
		}
		// fmt.Println(arr)
		wordsArr, _ := hash[arr]
		wordsArr = append(wordsArr, str)
		hash[arr] = wordsArr
	}
	var ans [][]string
	for _, vals := range hash {
		ans = append(ans, vals)
	}
	return ans
}
