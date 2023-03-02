package data

import (
	"fmt"
	"seckill/config"

	"github.com/streadway/amqp"
)

type Rabbitmq struct {
	Consumer *Amqp
	Producer *Amqp
}

type Amqp struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      string
}

func initRabbitMQ(config config.AppConfig) *Rabbitmq {
	return &Rabbitmq{
		Consumer: CreateConn(config),
		Producer: CreateConn(config),
	}
}

func CreateConn(config config.AppConfig) *Amqp {
	fmt.Println("---Start initializing RabbitMQ connection---")
	var err error
	rabbitq_conn, err := amqp.Dial(
		fmt.Sprintf("amqp://%s:%s@%s:%s/",
			config.RabbitMQUser,
			config.RabbitMQPaasword,
			config.RabbitMQHost,
			config.RabbitMQPort,
		),
	)
	if err != nil {
		fmt.Println("init rabbitmq connection failed, err:", err)
		panic(err)
	}
	ch, err := rabbitq_conn.Channel()
	if err != nil {
		fmt.Println("init rabbitmq channel failed, err:", err)
		panic(err)
	}
	_, err = ch.QueueDeclare(
		// 队列名称
		config.RabbitMQConfig.QueueName, // name
		false,                           // durable
		false,                           // delete when unused
		false,                           // exclusive
		false,                           // no-wait
		nil,                             // arguments
	)
	return &Amqp{
		Connection: rabbitq_conn,
		Channel:    ch,
		Queue:      config.RabbitMQConfig.QueueName,
	}
}
