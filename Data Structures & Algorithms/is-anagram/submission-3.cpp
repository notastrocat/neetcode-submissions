class Solution {
public:
    bool isAnagram(string s, string t) {
        if ( s.length() != t.length() ) {
            return false;
        }

        std::map<char, int> ts;  // table for s
        std::map<char, int> tt;  // table for t
        for (auto ch : s) {
            ts[ch]++;
        }
        for (auto ch : t) {
            tt[ch]++;
        }

        for (auto ch : t) {
            if (ts[ch] <= 0) return false;
            if (ts[ch] != tt[ch]) return false;
        }

        return true;
    }
};
