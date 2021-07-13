package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

//全局的kafka异步生产者
var KafkaAsyncProducer sarama.AsyncProducer
var KafkaProducer sarama.SyncProducer
var KafkaConsumer sarama.Consumer

func KafkaAsyncProducerInit() {
	var e error
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//注意，版本设置不对的话，kafka会返回很奇怪的错误，并且无法成功发送消息
	//config.Version = sarama.V0_10_0_1
	config.Version = sarama.V2_0_0_0

	fmt.Println("start make producer")
	//使用配置,新建一个异步生产者
	//KafkaAsyncProducer, e = sarama.NewAsyncProducer(strings.Split("10.94.0.1:9092, 10.94.0.2:9092，10.94.0.3:9092", ","), config)
	KafkaAsyncProducer, e = sarama.NewAsyncProducer(strings.Split(Cfg.Kafka.Urls, ","), config)

	if e != nil {
		fmt.Println(e)
		return
	}
	//defer producer.AsyncClose()

}

func KafkaAsyncProduce(topic string, content string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(content),
	}
	//使用通道发送
	KafkaAsyncProducer.Input() <- msg
	//msg = <- KafkaAsyncProducer.Successes()
}

//同步生产
func KafkaProducerInit() (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          //赋值为-1：这意味着producer在follower副本确认接收到数据后才算一次发送完成。
	config.Producer.Partitioner = sarama.NewRandomPartitioner //写到随机分区中，默认设置8个分区
	config.Producer.Return.Successes = true
	for i := 1; i <= 10; i++ {
		KafkaProducer, err = sarama.NewSyncProducer(strings.Split(Cfg.Kafka.Urls, ","), config)
		//defer client.Close()
		if err != nil {
			log.Println("kafka连接失败,尝试重连……")
		} else {
			break
		}
		if i == 10 {
			log.Println("kafka连接失败！")
			return err
		}
		time.Sleep(6 * time.Second)
	}
	return err
}

func KafkaProduce(topic string, content string) error {
	//fmt.Println("topic:")
	//fmt.Println(topic)
	//fmt.Println("content:")
	//fmt.Println(content)
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(content)

	//fmt.Println(msg)
	pid, offset, err := KafkaProducer.SendMessage(msg)
	c := 0
	for {
		c++
		if err == nil {
			break
		}
		if c > 10 {
			return err
		}
		pid, offset, err = KafkaProducer.SendMessage(msg)
	}
	fmt.Printf("分区ID:%v, offset:%v \n", pid, offset)
	return err
}

func KafkaConsumerInit() (err error) {

	fmt.Println("start consumer")
	config := sarama.NewConfig()
	//提交offset的间隔时间，每秒提交一次给kafka
	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	//config.Version = sarama.V0_10_0_1
	config.Version = sarama.V2_0_0_0
	//consumer新建的时候会新建一个client，这个client归属于这个consumer，并且这个client不能用作其他的consumer
	//KafkaConsumer, err = sarama.NewConsumer(strings.Split("10.94.0.1:9092, 10.94.0.2:9092，10.94.0.3:9092", ","), config)
	if Cfg.Kafka.Urls == "" {
		fmt.Println("---No config found! kafka onsumer stop!")
	}
	KafkaConsumer, err = sarama.NewConsumer(strings.Split(Cfg.Kafka.Urls, ","), config)
	if err != nil {
		return
	}
	fmt.Println("consumer build!")
	var wg sync.WaitGroup
	topic := strconv.FormatInt(1, 10)
	//topic :="6867305865050018969"
	partitionList, err := KafkaConsumer.Partitions(topic) //获得本网关topic的所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition:, ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := KafkaConsumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("Failed to start consumer for partition : ", err)
			return err
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { //为每个分区开一个go协程去取值
			for msg := range pc.Messages() { //阻塞直到有值发送过来，然后再继续等待
				fmt.Printf("kafka接收到消费信息：Partition:%d\n, Offset:%d\n, key:%s\n, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				if string(msg.Value) != "" { //此处为具体的方法逻辑或者对应的方法调用

					var msgMap map[string]interface{}
					if err := json.Unmarshal([]byte(string(msg.Value)), &msgMap); err == nil {
						fmt.Println(msgMap)
						//fmt.Println(dat["status"])
					} else {
						fmt.Println(err)
					}
					if msgMap["kafkaFuncNum"] == nil || msgMap["kafkaFuncNum"].(string) == "" {
						return
					}

					if msgMap["kafkaFuncNum"].(string) == "reboot" {
						rebootCmd := exec.Command("bash", "-c", "reboot")
						rebootCmd.Run()
						return
					}
					funcNum, err := strconv.Atoi(msgMap["kafkaFuncNum"].(string))
					if err != nil {
						fmt.Println(err)
					}
					url := "http://127.0.0.1:9000" + FuncDict[funcNum]
					contentType := "application/json;charset=utf-8"
					log.Println("********************************kafka消费******************************")
					log.Println(url)
					log.Println("********************************kafka消费******************************")
					body := bytes.NewBuffer([]byte(string(msg.Value)))
					http.Post(url, contentType, body)

				}
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
	return nil

	//KafkaConsumer.Close()
	return err
}
