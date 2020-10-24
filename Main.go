package main

import (
	"fmt"
	"strconv"
)

var (
	I    float32 = 42
	Name string
)

const ( // enumerator for enum
	iota1 = iota // equals 0
	iota2
	iota3
)

const ( // enumerator for enum
	_ = iota // write only constant (garbage collected)
	iota4
	iota5
)


func main() {
	fmt.Printf("%v, %T\n", I, I)
	j := 27
	fmt.Printf("%v, %T\n", j, j)
	Name = strconv.Itoa(j)
	fmt.Printf("%v, %T\n", Name, Name)
	boolean := true
	fmt.Printf("%v, %T\n", boolean, boolean)
	var n uint16 = 42
	fmt.Printf("%v, %T\n", n, n)

	a := 10 // 1010
	b := 3 // 0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 0011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b) // 0100
	fmt.Println(a >> 3)
	fmt.Println(a << 3)
	str := "this is a string"
	bytes := []byte(str)
	fmt.Printf("%v, %T\n", bytes, bytes)
	var roone rune = 'a' // int32 (uint8 is rune as well)
	fmt.Printf("%v, %T\n", roone, roone)

	const myConst int = 42 // only assignable at compile time, not run time
	fmt.Printf("%v, %T\n", myConst, myConst)
}
