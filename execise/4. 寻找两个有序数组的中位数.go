package main

func main(){

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	if len1 <= 2 && len2 <= 2 {
		return getMiddle(nums1 []int, nums2 []int)
	}

	var mid1, mid2 float64 
	var index1, index2 int
	if len1 % 2 == 0 {
		mid1 = (nums1[len1/2-1] + nums1[len1/2])/2
		index1 = len1/2 - 1
	} else {
		mid1 = nums1[(len1-1)/2]
		index1 = (len1 - 1)/2
	}
	if len2 % 2 == 0 {
		mid2 = (nums2[len2/2-1] + nums2[len2/2])/2
		index2 = len2/2 - 1
	} else {
		mid2 = nums2[(len2-1)/2]
		index2 = (len2 - 1)/2
	}
	
	
}



// func getMiddle(a, b, c int)int {
// 	if a > b{
// 		if a < c {
// 			return a
// 		}
// 		if b > c{
// 			return b
// 		}
// 		return c
// 	}
// 	if a > c {
// 		return a
// 	}
// 	if b > c{
// 		return c
// 	}
// 	return b
// }