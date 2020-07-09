package array

//MoveZeroes moves all zeros to the end of the array
func MoveZeroes(nums []int) {
	// Two Pointers Time O(n), Space O(1)
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[j] != 0 {
			j++
			continue
		}
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}