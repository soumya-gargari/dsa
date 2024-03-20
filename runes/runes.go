package main

import "fmt"

var a rune = 'b'
var b byte = 'b'

func main() {
	data := "Soumya"
	// if we are running normal for loop  and iterating over the string
	// we will get bytes value which will be uint8(means hold 8 bytes of data)
	for i := 0; i < len(data); i++ {
		fmt.Printf("type is %T\n", data[i])
	}
	fmt.Println()
	// if we are ranging over the string,
	// we will get rune value which will be int32(means hold 32 bytes of data)
	for _, val := range data {
		fmt.Printf("type is %T\n", val)
	}
	fmt.Println(a, b)
}
