
#include <string>
#include <vector>
#include <iostream>
#include <stdio.h>

using namespace std;

class Solution {
public:
    vector<int> findSubstring(string s, vector<string>& words) {
        m_over = false;

        vector<int> ret;
        if (words.size() == 0) {
            return ret;
        }
        m_word_len = words[0].length();

        for (int i = 0; i < s.length(); i++) {
            if (m_over) {
                break;
            }
            if (findStr(s.substr(i), words)) {
                ret.push_back(i);
            }
        }

        return ret;
    }
    bool findStr(string s, vector<string> words) {
        if (s.length() < words.size() * m_word_len) {
            m_over = true;
            return false;
        }
        while (s.size() > 0) {
            bool find = false;
            for (auto it=words.begin(); it!=words.end(); it++) {
                if (s.compare(0, m_word_len, *it) == 0) {
                    cout << "find " << *it << endl;
                    words.erase(it);
                    s = s.substr(m_word_len);
                    cout << "s now " << s << endl;
                    find = true;
                    break;
                }
            }
            if (words.size() == 0) {
                return true;
            }
            if (!find) {
                return false;
            }
        }
        return false;
    }

private:
    int m_word_len;
    bool m_over;
};

int main(int argc, const char *argv[])
{
    string str = "wordgoodgoodgoodbestword";
    vector<string> words = {"word","good","best","word"};

    Solution s;
    vector<int> v = s.findSubstring(str, words);
    for (int i : v) {
        printf("%d ", i);
    }
    printf("\n");
    return 0;
}
