package recursion

// func main() {

// 	nums := make([]int, 5)
// 	nums[0] = 1
// 	nums[1] = 5
// 	nums[2] = 7
// 	nums[3] = 3
// 	nums[4] = 6

// 	res := Sum(nums)
// 	fmt.Println(res)
// 	nums = append(nums, 1, 8)
// 	fmt.Println(Sum(nums))
// }

//Sum returns sum of all numbers in a input slice
func Sum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return nums[0] + Sum(nums[1:])
}
