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

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

class ContainsDuplicate {
    public static void main(String[] args) {
        for (String arg : args) {
           System.out.print(arg); 
           arg = arg.replaceFirst("\\[", "");
           arg = arg.replaceFirst("\\]", "");
           int[] nums = Arrays.stream(arg.split(",")).mapToInt(Integer::parseInt).toArray();
           System.out.println(" -> " + containsDuplicate(nums));
        }
    }

    static boolean containsDuplicate(int[] nums) {
        if (nums.length <= 1) return false;
        Map<Integer, Integer> hash = new HashMap<>();
        for (int num : nums) {
            if (hash.containsKey(num)) {
                return true;
            } else {
                hash.put(num, hash.getOrDefault(num, 0) + 1);
            }
        }
        return false;
    }
}
