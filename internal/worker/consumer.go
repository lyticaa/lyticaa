package worker

import (
	"os"

	"github.com/streadway/amqp"
)

func (w *Worker) Start() {
	w.Logger.Info().Msgf("listening for messages on %v....", w.queue())

	conn, err := amqp.Dial(os.Getenv("CLOUDAMQP_URL"))
	if err != nil {
		w.Logger.Fatal().Err(err).Msg("failed to connect to the queue")
	}

	ch, err := conn.Channel()
	if err != nil {
		w.Logger.Fatal().Err(err).Msg("failed to connect to open a channel")
	}

	q, err := w.declare(ch)
	if err != nil {
		w.Logger.Fatal().Err(err).Msg("failed to declare queue")
	}

	msgs, err := w.consume(ch, q)
	if err != nil {
		w.Logger.Fatal().Err(err).Msg("failed to register a consumer")
	}

	go w.run(msgs)
	go func() {
		<-w.Signalling.quit
		w.Logger.Info().Msgf("stop listening on %v....", w.queue())
		w.Signalling.cancel()
		_ = conn.Close()
		_ = ch.Close()
	}()
	<-w.Signalling.stop
}

func (w *Worker) queue() string {
	return os.Getenv("CLOUDAMQP_QUEUE_REPORTS_PUBLISHED")
}

func (w *Worker) run(msgs <-chan amqp.Delivery) {
	for {
		select {
		case <-w.Signalling.ctx.Done():
			w.Signalling.stop <- struct{}{}
			return
		default:
			for d := range msgs {
				w.parse(d)
			}
		}
	}
}

func (w *Worker) declare(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(w.queue(), true, true, false, false, nil)
}

func (w *Worker) consume(ch *amqp.Channel, q amqp.Queue) (<-chan amqp.Delivery, error) {
	return ch.Consume(q.Name, "", true, false, false, false, nil)
}
