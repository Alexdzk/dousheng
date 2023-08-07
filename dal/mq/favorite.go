package mq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Alexdzk/dousheng/dal/db"
	"github.com/streadway/amqp"
)

var mq *amqp.Connection

func Init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	mq = conn
}

func PublishFavoriteMsg(ctx context.Context, favoriteMsg *db.FavoriteRaw) error {
	ch, err := mq.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"favorite",
		false, // 是否持久化
		false, // 是否自动删除
		false, // 是否独占
		false, // 是否等待服务器响应
		nil,   // 额外参数
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return err
	}
	body, err := json.Marshal(favoriteMsg)
	if err != nil {
		log.Fatalf("Failed to marshal Person object: %v\n", err)
	}
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
		return err
	}
	log.Printf("Sent a message: %s", body)
	return nil
}

func ConsumeFavoriteMsg(ctx context.Context) {
	ch, err := mq.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"favorite", // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var fab db.FavoriteRaw
			if err := json.Unmarshal(d.Body, &fab); err != nil {
				log.Printf("Failed to decode message: %v\n", err)
				d.Ack(false)
				continue
			}
			db.CreateFavorite(ctx, &fab, fab.VideoId)
			d.Ack(false)
		}
	}()
	<-forever
}
