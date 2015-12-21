/*
 * add_two_numbers.cpp
 *
 *  Created on: 2015年12月21日
 *      Author: wangqiying
 */

// https://leetcode.com/problems/add-two-numbers/
#include <stdio.h>
struct ListNode {
	int val;
	ListNode *next;
	ListNode(int x) :
			val(x), next(NULL) {
	}
};
class Solution {
public:
	ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
		int rest = 0;
		ListNode* new_list = NULL;
		ListNode* new_prev = NULL;
		while(l1 != NULL || l2 != NULL || rest > 0)
		{
			int v = rest;
			if(NULL != l1)
			{
				v += l1->val;
				l1 = l1->next;
			}
			if(NULL != l2)
			{
				v += l2->val;
				l2 = l2->next;
			}
			if(v >= 10)
			{
				v -= 10;
				rest = 1;
			}else
			{
				rest = 0;
			}
			ListNode* node = new ListNode(v);
			if(NULL == new_list)
			{
				new_list = node;
			}
			if(NULL != new_prev)
			{
				new_prev->next = node;
			}
			new_prev = node;
		}
		return new_list;
	}
};

