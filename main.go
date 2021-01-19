package main

import (
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

//enreg or Struct

type Person struct {
	age      int
	name     string `required max:100` //just a tag go dont know about anything if u want to get the tag use "reflect" package
	lastname string
}

type Professor struct {
	Person //works like inheritance but its called Composition
	job    string
}

func main() {
	var i int = 1
	j := 3.15
	var code string = string(i) //it returns a unicode a the integer

	var s_i string = strconv.Itoa(i) // convert number to string
	//var s_i string = fmt.Sprintf("%d", i) 2nd method
	s_j := fmt.Sprintf("%.2f", j) // convert number to string
	fmt.Println("hello world", i, code, s_i, s_j)
	ok := true
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

	list := [...]int{0, 1, 2} //array of integers size = 3
	var list1 [3]string

	//list2 := list1 //creates a new array
	list2 := &list1 // pointes to the same array
	list1[0] = "test"
	fmt.Println(list)
	fmt.Println(list1, len(list1), list2)

	//Slices

	slice := []int{1, 2, 3, 5, 6} //or  make([]int, length, capacity)
	fmt.Println(slice)

	slice1 := []int{} //dynamic array its not so good because it goes whith the capacity of power 2

	slice1 = append(slice1, slice...) // push
	slice1 = append(slice1, 8)        // push
	fmt.Println(slice1)

	slice1 = append(slice1[:2], slice1[3:]...) // to remove the element of the index =2
	fmt.Println(slice1)

	//Maps = Dicts

	//array is a valid key type but Slices are NOT

	class := map[string]int{ // firt type is the key type qnd the 2nd is the value type
		"aziz":  15,
		"salah": 12,
		"test2": 13,
	}
	fmt.Println(class)
	class["test"] = 1
	delete(class, "test2") //delete a specif key value pair
	fmt.Println(class["aziz"], class)
	pop, exist := class["not found"] //if u want only the exist value u can use the throwaway _ exple= _, exist :=
	fmt.Println(pop, exist)

	//Struct

	person := Person{
		age:      20,
		name:     "aziz",
		lastname: "chagour",
	}
	fmt.Println(person, person.age)
	// ananymous strut
	p := struct{ name string }{name: "salah"}
	fmt.Println(p)
	p1 := p
	p1.name = "aziz"
	fmt.Println(p, p1) //it doesnt point to the struct

	pr := Professor{} //creating a professor instance

	pr.name = "swila7"
	pr.age = 52
	pr.lastname = "mendri"
	pr.job = "9awad"
	fmt.Println(pr)
}
