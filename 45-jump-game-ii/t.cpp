
#include <stdio.h>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    int jump(vector<int>& nums) {
        int i = 0;
        int cnt = 1;
        if (nums[0] >= nums.size()) {
            return cnt;
        }
        while (true) {
            vector<int> v = getMaxStart(nums, i);
            cnt++;
            i += v[0];
            if (v[0] >= v[1]) {
                return 0;
            }
            if (i + v[1] >= nums.size()) {
                break;
            }
        }
        return cnt;
    }
    vector<int> getMaxStart(vector<int>& nums, int idx) {
        int max = 0;
        int pos = 0;
        for (int i = 0; i < nums[idx]; i++) {
            int len = i + nums[idx+i];
            if (max < len) {
                max = len;
                pos = i;
            }
        }
        return vector<int>{pos, max};
    }
};

int main(int argc, const char *argv[])
{
    Solution s;
    vector<int> v = {2,3,1,1,4};
    cout << s.jump(v) << endl;
    return 0;
}
