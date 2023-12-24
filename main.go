package main

import (
	"dataspace/api"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello, World!")

	// Create a wait group
	var wg sync.WaitGroup

	wg.Add(1)

	go api.Start()

	wg.Wait()
}
