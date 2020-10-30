package main

import (
	"fmt"
	"reflect"
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

	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 0011
	fmt.Println(a ^ b)  // 1001
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

	// arrays
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)
	grades1 := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades1)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students[0])
	fmt.Printf("Number of students: %v\n", len(students))

	// arrays = values, not pointers to values in memory
	array1 := [...]int{1, 2, 3}
	array2 := array1
	array3 := &array1
	array1[1] = 5
	fmt.Println(array1) // 1, 5, 3
	fmt.Println(array2) // 1, 2, 3
	fmt.Println(array3) // 1, 5, 3
	// slices can come from array/slice
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Length of slice1: %v\n", len(slice1))
	fmt.Printf("Capacity of slice1: %v\n", cap(slice1))
	slice2 := slice1
	fmt.Println(slice2) // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	slice3 := slice1[:] // 1, 2, 3, 4, 5, 6, 7, 8, 9, 10
	fmt.Println(slice3)
	slice4 := slice1[3:] // 4, 5, 6, 7, 8, 9, 10
	fmt.Println(slice4)
	slice5 := slice1[:6] // 1, 2, 3, 4, 5, 6
	fmt.Println(slice5)
	slice6 := slice1[3:6] // 4, 5, 6
	fmt.Println(slice6)

	// make function
	madeSlice := make([]int, 3, 100)
	fmt.Println(madeSlice)
	fmt.Printf("Length of madeSlice: %v\n", len(madeSlice))
	fmt.Printf("Capacity of madeSlice: %v\n", cap(madeSlice))

	// append function
	appendedSlice := []int{}
	fmt.Println(appendedSlice)
	fmt.Printf("Length of madeSlice: %v\n", len(appendedSlice))
	fmt.Printf("Capacity of madeSlice: %v\n", cap(appendedSlice))
	appendedSlice = append(appendedSlice, 1)
	fmt.Println(appendedSlice)
	fmt.Printf("Length of madeSlice: %v\n", len(appendedSlice))
	fmt.Printf("Capacity of madeSlice: %v\n", cap(appendedSlice))
	appendedSlice = append(appendedSlice, 1, 2, 3, 4, 5)
	fmt.Println(appendedSlice)
	fmt.Printf("Length of madeSlice: %v\n", len(appendedSlice))
	fmt.Printf("Capacity of madeSlice: %v\n", cap(appendedSlice))

	// spreading second array values while appending this array items to first array
	appendedSlice = append(appendedSlice, []int{6, 7, 8, 9, 10}...)
	fmt.Println(appendedSlice)
	fmt.Printf("Length of madeSlice: %v\n", len(appendedSlice))
	fmt.Printf("Capacity of madeSlice: %v\n", cap(appendedSlice))

	// stack of slices
	stackedSlice1 := []int{1, 2, 3, 4, 5}
	// all values copied except for that last
	stackedSlice2 := stackedSlice1[:len(stackedSlice1)-1]
	fmt.Println(stackedSlice2)
	// removing element from the middle of stackedSlice1
	stackedSlice3 := append(stackedSlice1[:2], stackedSlice1[3:]...)
	fmt.Println(stackedSlice3)

	// changed stackedSlice1
	fmt.Println(stackedSlice1)

	// map data type: key-value
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"Cali":    1234234,
		"Texas":   25435234,
		"Florida": 2432432,
	}
	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 13123434
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	// check if the key-value pair exist
	pop, ok := statePopulations["Test"]
	fmt.Println(pop, ok) // 0 false

	// struct as a collection type (data type)
	type Doctor struct {
		number     int
		actorName  string
		episodes   []string
		companions []string
	}

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon",
		companions: []string{
			"Companion1",
			"Companion2",
			"Companion3",
		},
	}
	fmt.Println(aDoctor)

	// anonymous struct
	bDoctor := struct{ name string }{name: "John"}
	anotherDoctor := bDoctor
	anotherDoctor.name = "Tom"
	fmt.Println(bDoctor)
	fmt.Println(anotherDoctor)

	// embedded. Composition > Inheritance

	type Animal struct {
		Name   string `required max: "100"`
		Origin string
	}

	type Bird struct {
		Animal
		Speed  float32
		CanFly bool
	}

	birdy := Bird{
		Animal: Animal{
			Name:   "Birdy",
			Origin: "Sweden",
		},
		Speed:  50,
		CanFly: true,
	}
	fmt.Println(birdy)

	// Validation with Tags
	typee := reflect.TypeOf(Animal{})
	field, _ := typee.FieldByName("Name")
	fmt.Println(field.Tag)

	// if/else statements

	if true {
		fmt.Println("The test is true")
	}

	if pop, ok := statePopulations["Texas"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 105
	if guess < number || guess < 50 {
		fmt.Println("Too low")
	} else if guess >= 1 && guess <= number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("Got it!")
	}

	fmt.Println(number <= guess, number >= guess, number != guess)

	// short circuiting: as soon as 1 part of OR test is true, the code executed

	// switch statement, no break

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// standalone case
	y := 10
	switch {
	case y <= 10:
		fmt.Println("<=10")
		fallthrough
	case y <= 20:
		fmt.Println("<=20")
	default:
		fmt.Println(">20")
	}

	// type switch

	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		// break statement
		break
		fmt.Println("this will not print")
	case float64:
		fmt.Println("i is an float64")
	case string:
		fmt.Println("i is an string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")
	}

	// looping for statement
	// first form of loop
	// for init; test; incrementer {}
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Println(i)
		fmt.Println(j)
	}

	looper := 0
	for ; looper < 5; looper++ {
		fmt.Println(looper)
	}
	fmt.Println(looper + 5)

	// second form of loop
	// for test
	z := 0
	for z < 5 {
		fmt.Println(z)
		z++
	}

	// third form of loop
	// for {}
	// cb := 0
	// for {
	// 	fmt.Println(cb)
	// 	if cb == 5 {
	// 		break
	// 	}
	// }

	// continue statement
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// break {LABEL} in any place

	// for range loop
	rangee := []string{"first", "second", "third"}
	for _, v := range rangee {
		fmt.Println(v)
	}
}
