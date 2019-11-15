
#include <stdio.h>
#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    int trap(vector<int>& height) {
        if (height.size() < 3) {
            return 0;
        }

        int all = 0;
        for (int i = 1; i < height.size() - 1; i++) {
            int vol = volume(height, i);
            if (vol > 0) {
                all += vol;
            }
        }
        return all;
    }
    int volume(vector<int>& height, int idx) {
        int l = 0;
        for (int i = idx - 1; i >= 0; i--) {
            if (l < height[i]) {
                l = height[i];
            }
        }
        int r = 0;
        for (int i = idx + 1; i < height.size(); i++) {
            if (r < height[i]) {
                r = height[i];
            }
        }
        int min = (l<r)?l:r;
        return min - height[idx];
    }
};

int main(int argc, const char *argv[])
{
    vector<int> v = {2,0,2};
    //vector<int> v = {0,1,0,2,1,0,1,3,2,1,2,1};
    Solution s;
    cout << s.trap(v) << endl;
    return 0;
}
