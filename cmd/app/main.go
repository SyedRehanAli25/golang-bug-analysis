package main

import (
	"fmt"
	"github.com/yourusername/golang-bug-analysis/pkg/example"
)

func main() {
	fmt.Println("Starting GoLang Bug Analysis POC")
	result := example.Add(5, 3)
	fmt.Println("Example Add Result:", result)
}
