package main

import (
	"fmt"
)

func main() {
	fmt.Println(fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16))
}

func fullJustify(words []string, maxWidth int) []string {
	l := len(words)
	var res []string
	var p, q, rest, space, spacei int = 0, 0, maxWidth, 0, -1

	for q <= l-1 {
		// 第一个没有空格
		if q == p && len(words[q]) <= rest {
			rest = rest - len(words[q])
			q++
			continue
		}
		if (len(words[q]) + 1) <= rest {
			rest = rest - len(words[q]) - 1
			q++
			continue
		}

		var str = ""
		spacex := q - p - 1
		if spacex == 0 {
			space = rest
			str += words[p]
			for j := 0; j < space; j++ {
				str += " "
			}
		} else {
			space = rest / spacex
			spacei = rest % spacex
			str += words[p]
			for i := p + 1; i < q; i++ {
				// ==space是因为除了rest，加word时也减去了一个space
				for j := 0; j <= space; j++ {
					str += " "
				}
				if spacei+p >= i {
					str += " "
				}
				str += words[i]
			}
		}

		res = append(res, str)
		// reset
		p = q
		rest = maxWidth
		space, spacei = 0, -1
	}
	// last
	var str = ""
	str += words[p]
	for i := p + 1; i < q; i++ {
		str += " "
		str += words[i]
	}
	for j := 0; j < rest; j++ {
		str += " "
	}
	res = append(res, str)

	return res
}
