
#include <vector>
#include <stdio.h>

using namespace std;

class Solution {
public:
    Solution(void) {m_cnt = 0;m_done=false;}
    void solveSudoku(vector<vector<char>>& board) {
        solveSudokuHelper(board);
        board = m_board;
        //printBoard(board);
    }

    void solveSudokuHelper(vector<vector<char>> board) {
        if (m_done) {
            return;
        }
        bool done = true;
        m_cnt++;

        // sure
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] == '.') {
                    vector<char> v = getPossibleVal(board, i, j);
                    if (v.size() == 1) {
                        board[i][j] = v[0];
                    }
                }
            }
        }

        // maybe
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] == '.') {
                    done = false;
                    vector<char> v = getPossibleVal(board, i, j);
                    if (v.size() == 0) {
                        return;
                    }
                    //printf("line...(%d,%d)\n", i, j);
                    //printLine(v);
                    for (char c : v) {
                        board[i][j] = c;
                        solveSudokuHelper(board);
                    }
                }
            }
        }
        if (done) {
            m_done = true;
            //printf("done.....%d times\n", m_cnt);
            m_board = board;
        }
    }

    void printBoard(vector<vector<char>>& board) {
        for (vector<char> v : board) {
            for (char c : v) {
                printf("%c ", c);
            }
            printf("\n");
        }
        printf("\n");
    }
    void printLine(vector<char> line) {
        for (char c : line) {
            printf("%c ", c);
        }
        printf("\n");
    }
    vector<char> getPossibleVal(vector<vector<char>>& board, int x, int y) {
        vector<char> v = {'1','2','3','4','5','6','7','8','9'};
        // 行
        for (char c : board.at(x)) {
            if (c == '.') {
                continue;
            }
            vector<char>::iterator it = find(v.begin(), v.end(), c);
            if (it != v.end()) {
                v.erase(it);
            }
        }

        // 列
        for (int i = 0; i < 9; i++) {
            char c = board.at(i).at(y);
            if (c == '.') {
                continue;
            }
            vector<char>::iterator it = find(v.begin(), v.end(), c);
            if (it != v.end()) {
                v.erase(it);
            }
        }

        // 块
        int pos_x = (x/3)*3;
        int pos_y = (y/3)*3;
        for (int i = pos_x; i < pos_x+3; i++) {
            if (i == x) {
                continue;
            }
            for (int j = pos_y; j < pos_y+3; j++) {
                if (j == y) {
                    continue;
                }

                char c = board.at(i).at(j);
                if (c == '.') {
                    continue;
                }
                vector<char>::iterator it = find(v.begin(), v.end(), c);
                if (it != v.end()) {
                    v.erase(it);
                }
            }
        }
        return v;
    }

private:
    vector<vector<char>> m_board;
    int m_cnt;
    bool m_done;
};

int main(int argc, const char *argv[])
{
    Solution s;
    vector<vector<char>> board = {
        {'5', '3', '.', '.', '7', '.', '.', '.', '.'},
        {'6', '.', '.', '1', '9', '5', '.', '.', '.'},
        {'.', '9', '8', '.', '.', '.', '.', '6', '.'},
        {'8', '.', '.', '.', '6', '.', '.', '.', '3'},
        {'4', '.', '.', '8', '.', '3', '.', '.', '1'},
        {'7', '.', '.', '.', '2', '.', '.', '.', '6'},
        {'.', '6', '.', '.', '.', '.', '2', '8', '.'},
        {'.', '.', '.', '4', '1', '9', '.', '.', '5'},
        {'.', '.', '.', '.', '8', '.', '.', '7', '9'},
    };
    s.printBoard(board);
    s.solveSudoku(board);
    s.printBoard(board);
    return 0;
}
