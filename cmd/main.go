package main

import (
	"sync"

	"github.com/qwertyqq2/soc-server/internal/listner/provider"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	provider := &provider.Provider{}
	provider.SetUp()
	go provider.ListenEvent()
	wg.Wait()

	// provider := &provider.Provider{}
	// provider.SetUp()

	// conf := configs.GetConfig()
	// if err := apiserver.Start(conf, provider); err != nil {
	// 	log.Fatal(err)
	// }
}
