package main

import "fmt"

// nolint: gosimple // <- DO NOT delete this comment!
func main() {
	var r rune // write the data type after 'r'
	r = 'a'

	var i int // write the data type after 'i'
	i = 1

	var b bool // write the data type after 'b'
	b = true

	fmt.Printf("%c %d %t", r, i, b) // do not remove this line
}
