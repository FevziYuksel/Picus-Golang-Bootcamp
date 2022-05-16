package main

func main() {

	//// Primitive types:
	/*uint8 the set of all unsigned 8-bit integers (0 to 255)

	  uint16 the set of all unsigned 16-bit integers (0 to 65535)

	  uint32 the set of all unsigned 32-bit integers (0 to 4294967295)

	  uint64 the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	  int8 the set of all signed 8-bit integers (-128 to 127)

	  int16 the set of all signed 16-bit integers (-32768 to 32767)

	  int32 the set of all signed 32-bit integers (-2147483648 to 2147483647)

	  int64 the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	  float32 the set of all IEEE-754 32-bit floating-point numbers

	  float64 the set of all IEEE-754 64-bit floating-point numbers

	  complex64 the set of all complex numbers with float32 real and imaginary parts

	  complex128 the set of all complex numbers with float64 real and imaginary parts

	  byte alias for uint8 rune alias for int32
	*/
	// Boolean Types
	// var flag bool //default value is false
	// fmt.Println(flag)

	// var patika bool = true
	// fmt.Println(patika)

	// patika = "patika" //as type bool in assignment error

	// Integer Types
	// var num int // digit default value is 0
	//var num int64 = 10

	// Float
	//var num1 float32 = 10.2
	//var num2 float64 = 10.2

	// var vs :=  //assignment methods
	//var x int = 10
	//var x = 10
	//var x int
	//fmt.Println(x)

	//var x = 10
	//var y float64 = 20
	//y:=20

	// var x, y int = 10, 20 	//multi assignment and unpackin in go
	// var x, y = 10, "patika"

	// var ( //multiline assignment (works for const too)
	// 	x    int
	// 	y        = 20 //auto assignment is possible in multiline
	// 	z    int = 30
	// 	d, e     = 40, "hello"
	// 	f, g string
	// )
	// fmt.Print(x, y, z, d, e, f, g)

	//const
	//const pi = 3
	//fmt.Println(pi)
	//pi = 4
	//fmt.Println(pi)

	//const x float64 = 10
	//
	//const max_value = 1
	//const MAX_VALUE = 1 //use this case
	//
	//const MAXVALUE = 1
	//const maxValue = 1

	////Composite Types // arrays copy and  slices,maps are referance type

	//Arrays //Static Array
	//var x [3]int
	// var y = [3]int{10, 20, 30}
	// y := [3]int{10, 20, 30}
	//y = append(y, 40) // doesn't work for slices

	// var x = []int{10, 20, 30} //slice = Linked list
	// var x [][]int
	// x = append(x[1:], 40) //  new slice = any slice + other slice
	// x = append(x[1:], 40)
	// x = append(x[:2], 40)

	// Slices
	// [...] -> array
	// var y = [...]int{10, 20, 30} //auto count
	// [] -> slice

	//x := []int{10, 20, 30}
	//fmt.Println(len(x)) //length method of slice

	//var x []int
	//y := []int{40, 50, 60}
	//x = append(x, y...) //add slice to another slice
	//fmt.Println(x)

	//Capacity //Allocation limit of slice
	//var x []int
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 10)
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 20)
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 30)
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 40)
	//fmt.Println(x, len(x), cap(x))
	//x = append(x, 50)
	//fmt.Println(x, len(x), cap(x))

	//Make //create slice with default(0) value
	//x := make([]int, 5) // make(slice type , length of slice)
	//fmt.Println(x)
	//x = append(x, 99)
	//fmt.Println(x)

	// x := make([]int, 5, 10) // make(slice type , length of slice, capacity of slice)
	// fmt.Println(x)

	// Slicing Slices //create referance
	// x := []int{1, 2, 3, 4}
	// y := x[:2] //works exactly like python
	// fmt.Println(y)

	//z := x[1:]
	//fmt.Println(z)

	// e := x[:]
	// e := x
	// fmt.Println(e)

	//referance proof
	// e := x[:]
	// e[0] = 5
	// fmt.Println(e)
	// fmt.Println(x)

	//x := [4]int{1, 2, 3, 4}
	//y := x[:2]

	//// copy
	// x := []int{1, 2, 3, 4}
	// y := x //referance
	// y := make([]int, len(x)) //creates y =[]{0,0,0,0}
	// num := copy(y, x)        //clone of x , num is number of element copied
	// fmt.Println(y, num)

	//Maps
	//var x map[string]int
	//var x = map[string]int{}
	//x := map[string]int{}

	////commas has to be in multiline var , const ,map ,struct etc
	//teams := map[string][]string{
	//	"Ahmet": []string{"a", ""},
	//	"Ayşe":  []string{"e", "f", "g"},
	//}

	//teams["Sevde"] = []string{"a", "y"}
	//teams["Sevde"] = []string{"x", "o"}

	////fmt.Println(teams)

	//y := teams["Ayşe"]
	//y = append(y, "k")
	//teams["Ayşe"] = y
	//fmt.Println(teams)

	// ages := make(map[string]string, 10) //create empty map
	// fmt.Println(ages)
	// var ages map[string]string //create empty map idk the different
	// fmt.Println(ages)

	//// Reading and Writing a Map

	//totalWins := map[string]int{}
	//totalWins["Ayşe"] = 1
	//totalWins["Osman"] = 2
	//totalWins["Ayşe"]++ //map index is actually an int so it is possible like other languages
	//fmt.Println(totalWins)

	// The comma ok idiom
	// m := map[string]string{
	// 	"hello": "5",
	// 	"world": "0",
	// }

	// v, ok := m["goodbye"] //v is value , ok is isExist flag !!!!
	// v, ok := m["hello"]
	// fmt.Println(v, ok)

	// delete(m, "hello") //delete(map,key)
	// fmt.Println(m)

	//// Struct

	//Psudoclass of C and GO

	//"new" keyword is exist but rarely used return pointer of struct

	/*https://yourbasic.org/golang/structs-explained/*/

	// type person struct {
	// 	name string
	// 	age  int
	// 	pet  string
	// }

	// //Empty Structs
	// var ahmet person
	// ayse := person{} ("",0,"")

	// // Filling Structs indivually
	// osman := person{}
	// osman.age = 22
	// osman.pet = "turuncu"

	// // Contructor style
	// osman := person{
	// 	"Osman",
	// 	21,
	// 	"Boncuk",
	// }

	// // Spesified version Contructor style
	//osman := person{
	//	name: "Osman",
	//	age:  21,
	//	pet:  "Zeytin",
	//}

	// anonymous structs

	//Struct
	//type person struct {
	//	name string
	//	age  int
	//	pet  string
	//}
	//
	//var person struct {
	//	name string
	//	age  int
	//	pet  string
	//}

	//person.name = "Osman"
	//person.age = 21
	//person.name = "Zeytin"
	//
	//pet := struct {
	//	name string
	//	kind string
	//}{
	//	name: "Zeytin",
	//	kind: "Cat",
	//}
	//

	//
	//type firstPerson struct{
	//	name string
	//	age int
	//}
	//f = firstPerson{
	//	name: "Osman",
	//	age:  21,
	//}
	//

	//type person struct {
	//	name string
	//	age  int
	//	pet  string
	//}
	//var g struct {
	//	name string
	//	age  int
	//}
	//g.age = 21
	//g.name = "test"
	//g = f
	//
	//
	//fmt.Println(f == g)
	//

	// Shadowing Variables

	//x := 10
	//if x > 5 {
	//	fmt.Println(x)
	//	x = 5
	//	fmt.Println(x)
	//}
	//fmt.Println(x)

	// if
	//n := 10
	//if n == 0{
	//	//
	//}else if n == 1{
	//
	//}else{
	//
	//}

	// for
	// c-style
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	////condition-only
	//i := 1
	//for i < 100 {
	//	fmt.Println(i)
	//	i = i * 2
	//}
	//
	//// infinite
	//for {
	//	fmt.Println("HELLO")
	//}
	//
	//break
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//	if i == 5{
	//		fmt.Println("finished")
	//		break
	//	}
	//}

	//continue
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//	if i == 5{
	//		fmt.Println("finished")
	//		continue
	//	}
	//}

	//for-range
	//evenVals := []int{2, 4, 6, 8, 10}
	//for i, v := range evenVals {
	//	fmt.Println(i,v)
	//}

	//evenVals := []int{2, 4, 6, 8, 10}
	//for _, v := range evenVals {
	//	fmt.Println(v)
	//}

	// iterate
	//m := map[string]int{
	//	"a": 1,
	//	"c": 3,
	//	"b": 2,
	//}
	//
	//for i := 0; i < 3; i++ {
	//	fmt.Println(m)
	//
	//	fmt.Println("Loop", i)
	//	for k, v := range m {
	//		fmt.Println(k, v)
	//	}
	//}

	//evenVals := []int{2, 4, 6, 8, 10}
	//for _, v := range evenVals {
	//	v *= 2
	//}
	//fmt.Println(evenVals)

	//
	//words := []string{"a","cow","smile","gopher", "octopus"}
	//for _, word := range words{
	//	switch word{
	//	case "a":
	//		fmt.Println("its a")
	//	case "cow":
	//		fmt.Println("cow")
	//	default:
	//		fmt.Println("Default value")
	//	}
	//}

}
