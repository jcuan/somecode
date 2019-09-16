#include <iostream>

using namespace std;

int main()
{
    return 0;
}

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *p1 = l1;
        ListNode *p2 = l2;
        ListNode *resHead = new ListNode(0);
        ListNode *pres = resHead;
        ListNode *newNode;
        unsigned int carry = 0;
        unsigned int value;
        while (p1 != NULL || p2 != NULL) {
            value = carry;
            if (p1 != NULL) {
                value += p1->val;
                p1 = p1->next;
            }
            if (p2 != NULL) {
                value += p2->val;
                p2 = p2->next;
            }
            if (value > 9) {
                carry = 1;
                value %= 10;
            } else {
                carry = 0;
            }
            newNode = new ListNode(value);
            pres->next = newNode;
            pres = pres->next;
        }
        if (carry != 0) {
            newNode = new ListNode(carry);
            pres->next = newNode;
        }
        pres = resHead;
        resHead = resHead->next;
        delete pres;
        return resHead;
    }
};