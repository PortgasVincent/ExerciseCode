package leetcode

func main(){

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	index1 := (m + n + 1)/2
	index2 := (m + n + 2)/2
	return float64(getKth(nums1, 0, m-1, nums2, 0, n-1, index1) + getKth(nums1, 0, m-1, nums2, 0, n-1, index2)) * 0.5
}

func getKth(arr1 []int, left1, right1 int, arr2 []int, left2, right2, k int) int {
	len1 := right1 - left1 + 1
	len2 := right2 - left2 + 1
	if len1 > len2 {
		return getKth(arr2, left2, right2, arr1, left1, right1, k)
	}
	if len1 == 0 {
		return arr2[left2 + k - 1]
	}
	if k == 1{
		return min(arr1[left1], arr2[left2])
	}
	i := left1 + min(k/2, len1) - 1
	j := left2 + min(k/2, len2) - 1
	if arr1[i] < arr2[j] {
		return getKth(arr1, i+1, right1, arr2, left2, right2, k-min(k/2, len1))
	}
	return getKth(arr1, left1, right1, arr2, j+1, right2, k-min(k/2, len2))
}

func min(a, b int)int{
	if a >= b{
		return b
	}
	return a
}