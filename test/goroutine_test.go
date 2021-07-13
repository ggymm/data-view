package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestBatch(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("dog")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("fish")
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("cat")
		}
		wg.Done()
	}()

	wg.Wait()
}
