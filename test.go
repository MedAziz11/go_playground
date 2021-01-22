package main

import "fmt"

func sendValues(ch chan string) {
	ch <- "hello"
}

func main() {
	values := make(chan string)
	defer close(values)

	go sendValues(values)
	value := <-values
	fmt.Println(value)

}
