/*
 * ATOI.CPP
 *
 *  Created on: 2015年12月22日
 *      Author: wangqiying
 */
//https://leetcode.com/problems/string-to-integer-atoi/
#include <string>
using std::string;
class Solution {
public:
	bool isnumber(char c) {
		return c >= '0' && c <= '9';
	}
	bool isspace(char c) {
		return c == ' ' || c == '\t';
	}
	int myAtoi(string str) {
		if (str.empty()) {
			return 0;
		}
		int i = 0;
		bool negative = false;
		bool started = false;
		int64_t v = 0;
		for (i = 0; i < str.size(); i++) {
			if (started) {
				if (!isnumber(str[i])) {
					break;
				}
			} else {
				if (isspace(str[i])) {
					continue;
				} else if (str[i] == '-') {
					negative = true;
					started = true;
					continue;
				} else if (str[i] == '+') {
					started = true;
					continue;
				} else if (isnumber(str[i])) {
					started = true;
				}else
				{
					break;
				}
			}
			v *= 10;
			v += (str[i] - '0');
			if (!negative && v > 2147483647) {
				v = 2147483647;
				break;
			}
			if(negative && v >= 2147483648)
			{
				return -2147483648;
			}
		}
		return negative ? -v : v;
	}

};

