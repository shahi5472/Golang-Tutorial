package main

import "fmt"

func main() {
	var m1 int = 2

	var (
		m2 = 3
		m3 = 4
	)

	var a int32 = 2

	var b int64 = 3

	fmt.Println(m1 + m2 + m3)

	fmt.Println(a)

	fmt.Println(int64(a) + b)

	fmt.Println(a + int32(b))

	// := use initizalize variable
	i := 2
	j := 2

	fmt.Println(i + j)

	var name string = "hello"

	fmt.Println(name)
}

/*
Atomic data type
string
int
int32
int64
uint
uint32
uint64

Unsafe
Pointer

Abstract data type
map[]<datatype>
struct{}
interface{}
*/
