
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */

struct ListNode* addList(struct ListNode *h, int v) {
    struct ListNode *n = (struct ListNode *)malloc(sizeof(struct ListNode));
    n->val = v;
    n->next = NULL;
    if (h) h->next = n;
    
    return n;
}
struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    struct ListNode *h = NULL;
    struct ListNode *n = NULL;
    int cnt = 0;
    int v = 0;
    while (l1 || l2) {
        v = cnt;
        cnt = 0;
        if (l1) {
            v += l1->val;
            l1 = l1->next;
        }
        if (l2) {
            v += l2->val;
            l2 = l2->next;
        }
        if (v >= 10) {
            v -= 10;
            cnt = 1;
        }
        n = addList(n, v);
        if (!h) h = n;
    }
    if (cnt) n = addList(n, cnt);
    return h;
}
