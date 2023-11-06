package main

func main() {

}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	len1 := len(nums1)
// 	len2 := len(nums2)
// 	count1 := (len1 + len2 + 1) / 2
// 	count2 := (len1 + len2 + 2) / 2
// 	return float64(getKth(nums1, 0, len1-1, nums2, 0, len2-1, count1)+getKth(nums1, 0, len1-1, nums2, 0, len2-1, count2)) * 0.5
// }

// func getKth(num1 []int, left1, right1 int, num2 []int, left2, right2, k int) int {
// 	if right1-left1 > right2-left2 {
// 		return getKth(num2, left2, right2, num1, left1, right1, k)
// 	}
// 	if right1-left1+1 <= 0 {
// 		return num2[left2+k-1]
// 	}
// 	if k == 1 {
// 		return min(num1[left1], num2[left2])
// 	}
// 	n1 := num1[min(left1+k/2-1, right1)]
// 	n2 := num2[min(left2+k/2-1, right2)]
// 	if n1 < n2 {
// 		return getKth(num1, min(left1+k/2-1, right1)+1, right1, num2, left2, right2, k-min(left1+k/2-1, right1)+left1-1)
// 	}
// 	return getKth(num1, left1, right1, num2, min(left2+k/2-1, right2)+1, right2, k-min(left2+k/2-1, right2)+left2-1)
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	count1 := (l1 + l2 + 1) / 2
	count2 := (l1 + l2 + 2) / 2
	return 0.5 * float64(getKth(nums1, 0, l1-1, nums2, 0, l2-1, count1)+getKth(nums1, 0, l1-1, nums2, 0, l2-1, count2))
}

func getKth(num1 []int, left1, right1 int, num2 []int, left2, right2, k int) int {
	if right1-left1 > right2-left2 {
		return getKth(num2, left2, right2, num1, left1, right1, k)
	}
	if right1-left1 < 0 {
		return num2[left2+k-1]
	}
	if k == 1 {
		return min(num1[left1], num2[left2])
	}
	n1 := num1[min(left1+k/2-1, right1)]
	n2 := num2[min(left2+k/2-1, right2)]
	if n1 < n2 {
		return getKth(num1, min(left1+k/2-1, right1)+1, right1, num2, left2, right2, k-min(left1+k/2-1, right1)+left1-1)
	}
	return getKth(num1, left1, right1, num2, min(left2+k/2-1, right2)+1, right2, k-min(left2+k/2-1, right2)+left2-1)
}
