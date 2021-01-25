package nsq_lib

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"golang-study/utils"
	"time"
)

var tcpNsqAddr = "39.106.0.106:4150"

// topic
var topic = "topic-test"

// 生产数据
func NsqProducer() {
	config := nsq.NewConfig()

	for i := 0; i < 20; i++ {
		producer, e := nsq.NewProducer(tcpNsqAddr, config)
		if e != nil {
			fmt.Println(e)
		}

		// 内容
		tContext := "this is test data"

		// 发布消息
		e = producer.Publish(topic, utils.StringBytes(tContext))
		if e != nil {
			fmt.Println(e)
		}

	}
}

// --------------------------------消费数据--------------------------

// 声明一个结构体，实现HandleMessage接口方法
type NsqHandler struct {
	// 消息数
	nsqCount int64
	// 标识ID
	nsqHandlerID string
}

// 实现HandleMessage方法
func (n *NsqHandler) HandleMessage(message *nsq.Message) error {
	// 没收到一条消息 +1
	n.nsqCount++
	fmt.Println(n.nsqCount, n.nsqHandlerID)
	fmt.Printf("msg.timestamp=%v, msg.nsqaddress=%s,msg.body=%s \n", time.Unix(0, message.Timestamp).Format("2006-01-02 03:04:05"), message.NSQDAddress, utils.BytesString(message.Body))

	return nil
}

// 消费数据
func NsqConsumer() {
	config := nsq.NewConfig()
	// 创造消费者，参数一是订阅主题topic，参数二是使用的通道
	consumer, e := nsq.NewConsumer(topic, "channel", config)
	if e != nil {
		fmt.Println(e)
	}
	// 添加处理回调
	consumer.AddHandler(&NsqHandler{nsqHandlerID: "one"})
	// 连接对应的nsqd
	e = consumer.ConnectToNSQD(tcpNsqAddr)
	if e != nil {
		fmt.Println(e)
	}
}
