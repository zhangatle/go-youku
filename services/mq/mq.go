package mq

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
)

type Callback func(msg string)

func Connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	return conn, err
}

// 发送端函数
func Publish(exchange string, queueName string, body string) error {
	// 建立连接
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	// 创建通道channel
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	// 创建队列
	queue, err := channel.QueueDeclare(queueName, true, false, false,false, nil)
	if err != nil {
		return err
	}
	// 发送消息
	err = channel.Publish(exchange, queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType: "text/plain",
		Body: []byte(body),
	})
	return err
}

// 接收者方法
func Consumer(exchange string, queueName string, callback Callback) {
	// 建立连接
	conn, err := Connect()
	if err != nil {
		logs.Error(err)
		return
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		logs.Error(err)
		return
	}
	defer channel.Close()
	queue, err := channel.QueueDeclare(queueName, true, false,false, false,nil)
	if err != nil {
		logs.Error(err)
		return
	}
	msg, err := channel.Consume(queue.Name, "", false, false, false, false,nil)
	if err != nil {
		logs.Error(err)
		return
	}
	forever := make(chan int)
	go func() {
		for d := range msg{
			s := BytesToString(&(d.Body))
			callback(*s)
			_ = d.Ack(false)
		}
	}()
	fmt.Printf("Waiting for message")
	<-forever
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

func PublishEx(exchange string, types string, routingKey string, body string) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	// 创建交换机
	err = channel.ExchangeDeclare(exchange, types, true, false, false, false, nil)
	if err != nil {
		return err
	}
	err = channel.Publish(exchange, routingKey, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType: "text/plain",
		Body: []byte(body),
	})
	return err
}

func ConsumerEx(exchange string, types string, routingKey string, callback Callback) {
	conn, err := Connect()
	if err != nil {
		return
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		return
	}
	defer channel.Close()
	// 创建交换机
	err = channel.ExchangeDeclare(exchange, types, true, false, false, false, nil)
	if err != nil {
		return
	}
	// 创建队列
	queue, err := channel.QueueDeclare("", false, false, true,false, nil)
	if err != nil {
		return
	}
	// 绑定
	err = channel.QueueBind(queue.Name, routingKey, exchange, false, nil)
	if err != nil {
		return
	}
	msg, err := channel.Consume(queue.Name, "", false, false, false,false, nil)
	if err != nil {
		return
	}
	forever := make(chan int)
	go func() {
		for d := range msg{
			s := BytesToString(&(d.Body))
			callback(*s)
			_ = d.Ack(false)
		}
	}()
	fmt.Printf("Waiting for message\n")
	<-forever
}

func PublishDlx(exchangeA string, body string) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	// 消息发送到A交换机
	err = channel.Publish(exchangeA, "", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType: "text/plain",
		Body: []byte(body),
	})
	return err
}

func ConsumerDlx(exchangeA string, queueAName string, exchangeB string, queueBName string, ttl int, callback Callback)  {
	conn, err := Connect()
	if err != nil {
		return
	}
	defer conn.Close()
	channel, err := conn.Channel()
	if err != nil {
		return
	}
	// 创建A交换机，创建A队列，A交换机和A队列绑定
	err = channel.ExchangeDeclare(exchangeA, "fanout", true, false, false, false,nil)
	if err != nil {
		return
	}
	queueA, err := channel.QueueDeclare(queueAName, true, false, false, false, amqp.Table{
		// 当消息过期时把消息发送到exchangeB
		"x-dead-letter-exchange" : exchangeB,
		"x-message-ttl":ttl,
	})
	if err != nil {
		return
	}
	err = channel.QueueBind(queueA.Name, "", exchangeA, false, nil)
	if err != nil {
		return
	}
	// 创建B交换机，创建B队列，B交换机和B队列绑定
	err = channel.ExchangeDeclare(exchangeB, "fanout", true, false, false, false,nil)
	if err != nil {
		return
	}
	queueB, err := channel.QueueDeclare(queueBName,true, false, false, false, nil)
	if err != nil {
		return
	}
	err = channel.QueueBind(queueB.Name, "", exchangeB, false, nil)
	if err != nil {
		return
	}
	msg, err := channel.Consume(queueB.Name, "", false, false, false,false, nil)
	if err != nil {
		return
	}
	forever := make(chan int)
	go func() {
		for d := range msg{
			s := BytesToString(&(d.Body))
			callback(*s)
			_ = d.Ack(false)
		}
	}()
	fmt.Printf("Waiting for message\n")
	<-forever
}