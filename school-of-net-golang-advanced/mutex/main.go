package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// vai ser executado qnd aplicação for iniciada
func init() {
	runtime.GOMAXPROCS(1)
}

var result int
var m sync.Mutex

var waitGroup sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	waitGroup.Add(2)
	go runProcess("P1", 20)
	go runProcess("P2", 20)

	waitGroup.Wait()
	fmt.Println("Final Result ", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()
		result++
		fmt.Println(name, "->", i, "Partial Result", result)
		m.Unlock()
	}
	waitGroup.Done()
}
