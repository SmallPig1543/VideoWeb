package mq

import (
	"VideoWeb2/conf"
	"github.com/streadway/amqp"
)

var Connect *amqp.Connection

func LinkRabbitmq() {
	url := "amqp://" + conf.RabbitmqUser + ":" + conf.RabbitmqPassword + "@" + conf.RabbitmqIP + "/"
	connect, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}
	Connect = connect
}
