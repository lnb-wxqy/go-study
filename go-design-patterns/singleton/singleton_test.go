package singleton

import (
	"sync"
	"testing"
)

const count = 100

func TestGetInstance(t *testing.T) {

	ins1 := GetInstance()
	ins2 := GetInstance()

	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

//multi goroutine
func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(count)
	instances := [count]*singleton{}
	for i := 0; i < count; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < count-1; i++ {
		if instances[i] != instances[i+1] {
			t.Fatal("instance is not equal")
		}
	}
}
