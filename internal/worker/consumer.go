package app

import (
	"os"

	"github.com/streadway/amqp"
)

func (a *App) Start() {
	a.Logger.Info().Msgf("listening for messages on %v....", os.Getenv("CLOUDAMQP_QUEUE"))

	conn, err := amqp.Dial(os.Getenv("CLOUDAMQP_URL"))
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to connect to the queue")
	}

	ch, err := conn.Channel()
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to connect to open a channel")
	}

	q, err := a.declare(ch)
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to declare queue")
	}

	msgs, err := a.consume(ch, q)
	if err != nil {
		a.Logger.Fatal().Err(err).Msg("failed to register a consumer")
	}

	go a.run(msgs)
	go func() {
		<-a.Signalling.quit
		a.Logger.Info().Msgf("stop listening on %v....", os.Getenv("CLOUDAMQP_QUEUE"))
		a.Signalling.cancel()
		_ = conn.Close()
		_ = ch.Close()
	}()
	<-a.Signalling.stop
}

func (a *App) run(msgs <-chan amqp.Delivery) {
	for {
		select {
		case <-a.Signalling.ctx.Done():
			a.Signalling.stop <- struct{}{}
			return
		default:
			for d := range msgs {
				a.parse(d)
			}
		}
	}
}

func (a *App) declare(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		os.Getenv("CLOUDAMQP_QUEUE"),
		true,
		true,
		false,
		false,
		nil,
	)
}

func (a *App) consume(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}
