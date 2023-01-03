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

func test1() {
	a := 1
	b := 2

	swap(a, b)
	fmt.Printf("a = %d,b = %d\n", a, b)
	swapPointer(&a, &b)
	fmt.Printf("a = %d,b = %d\n", a, b)
}

type SomeStructLevel2 struct {
	value int
}
type SomeStruct struct {
	value  int
	level2 SomeStructLevel2
}

func ChangeStructValue(s SomeStruct) {
	s.value = 2
}

func ChangeStructValuePointer(s *SomeStruct) {
	s.value = 3
}

func test2() {
	a := SomeStruct{value: 1}

	ChangeStructValue(a)
	fmt.Printf("a value : %d \n", a.value)
	ChangeStructValuePointer(&a)
	fmt.Printf("a value : %d \n", a.value)
}

func (s SomeStruct) ChangeStructValue() {
	s.value = 4
}

func (s *SomeStruct) ChangeStructValuePointer() {
	s.value = 5
}
func test3() {
	a := SomeStruct{value: 1}
	a.ChangeStructValue()
	fmt.Printf("a value : %d \n", a.value)
	a.ChangeStructValuePointer()
	fmt.Printf("a value : %d \n", a.value)

}

func test4() {
	a := SomeStruct{value: 1}
	b := a

	c := &a

	a.ChangeStructValuePointer()
	fmt.Printf("a value %d b value %d\n", a.value, b.value)
	fmt.Printf("a value %d c value %d\n", a.value, c.value)

	//	a.level2 := SomeStructLevel2{value : 10}
	a.level2 = SomeStructLevel2{value: 10}
	fmt.Printf("a value %d b value %d\n", a.level2.value, b.level2.value)
	fmt.Printf("a value %d c value %d\n", a.level2.value, c.level2.value)

}

func main() {
	test1()
	test2()
	test3()
	test4()
}
