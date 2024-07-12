package app

import (
	"net/http"

	"github.com/ansxy/golang-boilerplate-gin/config"
	http_transport "github.com/ansxy/golang-boilerplate-gin/internal/transport/http"
	custome_http "github.com/ansxy/golang-boilerplate-gin/pkg/http"
	custome_kafka "github.com/ansxy/golang-boilerplate-gin/pkg/kafka"
)

func Exec() (err error) {
	handler := http_transport.NewHttpTransport()

	// w := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{"localhost:9092"},
	// 	Topic:   "api",
	// })

	// msg := kafka.Message{
	// 	Key:   []byte("key"),
	// 	Value: []byte("value"),
	// }

	// if err = w.WriteMessages(context.Background(), msg); err != nil {
	// 	log.Fatal(err)
	// }

	// if err = w.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// r := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers:     []string{"localhost:9092"},
	// 	Topic:       "api",
	// 	Partition:   0,
	// 	MinBytes:    10e3,
	// 	MaxBytes:    10e6,
	// 	Logger:      kafka.LoggerFunc(log.Printf),
	// 	ErrorLogger: kafka.LoggerFunc(log.Printf),
	// })

	// defer r.Close()

	// fmt.Println("start reading")

	// for {
	// 	msg, err := r.ReadMessage(context.Background())
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(string(msg.Value))
	// }

	custome_kafka.InitKafka(&config.KafkaConfig{
		KafkaConfigReader: config.KafkaConfigReader{
			Brokers:   []string{"localhost:9092"},
			Topic:     "api",
			Partition: 0,
		},
		KafkaConfigWriter: config.KafkaConfigWriter{
			Brokers: []string{"localhost:9092"},
			Topic:   "api",
		},
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err = custome_http.NewHttpServer(srv)
	if err != nil {
		return err
	}

	return
}
