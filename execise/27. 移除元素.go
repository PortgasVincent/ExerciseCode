package main

func main() {

}

func removeElement(nums []int, val int) int {
	p, q, l := 0, 0, 0
	for q < len(nums) {
		if nums[q] != val {
			l++
			if q > p {
				nums[p] = nums[q]
			}
			p++
			q++
			continue
		}
		q++
	}
	return l
}
