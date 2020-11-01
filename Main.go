package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	// defer function
	fmt.Println("start")
	fmt.Println("middle")
	fmt.Println("end")

	fmt.Println("real world defer example")

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	// defer
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	deferstart := "start"
	defer fmt.Println(deferstart)
	deferstart = "end"

	// panic = error handling, executes after defer
	// aa, bb := 1, 0
	// ans := aa / bb
	// defer fmt.Println(ans)

	// Pointers
	var pointerVariable int = 42
	// memory address (pointer to integer)
	var copiedValue *int = &pointerVariable
	fmt.Println(pointerVariable, *copiedValue)
	pointerVariable = 27
	fmt.Println(pointerVariable, *copiedValue)
	*copiedValue = 1234
	fmt.Println(pointerVariable, *copiedValue)

	// Pointer arithmetic = NOT ALLOWED
	// aaa := [3]int{1, 2, 3}
	// bbb := &aaa[0]
	// ccc := &aaa[1] - 4
	// fmt.Printf("%v %p %p\n", aaa, bbb, ccc)

	// pointer to struct
	type myStruct struct {
		foo int
	}

	// nil value. get/set
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Println((*ms).foo) // 42

	var ms2 *myStruct
	ms2 = new(myStruct)
	ms2.foo = 42
	fmt.Println(ms2.foo) // 42

	// slice points to the first variable of the underlying array
	// map contains a pointer to the underlying data too
	aaa1 := []int{1, 2, 3}
	bbb1 := aaa1
	fmt.Println(aaa1, bbb1)
	aaa1[1] = 42
	fmt.Println(aaa1, bbb1)

	// functions
	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	sum := sum(1, 2, 3, 4, 5)
	var msg string = "The sum is"
	fmt.Println(msg, *sum)
	sum2 := sumWithNamedReturn(1, 2, 3, 4, 5)
	fmt.Println(msg, sum2)

	ddd, err := devide(5.0, 0.1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ddd)

	// Anonymous function
	func() {
		fmt.Println("Hello go")
	}()

	var function func() = func() {
		fmt.Println("Hello go!")
	}
	function()

	// Method = function in known context (inside a type)

	ggg := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	ggg.greet()

	// Interfaces
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go Interface"))

	// Any type interface
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	// empty interface can be any type
	var emptyInterface interface{}
	emptyInterface = "emptyInterface"
	fmt.Println(emptyInterface)

}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*greeting, *name)
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func sumWithNamedReturn(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	// promoting the variable from the local stack of the function
	// to global heap (shared memory) of the computer
	// (to use this pointer futher)
	return
}

func devide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot devide by zero")
	}
	return a / b, nil
}
