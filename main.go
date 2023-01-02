package main

import "fmt"

func swap(a, b int) {
	temp := a
	a = b
	b = temp
}

func swapPointer(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 1
	b := 2

	swap(a, b)
	fmt.Printf("a = %d,b = %d\n", a, b)
	swapPointer(&a, &b)
	fmt.Printf("a = %d,b = %d\n", a, b)

}
