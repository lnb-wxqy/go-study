package utils

import (
	"context"
	"time"
)



/**
重试
*/
func Retry(f func() error, times int, space time.Duration) error {
	if times == 0 || times == 1 {
		return f()
	}
	var err error
	var i = 0
	for {
		err = f()
		if err == nil {
			return nil
		}
		i++
		if times > 1 && i >= times {
			break
		}
		time.Sleep(space)
	}
	return err
}

/**
无限重试
*/
func RetryCancelWithContext(f func() error, times int,space time.Duration,ctx context.Context) error {
	if times == 0 || times == 1 {
		return f()
	}
	var err error
	var i = 0
	for {
		select {
		case <-ctx.Done():
			return err
		default:
		}
		err = f()
		if err == nil {
			return nil
		}
		i++
		if times > 1 && i >= times {
			break
		}
		time.Sleep(space)
	}
	return  err
}
