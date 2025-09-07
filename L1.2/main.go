package main

import(
	"fmt"
)

func main(){
	array := []int{2, 4, 6, 8, 10}
	result := multiply(array)

	fmt.Println(result)
}
func multiply(n []int) []int{
	for i, elem := range n{
		n[i] = elem * elem
	}
	return n
}