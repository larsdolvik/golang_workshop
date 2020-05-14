package examples

import (
	"fmt"
	"net/http"
	"sync"
)

// Go routines can be thought of as lightweight threads,
// and is used to execute code concurrently.

// Multiple go routines can run on the same thread. Each
// only take up 2kb of memory.

// https://gobyexample.com/goroutines

func routine() {
	fmt.Print("Hello world")
}

// RoutineExample runs a single go routine without waiting
func RoutineExample() {
	go routine()
	fmt.Print("Hi there.")
	// time.Sleep(time.Second)
}

// WaitForAllRoutines waits for all go routines to finish before returning
func WaitForAllRoutines() {
	urls := []string{
		"https://www.google.com",
		"https://www.vg.no",
		"https://kode24.no",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, val := range urls {
		go getURL(val, &wg)
	}
	wg.Wait()
}

func getURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error retrieving url, %s", err)
	}

	fmt.Printf("%d\n", resp.StatusCode)
}
