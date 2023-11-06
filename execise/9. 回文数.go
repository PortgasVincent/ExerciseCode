package main

func main() {

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	x1 := x
	var reverse int
	for x1 != 0 {
		reverse = reverse*10 + x1%10
		x1 /= 10
	}
	return reverse == x
}
