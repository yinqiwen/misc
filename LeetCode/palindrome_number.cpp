/*
 * palindrome_number.cpp
 *
 *  Created on: 2015年12月22日
 *      Author: wangqiying
 */
//https://leetcode.com/problems/palindrome-number/
class Solution {
public:
	bool isPalindrome(int x) {
		if (x < 0) {
			return false;
		}
		if (x < 10) {
			return true;
		}
		int v = 0;
		int xv = x;
		while (x) {
			int left = x % 10;
			v *= 10;
			v += left;
			x /= 10;
		}
		return xv == v;

	}
};

