package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"

	"sync"
)

type testData struct {
	Name string
	Data string
}

func consumer(ch chan testData) {
	i := 0
	for {
		dt := <-ch
		i++
		if i%100000 == 0 {
			fmt.Printf("Data %d %s\n", i, dt)
		}
	}
}

func test(wg sync.WaitGroup) {
	defer wg.Done()
	c := make(chan testData)

	go consumer(c)
	for i := 0; i < 10000000; i++ {
		x := fmt.Sprintf("%d", i)
		c <- testData{Name: x}
	}

}

func main() {
	go func() {

		r := http.NewServeMux()

		// Register pprof handlers
		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)

		log.Println(http.ListenAndServe("localhost:6060", r))
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	test(wg)
	wg.Wait()
}
