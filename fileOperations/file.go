package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	buf := make([]byte, 100)
	for {
		data, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if data > 0 {
			fmt.Printf("%v\n", string(buf[:data]))
			fmt.Println("================")
		}
	}
}
