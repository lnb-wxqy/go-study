package aotmic_lib

import (
	"fmt"
	"sync/atomic"
)

var count int32

func atomic_addint()  {
	atomic.AddInt32(&count,12)
	fmt.Println(count)
}
