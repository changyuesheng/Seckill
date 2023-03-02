package repository

import (
	"errors"
	"fmt"
	"seckill/data"
	"seckill/model"

	"github.com/streadway/amqp"
)

type rabbitMQRepository struct {
	mq *data.Rabbitmq
}

func NewRabbitRepository(rabbitmq *data.Rabbitmq) model.RabbitMQRepository {
	rabbit := &rabbitMQRepository{
		mq: rabbitmq,
	}
	//go rabbit.GetData()
	return rabbit
}

func (r *rabbitMQRepository) Push(data string) error {
	err := r.mq.Producer.Channel.Publish(
		"",                  // exchange
		r.mq.Producer.Queue, // routing key
		false,               // mandatory
		false,               // immediate

		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		},
	)
	if err != nil {
		fmt.Println("push data to queue failed")
		return errors.New("push data to queue failed")
	}
	fmt.Println("push data to queue successfully")
	return nil
}

func (r *rabbitMQRepository) Getch() *amqp.Channel {
	return r.mq.Consumer.Channel

}

func (r *rabbitMQRepository) Getqname() string {
	return r.mq.Consumer.Queue

}

// func (r *rabbitMQRepository) GetData() {
// 	msgs, err := r.mq.Consumer.Channel.Consume(
// 		r.mq.Producer.Queue, // queue
// 		"toutiao",           // consumer
// 		true,                // auto-ack,true消费了就消失
// 		false,               // exclusive
// 		false,               // no-local
// 		false,               // no-wait
// 		nil,                 // args
// 	)
// 	if err != nil {
// 		fmt.Println("Failed to register a consumer")
// 	}
// 	for d := range msgs {
// 		temp := strings.Split(string(d.Body), ".")
// 		oder := &model.Order{
// 			Username:  temp[0],
// 			GoodsName: temp[1],
// 		}
// 		err := m.DB.Create(&o).Error

// 	}
// }
