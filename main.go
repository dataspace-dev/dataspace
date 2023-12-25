package main

import (
	"dataspace/bootstrap"
	"sync"
)

func main() {

	bootstrap.LoadEnv() // Load the .env file
	bootstrap.SetGitMode() // Set the Gin mode (release or debug)

	// Load the subprocesses (api, socket manager, etc.)
	var wg sync.WaitGroup
	bootstrap.LoadSubprocesses(&wg) 
	wg.Wait()
}

