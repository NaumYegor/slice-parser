package main

import (
	"fmt"
	"github.com/naumyegor/slice-parser/slicing"
	"log"
)

func main() {
	slices := [][]int {
		{100, 15, 27},
		{30, 40, 207},
		{400, 400, 1},
	}
	m, err := slicing.NewMatrix(slices)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(m.Transform())
}
