package context_lib

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//控制并发有两种经典方式，一种WaitGroup，一种是Context

//1.WaitGroup 控制多个goroutine同时完成，同步方式
//适用于，多个goroutine协同做一件事
//实际业务中，会有这么这一中场景：需要我们主动通知某一个goroutine结束，因此引入chan
func WaitGroup() {
	var wg sync.WaitGroup

	wg.Add(2) //添加任务总数
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		wg.Done() //任务执行完毕，调用done方式，任务书减一
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()

	wg.Wait() //等待所有goroutine执行完毕
	fmt.Println("任务都已执行完毕，放行")
}

//1.1 chan通知
//chan+select 结束goroutine
//局限性：如果有很多goroutine需要控制，而且这些goroutine又衍生了更过的goroutine。。。
func Chan() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了")
				return
			default:
				fmt.Println("goroutine 监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("可以通知监控停止了")
	stop <- true
	time.Sleep(5 * time.Second)
}

//2.context
//上面说的这种场景是存在的，比如一个网络请求Request，每个Request都需要开启一个goroutine做一些事情，
// 这些goroutine又可能会开启其他的goroutine。所以我们需要一种可以跟踪goroutine的方案，
// 才可以达到控制他们的目的，这就是Go语言为我们提供的Context，称之为上下文非常贴切，它就是goroutine的上下文。
//2.1 用context重写Chan方法
func Context() {
	ctx, cancel := context.WithCancel(context.Background())
	//ctx当做参数传入，可以跟踪这goroutine
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出了...", ctx.Err())
				return
			default:
				fmt.Println("监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("通知goroutine停止...")
	cancel()
	time.Sleep(5 * time.Second)
}

//2.2 Context控制多个goroutine
func MultiGoroutine() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "【goroutine1】")
	go watch(ctx, "【goroutine2】")
	go watch(ctx, "【goroutine3】")

	time.Sleep(2 * time.Second)
	fmt.Println("可以结束所有goroutine了")
	cancel()
	time.Sleep(4 * time.Second)

}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "运行结束...", ctx.Err())
			return
		default:
			fmt.Println(name, "运行中...")
			time.Sleep(3 * time.Second)
		}
	}
}

//2.3 WithCancel
var key = "name"

func ContextWithCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watchV(valueCtx)
	time.Sleep(2 * time.Second)
	fmt.Println("可以结束了...")
	cancel()
	time.Sleep(4 * time.Second)
}

func watchV(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), "监控退出了，停止了")
			return
		default:
			fmt.Println(ctx.Value(key), "goroutine运行中...")
			time.Sleep(2 * time.Second)
		}
	}
}
