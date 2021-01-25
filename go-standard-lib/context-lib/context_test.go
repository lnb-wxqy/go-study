package context_lib

import "testing"

func TestWaitGroup(t *testing.T) {
	WaitGroup()
	t.Log("测试完毕")
}

func TestChan(t *testing.T) {
	Chan()
}

func TestContext(t *testing.T) {
	Context()
}

func TestMultiGoroutine(t *testing.T) {
	MultiGoroutine()
}

func TestContextWithCancel(t *testing.T) {
	ContextWithCancel()
}