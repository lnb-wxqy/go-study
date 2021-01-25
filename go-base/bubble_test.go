package go_base

import (
	"fmt"
	"testing"
)

func TestBuffleSort(t *testing.T) {
	arr :=[]int{3,2,5,6,4,12,8,7,9}
	fmt.Println("排序前： ",arr)
	BuffleSort(arr)
	fmt.Println("排序后： ",arr)
}
