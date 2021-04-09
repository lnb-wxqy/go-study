package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Shopify/sarama"
	"go-study/utils"
	"sync"
	"time"
)

type KafkaProducer struct {
	sync.Mutex

	BootStrap            []string
	Retry                int
	SaramakafkaProducer  sarama.SyncProducer
	KafkaProducerVersion string
	Topic                string
	ctx                  context.Context
}

func (k *KafkaProducer) InitKafkaProduceConnection() {
	k.Lock()
	defer k.Unlock()
	if len(k.BootStrap) == 0 || k.Topic == "" {
		fmt.Println("bootstrap is nil or empty or topic is empty!")
		return
	}

	if k.Retry == 0 {
		k.Retry = 3
	}
	if k.SaramakafkaProducer != nil {
		_ = k.SaramakafkaProducer.Close()
	}
	config := sarama.NewConfig()
	config.Version = sarama.V2_3_0_0
	config.Producer.MaxMessageBytes = 5 * 1024 * 1024
	var err error
	if k.KafkaProducerVersion != "" {
		config.Version, err = sarama.ParseKafkaVersion(k.KafkaProducerVersion)
		if err != nil {
			//logger.LOG_ERROR("ParseKafkaVersion is error", err)
			fmt.Println("ParseKafkaVersion is error: ", err)
		}
	}

	config.Producer.Return.Successes = true
	syncProducer, err := sarama.NewSyncProducer(k.BootStrap, config)
	k.SaramakafkaProducer = syncProducer
	if err != nil {
		fmt.Println("创建kafka-producer 失败： ", err)
		k.SaramakafkaProducer = nil
	}
	k.ctx, _ = context.WithCancel(context.Background())
}

func (k *KafkaProducer) Producer(data interface{}) error {
	dataByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json.marshal data error: ", err)
		return err
	}
	kafkaMsg := &sarama.ProducerMessage{
		Value: sarama.ByteEncoder(dataByte),
		Topic: k.Topic,
	}
	err = utils.RetryCancelWithContext(func() error {
		k.Lock()
		kafkaProducer := k.SaramakafkaProducer
		k.Unlock()
		if kafkaProducer == nil {
			// 发生异常，重连
			k.InitKafkaProduceConnection()
			return errors.New("kafka producer not connect")
		}
		partition, offset, err := kafkaProducer.SendMessage(kafkaMsg)
		//err := kafkaProducer.SendMessages(kafkaMsgs) 批量发送
		if err != nil {
			fmt.Println("kafka sendMessage error: ", err, "dst address： ", k.BootStrap, "message size is: ", kafkaMsg.Value.Length())
			//pe, ok := err.(sarama.ProducerErrors)
			er, ok := err.(sarama.ProducerError)
			if ok {
				fmt.Println("kafak sendMessage error detail: ", er.Error())
			}
			// 发送异常、重连
			k.Lock()
			k.SaramakafkaProducer = nil
			k.Unlock()
		}

		fmt.Println("send kafka mesage success! partition: ", partition, "offset: ", offset)
		return err

	}, k.Retry, 1*time.Second, k.ctx)

	return err
}
