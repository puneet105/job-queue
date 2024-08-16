package queue

import (
    "gopkg.in/yaml.v2"
    "github.com/IBM/sarama"
    "os"
)

type KafkaQueue struct {
    Producer sarama.SyncProducer
    Consumer sarama.Consumer
}

func NewKafkaQueue() (*KafkaQueue, error) {
    config := struct {
        Kafka struct {
            Brokers []string `yaml:"brokers"`
        } `yaml:"kafka"`
    }{}

    data, err := os.ReadFile("config.yaml")
    if err != nil {
        return nil, err
    }

    err = yaml.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }

    producer, err := sarama.NewSyncProducer(config.Kafka.Brokers, nil)
    if err != nil {
        return nil, err
    }

    consumer, err := sarama.NewConsumer(config.Kafka.Brokers, nil)
    if err != nil {
        return nil, err
    }

    return &KafkaQueue{Producer: producer, Consumer: consumer}, nil
}

func (k *KafkaQueue) Publish(topic string, message []byte) error {
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.ByteEncoder(message),
    }
    _, _, err := k.Producer.SendMessage(msg)
    return err
}

func (k *KafkaQueue) Consume(topic string) (<-chan *sarama.ConsumerMessage, error) {
    partitionConsumer, err := k.Consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
    if err != nil {
        return nil, err
    }

    return partitionConsumer.Messages(), nil
}
