package main

import (
    "fmt"
    "github.com/your-organization/basics-go/slices"
)

func main() {
    fmt.Println("Testing basics-go")

    // test all the packages here 
	original := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", original)
	reversed :=	slices.ReverseElements(original)
	fmt.Println("Reversed:", reversed)

   
}
