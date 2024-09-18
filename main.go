// 1} simple hello prog

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello")
// }

// 2} Return multiple values from return
package main

import "fmt"

func add_remove() (int, int) {
	x := 9
	y := 6
	m := x + y
	n := x - y
	return m, n
}

func main() {
	p, o := add_remove()
	fmt.Println(p, o)
}
