package bootstrap

import (
	"dataspace/api"
	"sync"
)

// loadSubprocesses is a function that loads and starts all the subprocesses such as the API
// websockets, etc.
func LoadSubprocesses(wg *sync.WaitGroup) {
	wg.Add(1)
	go api.Start()
}
