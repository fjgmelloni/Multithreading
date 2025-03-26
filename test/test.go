package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp, err := http.Get("http://localhost:8080/cep/14170420")
			if err != nil {
				fmt.Printf("Erro [%d]: %v\n", i, err)
				return
			}
			fmt.Printf("Resposta [%d]: %s\n", i, resp.Status)
			resp.Body.Close()
		}(i)
	}
	wg.Wait()
}
