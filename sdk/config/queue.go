package config

import (
	"fmt"
	"log"
	"time"

	"github.com/nicelizhi/easy-admin-core/storage"
	"github.com/nicelizhi/easy-admin-core/storage/queue"

	"github.com/go-redis/redis/v7"
	"github.com/robinjoseph08/redisqueue/v2"
	//"github.com/redis/go-redis/v9"
)

type Queue struct {
	Redis  *QueueRedis
	Memory *QueueMemory
	NSQ    *QueueNSQ `json:"nsq" yaml:"nsq"`
}

type QueueRedis struct {
	RedisConnectOptions
	Producer *redisqueue.ProducerOptions
	Consumer *redisqueue.ConsumerOptions
}

type QueueMemory struct {
	PoolSize uint
}

type QueueNSQ struct {
	NSQOptions
	ChannelPrefix string
}

var QueueConfig = new(Queue)

// Empty 空设置
func (e Queue) Empty() bool {
	return e.Memory == nil && e.Redis == nil && e.NSQ == nil
}

// Setup 启用顺序 redis > 其他 > memory
func (e Queue) Setup() (storage.AdapterQueue, error) {
	if e.Redis != nil {
		e.Redis.Consumer.ReclaimInterval = e.Redis.Consumer.ReclaimInterval * time.Second
		e.Redis.Consumer.BlockingTimeout = e.Redis.Consumer.BlockingTimeout * time.Second
		e.Redis.Consumer.VisibilityTimeout = e.Redis.Consumer.VisibilityTimeout * time.Second
		client := GetRedisClient()
		//client * redis.Client
		if client == nil {
			options, err := e.Redis.RedisConnectOptions.GetRedisOptions()
			fmt.Printf("options " + options.Addr + "\r\n")

			if err != nil {
				log.Fatal(err)
				return nil, err
			}
			client = redis.NewClient(options)
			fmt.Printf("options " + options.Password + "\r\n")
			fmt.Printf("options client" + client.String())
			_redis = client
		}

		e.Redis.Producer.RedisClient = client
		e.Redis.Consumer.RedisClient = client
		//return queue.NewRedis()

		return queue.NewRedis(e.Redis.Producer, e.Redis.Consumer)
	}
	if e.NSQ != nil {
		cfg, err := e.NSQ.GetNSQOptions()
		if err != nil {
			return nil, err
		}
		return queue.NewNSQ(e.NSQ.Addresses, cfg, e.NSQ.ChannelPrefix)
	}
	return queue.NewMemory(e.Memory.PoolSize), nil
}
