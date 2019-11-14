
#include <stdio.h>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int longestValidParentheses(string s) {
        int max = 0;
        for (int i = 0; i < s.length();) {
            if (s[i] == ')') {
                i++;
                continue;
            }
            int len = calc(s.substr(i));
            if (max < len) {
                max = len;
            }
            if (len == 0) {
                len++;
            }
            i += len;
        }
        return max;
    }
    int calc(string s) {
        int v = 0;
        int cnt = 0;
        int base = 0;
        int i = 0;
        for (i = 0; i < s.length(); i++) {
            if (s[i] == '(') {
                v++;
            } else {
                if (v == 0) {
                    break;
                }
                v--;
                cnt++;
                if (v == 0) {
                    base += cnt;
                    cnt = 0;
                }
            }
        }
        return base*2;
    }
};

int main(int argc, const char *argv[])
{
    Solution s;
    //string str = "((()))())";
    string str = "()(()";
    cout << s.longestValidParentheses(str) << endl;
    return 0;
}
