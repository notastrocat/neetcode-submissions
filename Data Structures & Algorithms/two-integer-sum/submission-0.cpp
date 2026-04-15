class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        std::map<int, int> table;
        std::vector<int> res;

        for (int i=0; i<nums.size(); i++) {
            table[nums[i]] = i;
        }

        for (int i=0; i<nums.size(); i++) {
            int diff = target - nums[i];

            if (table[diff] && table[diff] != i) {
                res.push_back(i);
                res.push_back(table[diff]);
                break;
            }
        }

        return res;
    }
};
