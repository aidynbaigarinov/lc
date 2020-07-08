package array

//MajorityElement returns element that appears in array more than len(nums)/2 times
func MajorityElement(nums []int) int {
	// Hashmap approach Time O(n), Space O(n)
	m := make(map[int]int)

	for i := range nums {
		m[nums[i]]++
	}

	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0

	// TODO: Look another solutions
}