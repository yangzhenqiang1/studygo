package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func speak(i int) {
	defer wg.Done()
	fmt.Printf("speak,index:%#v\n", i)
}

func main() {
	fmt.Printf("start，%v\n", time.Now())
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go speak(i)
	}

	// time.Sleep(3 * time.Second)
	wg.Wait()
	fmt.Printf("end，%v\n", time.Now().Format("2006-01-02 15:04:05"))
}
