package main

func main() {

}

func strStr(haystack string, needle string) int {
	l := len(needle)
	for i := 0; i < len(haystack); i++ {
		if i+l <= len(haystack) && haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}
