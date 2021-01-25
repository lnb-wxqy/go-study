package kafka

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type FromOffset int

const (
	FROM_OFFSET_NEW FromOffset = 0 // 实时消费
	FROM_OFFSET_OLD FromOffset = 1 // 从头消费
)

type KafkaConsumer struct {
	sync.Mutex

	Bootstrap []string
	Topic     string
	GroupId   string
	FromOffset
	IsFromEarliestOffset bool
	BatchSize            int // 批次大小，即每次消费kafka的数据量
	BatchDelay           int
	parallelSize         int
	topicSleep           map[string]time.Duration
	startTime            string
	endTime              string
	KafkaVersion         string

	cancelCtx context.Context
	cancel    func()
}

func (kc *KafkaConsumer) InitKafkaConsumerConnect() {
	kc.Lock()
	defer kc.Unlock()
	if len(kc.Bootstrap) == 0 || kc.Topic == "" || kc.GroupId == "" {
		fmt.Println("bootstrap or topic  or groupId is null 0r empty!")
		return
	}
	kc.FromOffset = FROM_OFFSET_OLD
	if kc.IsFromEarliestOffset {
		kc.FromOffset = FROM_OFFSET_NEW
	}
	if kc.BatchSize < 1 {
		kc.BatchSize = 1
	}


}
