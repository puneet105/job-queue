package queue

import (
    "github.com/streadway/amqp"
    "gopkg.in/yaml.v2"
    "os"
    "fmt"
)

type RabbitMQ struct {
    Conn    *amqp.Connection
    Channel *amqp.Channel
}

func NewRabbitMQ() (*RabbitMQ, error) {
    config := struct {
        RabbitMQ struct {
            URL string `yaml:"url"`
        } `yaml:"rabbitmq"`
    }{}

    data, err := os.ReadFile("config.yaml")
    if err != nil {
        return nil, err
    }

    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }
    fmt.Println(config.RabbitMQ.URL)
    conn, err := amqp.Dial(config.RabbitMQ.URL)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &RabbitMQ{Conn: conn, Channel: ch}, nil
}

func (r *RabbitMQ) Publish(queueName string, message string) error {
    err := r.Channel.Publish("", queueName, false, false, amqp.Publishing{
        ContentType: "text/plain",
        Body:        []byte(message),   
    })
    fmt.Printf("Published message To RabbitMQ: %s\n", message)
    return err
}

func (r *RabbitMQ) Consume(queueName string) (<-chan amqp.Delivery, error) {
    msgs, err := r.Channel.Consume(queueName, "", true, false, false, false, nil)
    if err != nil {
        return nil, err
    }
    return msgs, nil
}
