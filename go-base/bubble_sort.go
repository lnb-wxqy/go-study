package go_base

// 冒泡排序
func BuffleSort(arr []int) {

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 如果前者比后者大，则执行交换
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
