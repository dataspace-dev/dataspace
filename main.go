package main

import (
	"dataspace/api"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	// This starts the API REST where the routes are defined and the server is run
	// The frontend will call the routes in it
	go api.Start()

	wg.Wait()
}
