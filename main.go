package main

import (
	"fmt"
	"log"
	"rgr/lib"
)

func main() {
	data, err := lib.ParseData("test.json")
	if err != nil {
		log.Fatal(err)
	}
	res := make([]float32, len(data.A))
	for i := 0; i < len(data.A); i++ {
		fmt.Println("Data:")
		fmt.Printf("a = %v, b = %v\n", data.A[i], data.B[i])
		fmt.Println()
		res[i], err = lib.Count(data.A[i], data.B[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()

	}
	tree := &lib.BinaryTree{}
	for i := 0; i < len(res); i++ {
		tree.Insert(res[i])
	}
	fmt.Print("Resulting numbers:", res)
	fmt.Println()
	fmt.Println("Binary tree:")
	tree.GetTree()

	var s string
	fmt.Println("Do you want to add Node [y/n]?")
	fmt.Scanln(&s)
	if s == "y" {
		fmt.Println("Enter value for a\n a > 0:")
		var a float32
		var b float32
		fmt.Scanln(&a)
		fmt.Println("Enter value for b:")
		fmt.Scanln(&b)
		var r, err = lib.Count(a, b)
		if err != nil {
			log.Fatal(err)
		}
		tree.Insert(r)
		fmt.Println("Binary tree:")
		tree.GetTree()
	} else {
		fmt.Println("Bye!")
	}
}
