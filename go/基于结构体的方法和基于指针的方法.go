package main

import "fmt"

func main() {
	a := cache{}
	a.update(11, "struct")
	fmt.Println(a.A, a.B)
	p := &a
	p.put(222, "point")
	fmt.Println(p.A, p.B)
	fmt.Println(a.A, a.B)
}

type cache struct {
	A int
	B string
}

func (c cache) update(a int, b string) {
	c.A = a
	c.B = b
}

func (c *cache) put(a int, b string) {
	c.A = a
	c.B = b
}
