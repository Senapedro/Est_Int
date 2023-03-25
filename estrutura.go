package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(pessoa{"Ana", 25})
	fmt.Println(pessoa{"joao", 20})
}
