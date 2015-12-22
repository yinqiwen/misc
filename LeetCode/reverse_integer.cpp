/*
 * reverse_integer.cpp
 *
 *  Created on: 2015年12月22日
 *      Author: wangqiying
 */
//https://leetcode.com/problems/reverse-integer/
class Solution {
public:
	int reverse(int x) {
		int v = 0;
		while (x) {
			int last = x % 10;
			x = x / 10;
			if (v > 214748364 || v < -214748364) {
				return 0;
			}
			v *= 10;
			v += last;
		}
		return v;
	}
};

