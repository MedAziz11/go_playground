package main

import(
	"fmt"
	"strconv"
)
// uppercase variable is visible and exported to other packages
// lowercase vars are visible to any file in the same package

const (
	_ = iota // _ read only val : read it and throw the val
	a 
	b 
	c // automatically enumerate the consts
)

func main(){
	var i int = 1
	j := 3.15;
	var code string = string(i) //it returns a unicode a the integer

	var s_i string = strconv.Itoa(i) // convert number to string
	//var s_i string = fmt.Sprintf("%d", i) 2nd method
	s_j := fmt.Sprintf("%.2f", j)   // convert number to string
	fmt.Println("hello world",i, code, s_i, s_j)
	ok:= true
	fmt.Printf("%v, %T\n", ok, ok) //%v value, %T type

	var k int // automatically is a 0 
	fmt.Println(k)

	s := "im a string"
	fmt.Println(s)
	fmt.Println(s[3]) // returns the unicode of the character
	fmt.Println(string(s[3]))
	fmt.Println([]byte(s)) // returns an array of bytes of each character


	const cC int = 5 // constant
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	//ARRAYS 

	list := [...]int{0,1,2} //array of integers size = 3
	var list1 [3]string
	
	//list2 := list1 //creates a new array
	list2 := &list1 // pointes to the same array
	list1[0] = "test"
	fmt.Println(list)
	fmt.Println(list1, len(list1), list2)

	//Slices

	slice := []int{1,2,3,5,6} //or  make([]int, length, capacity)
	fmt.Println(slice)

	slice1 := []int{} //dynamic array its not so good because it goes whith the capacity of power 2

	slice1 = append(slice1, slice...) // push
	slice1 = append(slice1, 8) // push
	fmt.Println(slice1)

	slice1 = append(slice1[:2], slice1[3:]...) // to remove the element of the index =2
	fmt.Println(slice1)

}
