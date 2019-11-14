
#include <vector>
#include <stdio.h>

using namespace std;

class Solution {
public:
    void set(vector<int>& nums, int val) {
        if (val <= 0 || val > nums.size()) {
            return;
        }
        if (nums[val-1] == val) {
            return;
        }
        int tmp = nums[val-1];
        nums[val-1] = val;
        set(nums, tmp);
    }
    int firstMissingPositive(vector<int>& nums) {
        for (int i : nums) {
            if (i <= 0 || i > nums.size()) {
                continue;
            }
            if (nums[i-1] == i) {
                continue;
            }
            set(nums, i);
        }
        for (int i = 1; i <= nums.size(); i++) {
            if (nums[i-1] != i) {
                return i;
            }
        }
        return nums.size()+1;
    }
    void printLine(vector<int>& nums) {
        for (int i : nums) {
            printf("%d ", i);
        }
        printf("\n");
    }
};

int main(int argc, const char *argv[])
{
    Solution s;
    vector<int> nums = {};
    printf("%d\n", s.firstMissingPositive(nums));
    return 0;
}
