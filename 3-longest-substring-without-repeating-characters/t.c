inline int find(char *s, int start, int end, char c) {
    for (; start <= end; start++) {
       if (s[start] == c) break; 
    }
    if (start>end) return -1;
    else return start;
}

int lengthOfLongestSubstring(char * s){
    int start = 0;
    int end = 0;
    int max = 0;
    while (s[end]) {
        int pos = find(s, start, end, s[end+1]);                
        int len = end-start+1;
        if (pos >= 0) {
            start = pos+1;
        }
        if (max < len) max = len;
        end++;
    }
    return max;
}
