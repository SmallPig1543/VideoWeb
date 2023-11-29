package mq

// Receive 消费者接受消息
func Receive(receiverName string, messages chan string, channel chan bool) {
	ch, _ := Connect.Channel()
	defer ch.Close()
	_, _ = ch.QueueDeclare(receiverName, true, false, false, false, nil)
	msgs, _ := ch.Consume( // 注册一个消费者（接收消息）
		receiverName, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	go func() {
		for d := range msgs {
			messages <- string(d.Body)
		}
		close(messages)
	}()
	channel <- true
}
