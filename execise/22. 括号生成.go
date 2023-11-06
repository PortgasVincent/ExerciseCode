package main

func main() {

}

func generateParenthesis(n int) []string {
	strs := []string{"("}
	return generate(strs, 1, n-1)
}

func generate(strs []string, leftCount, n int) []string {
	if n == 0 && leftCount == 0 {
		return strs
	}
	if n == 0 && leftCount != 0 {
		var temp string
		for i := 0; i < leftCount; i++ {
			temp += ")"
		}
		for i := 0; i < len(strs); i++ {
			strs[i] = strs[i] + temp
		}
		return strs
	}
	if leftCount == 0 {
		for i := 0; i < len(strs); i++ {
			strs[i] = strs[i] + "("
		}
		leftCount++
		n--
		return generate(strs, leftCount, n)
	}

	// 加 "("
	strs1 := make([]string, len(strs))
	left1 := leftCount
	n1 := n
	for i := 0; i < len(strs); i++ {
		strs1[i] = strs[i] + "("
	}
	left1++
	n1--
	strs1 = generate(strs1, left1, n1)

	// 加 ")"
	strs2 := make([]string, len(strs))
	left2 := leftCount
	n2 := n
	for i := 0; i < len(strs); i++ {
		strs2[i] = strs[i] + ")"
	}
	left2--
	strs2 = generate(strs2, left2, n2)

	return append(strs1, strs2...)
}
