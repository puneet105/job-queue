package queue

import (
    "github.com/go-redis/redis/v8"
    "context"
    "github.com/puneet105/job-queue/internal/config"
    // "gopkg.in/yaml.v2"
    // "os"
    "fmt"
)

type RedisQueue struct {
    Client *redis.Client
}

func NewRedisQueue(config *config.Config) (*RedisQueue, error) {
    // config := struct {
    //     Redis struct {
    //         Addr     string `yaml:"addr"`
    //         Password string `yaml:"password"`
    //         DB       int    `yaml:"db"`
    //     } `yaml:"redis"`
    // }{}

    // data, err := os.ReadFile("config.yaml")
    // if err != nil {
    //     return nil, err
    // }

    // err = yaml.Unmarshal(data, &config)
    // if err != nil {
    //     return nil, err
    // }

    // client := redis.NewClient(&redis.Options{
    //     Addr:     config.Redis.Addr,
    //     Password: config.Redis.Password,
    //     DB:       config.Redis.DB,
    // })

    redisURL := fmt.Sprintf("%s:%s", config.RedisHost, config.RedisPort)
    client := redis.NewClient(&redis.Options{
        Addr:     redisURL,
        Password: config.RedisPassword,
        DB:       0,
    })

    return &RedisQueue{Client: client}, nil
}

func (r *RedisQueue) Publish(queueName string, message string) error {
    ctx := context.Background()
     
    err := r.Client.LPush(ctx, queueName, message)
    fmt.Printf("Published message to Redis: %s\n", message)
    return err.Err()
}

func (r *RedisQueue) Consume(queueName string) (chan string, error) {
    ctx := context.Background()
    ch := make(chan string)
    go func() {
        for {
            result, err := r.Client.BRPop(ctx, 0, queueName).Result()
            if err != nil {
                continue
            }
            ch <- result[1]
        }
    }()

    return ch, nil
}
