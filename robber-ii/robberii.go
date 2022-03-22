package robberii

// https://leetcode.com/problems/house-robber-ii/

// Examples
// 1, 3, 1    = 3
// 1, 2, 3, 1 = 4
// 1, 2, 3    = 3

// constraints:
// 1 <= len(nums) <= 100
// 0 <= nums[i] <= 1000

func Rob(nums []int) int {
    count := len(nums)
    if count == 1 {
        return nums[0]
    }

    if count == 2 {
        return max(nums[0], nums[1])
    }

    x := robAlternate(nums[0:count-1])
    y := robAlternate(nums[1:count])

    return max(x,y)
}

func robAlternate(nums []int) int {
    var money, rob1, rob2 int
    for _, m := range nums {
        money = max(rob1 + m, rob2)
        rob1 = rob2
        rob2 = money
    }
    return money
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
