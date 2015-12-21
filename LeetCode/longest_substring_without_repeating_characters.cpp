/*
 * longest_substring_without_repeating_characters.cpp
 *
 *  Created on: 2015年12月21日
 *      Author: wangqiying
 */

//https://leetcode.com/problems/longest-substring-without-repeating-characters/
#include <string>
#include <vector>
using std::string;

class Solution {
public:
	int lengthOfLongestSubstring(string s) {
		int max_len = -1;
		std::vector<int> flags(256);
		int start_idx = 0;
		for (int i = 0; i < s.size(); i++) {
			int flag = flags[s[i]];
			int last_duplicate_idx = (flag >> 1);
			int len = -1;
			if ((flag & 0x1) && last_duplicate_idx >= start_idx) {
				len = i - start_idx;
				start_idx = last_duplicate_idx + 1;
			} else if (i == s.size() - 1) {
				len = i - start_idx + 1;
			}
			if (len > max_len) {
				max_len = len;
			}
			flags[s[i]] = (i << 1) + 1;
		}
		if (max_len == -1) {
			max_len = s.size();
		}
		return max_len;
	}
};
