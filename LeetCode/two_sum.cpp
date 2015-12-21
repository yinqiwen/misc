/*
 * two_sum.cpp
 *
 *  Created on: 2015年12月21日
 *      Author: wangqiying
 */

// https://leetcode.com/problems/two-sum/
#include <vector>
using std::vector;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
       vector<int> res;
    	for(int i = 0; i< nums.size(); i++)
    	{
    		int next =  target - nums[i];
    		for(int j = i+1; j < nums.size(); j++)
    		{
    			if(next == nums[j])
    			{
    				res.push_back(i);
    				res.push_back(j);
    				return res;
    			}
    		}
    	}
    	return res;
    }
};


