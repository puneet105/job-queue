package main

import (
    "log"
	"time"
	"fmt"
    "net/http"
    "github.com/puneet105/job-queue/internal/auth"
    "github.com/puneet105/job-queue/internal/queue"
    "github.com/puneet105/job-queue/internal/worker"
    "github.com/puneet105/job-queue/internal/config"
)

func main() {
	var rabbitMQ *queue.RabbitMQ
	var redisQueue *queue.RedisQueue
    var err error
	publishMessages := []string{
        "Message 1 from Publisher 1",
        "Message 1 from Publisher 2",
		"Message 1 from Publisher 3",
    }
    cfg := config.LoadConfig()
    QueueName := cfg.QueueName
	for retries := 0; retries < 5; retries++ {
        rabbitMQ, err = queue.NewRabbitMQ(cfg)
        if err == nil {
            break
        }
        log.Printf("Failed to initialize RabbitMQ, retrying... (%d/10)", retries+1)
        time.Sleep(5 * time.Second)
    }
	if err != nil {
        log.Fatalf("Failed to initialize RabbitMQ after retries: %v", err)
    }

	channel, err := rabbitMQ.Conn.Channel()
    if err != nil {
    	fmt.Errorf("Failed to open a channel: %w", err)
    }

    _, err = channel.QueueDeclare(
        cfg.QueueName,    // name of the queue
        true,             // durable
        false,            // delete when unused
        false,            // exclusive
        false,            // no-wait
        nil,              // arguments
    )
    if err != nil {
    	fmt.Errorf("Failed to declare a queue: %w", err)
    }
    defer rabbitMQ.Conn.Close()

   
	for retries := 0; retries < 5; retries++ { 
		redisQueue, err = queue.NewRedisQueue(cfg)
		if err == nil {
			break
		}
		log.Printf("Failed to initialize Redis, retrying... (%d/10)", retries+1)
        time.Sleep(5 * time.Second)
	}
	if err != nil{
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
    defer redisQueue.Client.Close()

    
    go worker.ProcessRabbitMQJob(rabbitMQ, QueueName)
    go worker.ProcessRedisJob(redisQueue, QueueName)

    http.HandleFunc("/login", auth.LoginHandler) 

    http.Handle("/publish", auth.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
            return
        }

        for i, messages := range publishMessages {
            go func(i int, messages string) {
                err := rabbitMQ.Publish(QueueName, messages)
                if err != nil {
                    log.Printf("Publish %d failed: %v", i+1, err)
                }
            }(i, messages)
        }

        for i, messages := range publishMessages {
            go func(i int, messages string) {
                err = redisQueue.Publish(QueueName, messages)
                if err != nil {
                    log.Printf("Publish %d failed: %v", i+1, err)
                }
            }(i, messages)
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Published messages to all queues"))
    })))

    log.Println("Server is running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
