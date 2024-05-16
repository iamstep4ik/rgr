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
	fmt.Println("Data:")
	for i := 0; i < len(data.A); i++ {
		fmt.Printf("a = %v, b = %v\n", data.A[i], data.B[i])
		fmt.Println()
		res[i], err = lib.Count(data.A[i], data.B[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println()

	}
	tree := &lib.Node{Value: res[0]}
	for i := 1; i < len(res); i++ {
		tree.Insert(res[i])
	}
	fmt.Printf("Resulting numbers: %f", res)
	fmt.Println()
	fmt.Println("Binary tree:")
	lib.PrintTree(tree, 0)
}
