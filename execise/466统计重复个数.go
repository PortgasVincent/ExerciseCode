package main

func main466() {
	// fmt.Printf("%v", getMaxRepetitions("acb", 4, "ab", 2))
	// fmt.Printf("%v", getMaxRepetitions("aaa", 3, "aa", 1))
}

// func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int)int{
// 	c := getIncludeCount(s1, s2)

// 	checkSpecial := ""
// 	length := len([]rune(s2))
// 	for i:=0; i<length; i++ {
// 		checkSpecial += s1
// 	}
// 	specialCount := getIncludeCount(checkSpecial, s2)
// 	special := specialCount > length * c

// 	if n1 == 1 || !special{
// 		return c * n1 / n2
// 	}

// 	if specialCount == length * c + length - 1{
// 		return (n1 * c + (n1 - 1))/n2
// 	}

// 	hasSpecialNum := n1 / length
// 	rest := n1 % length

// 	str1 := s1 + s1
// 	neighborhood := getIncludeCount(str1, s2) == 2 * c + 1
// 	if neighborhood {
// 		restCount := rest * c + rest - 1
// 		return specialCount * hasSpecialNum + restCount + 1 + restCount * c
// 	}
	
// 	return specialCount * hasSpecialNum + rest * c
// }

// func getIncludeCount(a string, sub string) int{
// 	b := []rune(a)
// 	s := []rune(sub)
// 	count := 0
// 	for i, j := 0, 0; j < len(b); {
// 		if b[j] == s[i]{
// 			i++
// 			if i == len(s){
// 				count ++
// 				i = 0
// 			}
// 		}
// 		j++
// 	}
// 	return count
// }