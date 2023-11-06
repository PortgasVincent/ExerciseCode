package main

func main() {

}

func removeDuplicates(nums []int) int {
	m := map[int]bool{}
	l := len(nums)
	for i := 0; i < l; {
		if m[nums[i]] {
			count := 1
			for k := i + 1; k < l; k++ {
				if nums[k] == nums[k-1] {
					count++
				} else {
					break
				}
			}
			for j := i; j < l-count; j++ {
				nums[j] = nums[j+count]
			}
			l -= count
			continue
		}
		m[nums[i]] = true
		i++
	}
	return l
}
