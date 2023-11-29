package mq

import (
	"VideoWeb2/conf"
	"github.com/streadway/amqp"
)

// CreatePublisher 设置订阅者
func CreatePublisher(receiverName string, content string) error {
	ch, err := Connect.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	//声明交换机
	err = ch.ExchangeDeclare(conf.ExchangeName, amqp.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		return err
	}

	//声明队列
	_, err = ch.QueueDeclare(receiverName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	//队列和交换机绑定
	key := receiverName
	err = ch.QueueBind(receiverName, key, conf.ExchangeName, false, nil)
	if err != nil {
		return err
	}
	//将消息发送到队列
	err = ch.Publish(
		conf.ExchangeName, // exchange 交换机名称
		key,               // routing key 路由名称
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{ //发送消息数据
			ContentType: "text/plain",
			Body:        []byte(content),
		})
	if err != nil {
		return err
	}
	return nil
}
