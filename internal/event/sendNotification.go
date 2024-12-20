package event

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/KKhimmoon/yuemnoi-reserve/internal/dto"

	"github.com/KKhimmoon/yuemnoi-reserve/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

// SendNotification is responsible for sending a notification message to RabbitMQ
func SendNotification(notification dto.NotificationRequest) {
	cfg := config.Load()

	// Marshal the notification struct into JSON
	body, err := json.Marshal(notification)
	if err != nil {
		log.Fatalf("[client]: failed to marshal notification %+v", err)
	}

	log.Print(body)
	log.Print(reflect.TypeOf(body))

	conn, err := amqp.Dial(cfg.RabbitMQUrl)
	if err != nil {
		log.Fatalf("[client]: unable to connect RabbitMQ %+v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("[client]: failed to open a channel %+v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"notification", // name (use a common queue name for communication)
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Fatalf("[client]: failed to declare a queue %+v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("[client]: failed to publish a message %+v", err)
	}

	log.Print(" [x] Notification Sent from yuemnoi-reserve \n")
	log.Print(notification)
}
