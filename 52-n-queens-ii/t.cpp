
#include <stdio.h>
#include <vector>
#include <string>
#include <iostream>

using namespace std;

class Solution {
    /* 标识方法:
     * ..Q
     * Q..
     * .Q.
     * 记录为 {2,0,1}
     */
public:
    int totalNQueens(int n) {
        m_n = n;
        getPossibleMap(vector<int>{});

        vector<vector<string>> vvs;
        for (vector<int> map : m_map) {
            vvs.push_back(parseMap(map));
        }
        return vvs.size();
    }

    void getPossibleMap(vector<int> map) {

        if (map.size() == m_n) {
            m_map.push_back(map);
            return;
        }

        vector<int> pp = getPossiblePos(map, map.size());

        for (int p : pp) {
            if (p == -1) {
                continue;
            }
            //cout << "push back " << p << endl;
            vector<int> tmp = map;
            tmp.push_back(p);
            getPossibleMap(tmp);
        }
    }

    void printPp(const vector<int> &pp) {
        cout << "pp: ";
        for (int i : pp) {
            cout << i << " ";
        }
        cout << endl;
    }

    vector<int> getPossiblePos(const vector<int> &map, int line) {
        vector<int> pp;
        for (int i = 0; i < m_n; i++) {
            pp.push_back(i);
        }
        for (int i = 0; i < line; i++) {
            // 列
            pp[map[i]] = -1;

            // 斜
            int l = map[i] - (line - i);
            int r = map[i] + (line - i);
            if (l >= 0) {
                pp[l] = -1;
            }
            if (r < m_n) {
                pp[r] = -1;
            }
        }
        return pp;
    }

    vector<string> parseMap(const vector<int> &map) {
        vector<string> vs;
        //cout << "parse map:" << endl;
        for (int val : map) {
            string s = string(val, '.') + "Q" + string(m_n-val-1, '.');
            vs.push_back(s);
        }
        return vs;
    }

    void printVs(const vector<string> &vs) {
        //cout << "map:" << endl;
        for (string s : vs) {
            cout << s << endl;
        }
    }

private:
    int m_n;
    vector<vector<int>> m_map;
};

int main(int argc, const char *argv[])
{
    Solution s;
    vector<vector<string>> vvs = s.solveNQueens(4);
    for (vector<string> vs : vvs) {
        cout << "map: " << endl;
        s.printVs(vs);
        cout << endl;
    }
    return 0;
}
