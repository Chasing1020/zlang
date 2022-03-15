// Given an array of integers nums and an integer target
// return indices of the two numbers such that they add up to target

let nums = [2, 7, 11, 15];
let target = 9;
let dict = {};
for (let i = 0; i < len(nums); i = i + 1) {
   if (dict[nums[i]] == null) {
       dict[target - nums[i]] = i;
   } else {
       printf("Answer found: %s %s", i, dict[nums[i]]);
   }
}
