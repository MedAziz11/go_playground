package main

import (
	"fmt"
)

//Interface
type Writer interface {
	Write([]byte) (int, error) //methode declaratin
}

type Number interface{}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	i, err := fmt.Println(string(data))
	return i, err
}

func main() {
	var w Writer = ConsoleWriter{} // when u do like this your struct must have a Writer function
	w.Write([]byte("Hello World!"))
	var n Number = 0
	if v, ok := n.(Number); ok {
		fmt.Println(v)
	} else {
		fmt.Println("not int")
	}
}
