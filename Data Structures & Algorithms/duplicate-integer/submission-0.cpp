class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        std::map<int, int> table;
        for (auto num : nums) {
            table[num]++;
        }

        for (auto num : nums) {
            if (table[num] > 1) return true;
        }

        return false;
    }
};