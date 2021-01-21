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

	//NB: = is to increment a value to a variable that u already declared its type
	// := is to increment a value to a variable that u dont know its type
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

	// If and Switch statement

	if true { //condition
		// do smmthing
		fmt.Println("hello from if")
	}

	//if with an initiliazer

	if pop, exist := class["aaa"]; exist {
		fmt.Println(pop)
	}

	//Switch

	switch 2 {
	case 1:
		fmt.Println("thats 1") //u can add fallthrough key word to make it go to the next test even if it get the right result
		//fallthrough
	case 2:
		fmt.Println("thats 2")
	default:
		fmt.Println("None")

	}

	//Loops

	for i := 0; i < 5; i++ { //like C
		fmt.Println(i)
	}

	for i, j := 0, 5; i < 5; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}

	//Labels in Looops
	fmt.Println("************")
Loop:
	for i := 1; i < 5; i++ {
		for j := 1; j < 3; j++ {
			// if we need it to break if i*j >3 then we use the labels
			// because the break will only break the loop that it is in

			if i*j >= 3 {
				break Loop
			}
			fmt.Println(i * j)

		}

	}
	//Loop with range function (the python for 🤩)

	for i, v := range class {
		fmt.Println(i, v) //i : index, v: value

	}

	//defer is special when u put is on a finction it will run the last thing
	// defer uses a LIFO means the last defer function will run the first
	// defer some_function()

	//panic or Exceptions

	// i = 0
	// if i == 0 {
	// 	panic("cant devide by zero")
	// } else {
	// 	fmt.Println("nthing")
	// }

	//Poiters

	var x int = 42
	var y *int //poiter to an integer
	y = &x

	fmt.Println(x, y, *y) // now y is the adress of x and *y is the value of x
	x = 15
	fmt.Println(x, y, *y) // now when i change y i change x  too

	//when you initialize a pointer like var i *int
	// the pointer will get the value <nil>

	var ms *Person //it will point to nil
	ms = new(Person)
	ms.lastname = "aziz"

	//Functions
	yoyo(12)
	fmt.Println(test(1, 2, 3, 4, 6))
	if v, err := devide(3, 2); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(v)
	}
	p2 := Person{
		name:     "aziz",
		lastname: "chagour",
		age:      20,
	}
	p2.getName()

}

func yoyo(msg int) {
	fmt.Println(msg, "from yoyo function")
}

func test(args ...int) int { //same as *args in python it gets a list or slice
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

// returning a response with an error
func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot devide by 0")
	}
	return a / b, nil
}

//methods

func (p *Person) getName() {
	fmt.Println(p.name)
}
