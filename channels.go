//Channels and let the fun begin

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) //strongly typed

	wg.Add(2)

	go func() {
		//receiving go routine
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		//Sending go routine
		ch <- 42 // u r passing a copy of the data not the actual data
		wg.Done()
	}()
	wg.Wait()

	//receive only channels
	ch1 := make(chan int)

	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		} // recieve means we gonna get the value of the channel
		wg.Done()

	}(ch1)

	//send only channel
	go func(ch chan<- int) {
		ch <- 42 // send means the channel will get a value
		ch <- 2
		ch <- 4
		close(ch)
		wg.Done()

	}(ch1)

	wg.Wait()
}
