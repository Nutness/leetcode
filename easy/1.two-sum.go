/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    arrayLength := len(nums)
    var firstValue int
    var secondValue int
	for i := 0; i < arrayLength; i++ {
		for j := i + 1; j < arrayLength; j++ {
			if (nums[i] + nums[j] == target) {
				firstValue = i
                secondValue = j
                break
			}
		}
	}
    return []int{firstValue, secondValue}
}
// @lc code=end

