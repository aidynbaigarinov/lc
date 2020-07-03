package array

// SingleNumber using HashMap
// Time O(n), Space O(n)
func SingleNumber(nums []int) int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
// SingleNumberBitManipultaion using bit manipultaion
// Time O(n), Space O(1)
func SingleNumberBitManipultaion(nums []int) int {
	var res int

	for num := range nums {
		res ^= num
	}
	return res
}