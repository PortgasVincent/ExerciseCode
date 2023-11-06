package main

// func getSumIndex(arr []int, sum int) []int{
// 	m := map[int]int{}
// 	for i := 0; i<len(arr); i++{
// 		index, ok := m[sum - arr[i]]
// 		if ok {
// 			return []int{i, index}
// 		}
// 		m[arr[i]] = i
// 	}
// 	return nil
// }

func main1() {

}

func twoSum(nums []int, target int) []int {
	var m = map[int]int{}
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = i
	}
	return nil
}
