package worker

import (
	"fmt"
	"log"
	"github.com/puneet105/job-queue/internal/queue"
)

func ProcessRabbitMQJob(rabbitMQ *queue.RabbitMQ, queueName string) {
	msgs, err := rabbitMQ.Consume(queueName)
	if err != nil {
		log.Fatalf("Failed to consume RabbitMQ messages: %v", err)
	}

	for msg := range msgs {
		fmt.Printf("RabbitMQ: Received a message: %s\n", msg.Body)
	}
}

func ProcessKafkaJob(kafkaQueue *queue.KafkaQueue, topic string) {
	msgs, err := kafkaQueue.Consume(topic)
	if err != nil {
		log.Fatalf("Failed to consume Kafka messages: %v", err)
	}

	for msg := range msgs {
		fmt.Printf("Kafka: Received a message: %s\n", string(msg.Value))	
	}
}

func ProcessRedisJob(redisQueue *queue.RedisQueue, queueName string) {
	msgs, err := redisQueue.Consume(queueName)
	if err != nil {
		log.Fatalf("Failed to consume Redis messages: %v", err)
	}

	for msg := range msgs {
		fmt.Printf("Redis: Received a message: %s\n", msg)	
	}
}
