package main

import (
	"math"
	"runtime"
	"sync"
)

func heat() {
	var t float64 = math.MaxFloat64
	for {
		t /= 2.0
		if t < math.SmallestNonzeroFloat32 {
			t = math.MaxFloat64
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < runtime.NumCPU(); i++ {
		go heat()
	}
	wg.Wait()
}
