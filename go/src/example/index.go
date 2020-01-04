//https://www.youtube.com/watch?v=C8LgvuEBraI

package main

import (
	"fmt"
	"errors"
	"math"
)

func echo(){
	fmt.Println("Hello girl!")
}

func sum(){
	var x int
	x = 5
	var y int = 6
	var res int = x + y
	fmt.Println(res)
}

func ifelse(ops string){
	x := 5
	y := 6

	if ops == "+" {
		fmt.Println(x+y)
	} else if ops == "-" {
		fmt.Println(x-y)
	} else if ops == "*" {
		fmt.Println(x*y)
	} else if ops == "/" {
		fmt.Println(x/y)
	} else {
		fmt.Println("no ops found!")
	}
}

func array(){
	var a [5] int
	a[2] = 7

	b := [5]int{1,2,3,4,5} //Shorter version

	c := []int{1,2,3,4,5,6,7,8} //Shorter version

	d := []int{1,2,3}
	d = append(d, 4, 5, 6)

	fmt.Println(a, b, c, d)
}

func mapz(){
	vertices := make(map[string]string)

	vertices["fname"] = "John"
	vertices["lname"] = "Doe"
	vertices["username"] = "john1234"
	vertices["password"] = "********"

	fmt.Println(vertices)

	fmt.Println("Delete some data from map.")
	delete(vertices, "fname")
	fmt.Println(vertices)
}

func forz(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//Shorter version
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	arr := []string{"a", "b", "c"}
	for idx, val := range arr {
		fmt.Println("idx:", idx, "val:", val)
	}

	mapz := make(map[string]string)
	mapz["fname"] = "John"
	mapz["lname"] = "Doe"
	mapz["username"] = "john1234"
	mapz["password"] = "********"
	for key, val := range mapz {
		fmt.Println("key:", key, "val:", val)
	}
}

func sumArg(x int, y int) int{
	return x+y
}

func sqrt(x float64) (float64, error){
	if x < 0 {
		return 0, errors.New("Undefiend for negative numbers")
	}

	return math.Sqrt(x), nil
}

//Object
type person struct {
	name string
	age int
}

//Pointer
func inc(x *int){
	*x++
}

func main(){
	echo()
	sum()
	ifelse("*")
	array()
	mapz()
	forz()
	fmt.Println(sumArg(1, 2))

	//Sqrt
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//Object
	p := person{name: "Jane", age: 21}
	fmt.Println(p)
	fmt.Println(p.age)

	//Pointer
	i := 7
	inc(&i)
	fmt.Println(i)
}

