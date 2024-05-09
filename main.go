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
	fmt.Printf("a = %v, b = %v\n", data.A, data.B)
	fmt.Println()
	res := make([]float32, len(data.A))
	for i := 0; i < len(data.A); i++ {
		res[i], err = lib.Count(data.A[i], data.B[i])
		if err != nil {
			fmt.Printf("Something went wrong")
		}
	}

	tree := &lib.Node{Value: res[0]}

	for i := 1; i < len(res); i++ {
		tree.Insert(res[i])
	}
	fmt.Print("Binary tree:\n")
	lib.PrintTree(tree, 0)
}
