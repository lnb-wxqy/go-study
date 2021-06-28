## Go 语言常用包学习

#### 一、sync包

##### （一）waitgroup

- WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程的结束。

  例如：

  ```go
  func main() {
  
  	var wg sync.WaitGroup
  
  	for i := 0; i <= 10; i++ {
  		// increment the WaitGroup counter
  		wg.Add(1)
  
  		go func(i int) {
  			// decrement the counter when the goroutine completes
  			defer wg.Done()
  			fmt.Println(i)
  		}(i)
  	}
  
  	// wait for all go func to complete
  	wg.Wait()
  }
  ```

  1. `func (wg *WaitGroup) Add(delta int)` 
     - Add方法向内部计数加上delta，delta可以是负数；如果内部计数器变为0，Wait方法阻塞等待的所有线程都会释放，如果计数器小于0，方法panic。注意Add加上正数的调用应在Wait之前，否则Wait可能只会等待很少的线程。一般来说本方法应在创建新的线程或者其他应等待的事件之前调用。
  2. `func (wg *WaitGroup) Done()`
     - Done方法减少WaitGroup计数器的值，应在线程的最后执行。
  3. `func (wg *WaitGroup) Wait()` 
     - Wait方法阻塞直到WaitGroup计数器减为0。

##### （二）Once

- Once 是只执行一次动作的对象

- 例如：

  ```go
  func main() {
  
  	var once sync.Once
  
  	onceBody := func() {
  		fmt.Println("Only once")
  	}
  
  	done := make(chan bool)
  	for i := 0; i < 10; i++ {
  		go func() {
  			once.Do(onceBody)
  			done <- true
  		}()
  	}
  
  	for i := 0; i < 10; i++ {
  		<-done
  	}
  }
  
  // 运行结果
  Only once
  ```

1.  `func (o *Once) Do(f func())` 

   Do 方法当且仅当第一次调用是才会执行函数f。

   Do 用于必须刚好运行一次的初始化。因为f没有参数的，因为可能需要使用闭包来提供给Do方法调用：

   ```go
   config.once.Do(func(){ config.init(filename)})
   ```

   因为只有f返回后Do方法才会返回，f若引起了Do的调用，会导致死锁

   

##### 二、sync/atomic



#### 二、sync/atomic

#### 三、context

#### 四、strconv

#### 五、time

#### 六、strings

#### 七、bufio

#### 八、os

#### 九、json

#### 十 、flag

#### 十一、io

#### 十二、math

#### 十三、net

#### 十四、reflect

#### 十五、regexp

#### 十六、encoding

#### 十七、

