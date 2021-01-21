package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // read write Mutex it means alot of proc can read but only one can write

func main() {
	//runtime.GOMAXPROCS(100)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock() //write lock
		go count()
	}
	//bad example because we locking and unlocking

	wg.Wait()

}

func sayHello() {

	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func count() {

	counter++ // write a new value to the counter
	m.Unlock()
	wg.Done()
}
