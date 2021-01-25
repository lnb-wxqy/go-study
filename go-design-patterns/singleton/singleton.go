package singleton

import "sync"

// 单例模式
type singleton struct {
}

var instance *singleton

var once sync.Once

// sync.Once的do方法在程序运行过程中只执行一次
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
