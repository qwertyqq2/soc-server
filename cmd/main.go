package main

import (
	"log"
	"sync"

	"github.com/qwertyqq2/soc-server/internal/app/apiserver"
	"github.com/qwertyqq2/soc-server/internal/configs"
)

var wg sync.WaitGroup

func main() {
	// wg.Add(1)
	// provider := &provider.Provider{}
	// provider.SetUp()
	// go provider.ListenEvent()
	// wg.Wait()

	conf := configs.GetConfig()
	if err := apiserver.Start(conf); err != nil {
		log.Fatal(err)
	}
}
