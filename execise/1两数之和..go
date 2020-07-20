package main

func main1() {
	
}

func getSumIndex(arr []int, sum int) []int{
	m := map[int]int{}
	for i := 0; i<len(arr); i++{
		index, ok := m[sum - arr[i]]
		if ok {
			return []int{i, index}
		}
		m[arr[i]] = i
	}
	return nil
}