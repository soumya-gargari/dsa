package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// generics used in structure
// below is the example

type CustomData interface {
	constraints.Ordered
}
type user[T CustomData] struct {
	name string
	age  int
	data T
}

// generics used in function
// below is the example
func Test[T constraints.Ordered](a []T, c func(T) T) []T {
	var datas []T
	for _, val := range a {
		data := c(val)
		datas = append(datas, data)
	}
	return datas
}

func Add[T constraints.Ordered](a, b T) T {
	return a + b
}

// generics used in map
// below is the example
type data[T constraints.Ordered, V constraints.Ordered] map[T]V

func main() {
	// using generics in structure
	result := user[float64]{
		name: "Soumya",
		age:  30,
		data: 90.34567890,
	}
	fmt.Printf("%v\n", result)

	// using generics in function
	result1 := Test([]float32{2.9, 3.3, 4.6, 5.2, 6.1}, func(i float32) float32 { return i * 2 })
	fmt.Println("result is:", result1)
	mapData := make(data[int, string])
	mapData[3] = "Soumya"
	mapData[5] = "Gargari"
	fmt.Println("map data is:", mapData)
	c := Add(1.3, 2.6)
	fmt.Println("Addition result is:", c)
}
