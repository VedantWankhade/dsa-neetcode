/*
   Contains Duplicate

   https://leetcode.com/problems/contains-duplicate/description/

   Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

   Example 1:
   Input: nums = [1,2,3,1]
   Output: true

   Example 2:
   Input: nums = [1,2,3,4]
   Output: false

   Example 3:
   Input: nums = [1,1,1,3,3,4,3,2,4,2]
   Output: true

   Constraints:
   1 <= nums.length <= 105
   -109 <= nums[i] <= 109
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	for _, in := range os.Args[1:] {
		in = strings.TrimLeft(in, "[")
		in = strings.TrimRight(in, "]")
		arr := strings.Split(in, ",")
		var newArr []int
		for _, i := range arr {
			in, _ := strconv.Atoi(i)
			newArr = append(newArr, in)
		}
		fmt.Println(arr, "->", containsDuplicate(newArr))
	}
}

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	m := make(map[int]int)
	for _, e := range nums {
		_, ok := m[e]
		if ok {
			return true
		} else {
			m[e] = 1
		}
	}
	return false
}
