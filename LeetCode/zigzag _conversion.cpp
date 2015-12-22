/*
 * zigzag _conversion.cpp
 *
 *  Created on: 2015年12月22日
 *      Author: wangqiying
 */
//https://leetcode.com/problems/zigzag-conversion/
#include <string>
using std::string;

class Solution {
public:
	string convert(string s, int numRows) {
		if(numRows <= 1) return s;
		string ret;
		for (int j = 0; j < numRows; j++) {
			for (int i = j, k = 0; i < s.size(); k++,i = (2*numRows-2)*k +j) {
				ret.append(1, s[i]);
				if(j ==0 || j == numRows-1)
				{
					continue;
				}
				if(i+(numRows- j-1)*2 < s.size())
				{
					ret.append(1, s[i+(numRows- j-1)*2]);
				}
			}
		}
		return ret;
	}
};

