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

import java.util.HashMap;

class ValidAnagram {
    public static void main(String[] args) {
        for (String arg: args) {
            System.out.print(arg + " -> ");
            System.out.println(isAnagram(arg.split(" ")[0], arg.split(" ")[1]));
        }
    }

    static boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) return false;

        HashMap<Character, Integer> map = new HashMap<>();

        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            map.put(c, map.getOrDefault(c, 0) + 1);
        }

        for (int i = 0; i < t.length(); i++) {
            char c = t.charAt(i);
            if (!map.containsKey(c)) return false;
            map.put(c, map.get(c) - 1);
        }

        return map.values().stream().allMatch(num -> num == 0);
    }
}
