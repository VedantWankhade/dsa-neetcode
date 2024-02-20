/*
	Two Sum

	https://leetcode.com/problems/two-sum/description/

	Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

	You may assume that each input would have exactly one solution, and you may not use the same element twice.

	You can return the answer in any order.

	Example 1:
	Input: nums = [2,7,11,15], target = 9
	Output: [0,1]
	Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

	Example 2:
	Input: nums = [3,2,4], target = 6
	Output: [1,2]

	Example 3:
	Input: nums = [3,3], target = 6
	Output: [0,1]

	Constraints:
	2 <= nums.length <= 104
	-109 <= nums[i] <= 109
	-109 <= target <= 109
	Only one valid answer exists.

	Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(os.Args)
SKIP:
	for _, arg := range os.Args[1:] {
		in := strings.Split(arg, " ")
		numsStr, targetStr := in[0], in[1]
		target, err := strconv.Atoi(targetStr)
		if err != nil {
			continue SKIP
		}
		numsStr = strings.TrimPrefix(numsStr, "[")
		numsStr = strings.TrimSuffix(numsStr, "]")
		numsStrArr := strings.Split(numsStr, ",")
		var nums []int
		for _, num := range numsStrArr {
			anum, err := strconv.Atoi(num)
			if err != nil {
				continue SKIP
			}
			nums = append(nums, anum)
		}
		fmt.Printf("%v -> %v\n", arg, twoSum(nums, target))
	}
}

func twoSum(nums []int, target int) []int {
	// this exit is not needed since 2 <= nums.length - stated in the problem statement
	if len(nums) < 2 {
		return nil
	}
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		comp := target - nums[i]
		if cmp, ok := hash[comp]; ok {
			return []int{i, cmp}
		} else {
			hash[nums[i]] = i
		}
	}
	return []int{}
}
