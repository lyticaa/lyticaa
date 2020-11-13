package expenses

import (
	"encoding/json"
	"os"

	"github.com/streadway/amqp"
)

func NewPayload() *Payload {
	return &Payload{Message: Message{}}
}

type Message struct {
	Op     string `json:"op,omitempty"`
	UserID string `json:"userID"`
	Body   string `json:"body"`
}

type Payload struct {
	Message `json:"message"`
}

func Send(queueName string, message []byte) error {
	conn, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := channel(conn)
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := queue(ch, queueName)
	if err != nil {
		return err
	}
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        message,
	})
	if err != nil {
		return err
	}

	return nil
}

func connect() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("CLOUDAMQP_URL"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func channel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func queue(ch *amqp.Channel, name string) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(name, true, true, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &q, nil
}

func sendMessage(userID, op, data string) error {

	payload := NewPayload()
	payload.Op = op
	payload.Message.UserID = userID
	payload.Message.Body = data

	js, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err = Send(os.Getenv("CLOUDAMQP_QUEUE_EXPENSES_PENDING"), js); err != nil {
		return err
	}

	return nil
}
