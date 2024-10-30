package redis_pools

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type (
	RedisPoolOpts struct {
		DSN          string
		MinIdleConns int
	}

	RedisPool struct {
		Client *redis.Client
	}
)

// NewRedisPool creates a reusable connection across your application.
// Note: Each db namespace requires its own connection pool. See go-redis for more info.
func NewRedisPool(o RedisPoolOpts) (*RedisPool, error) {
	redisOpts, err := redis.ParseURL(o.DSN)
	if err != nil {
		return nil, err
	}

	redisOpts.MinIdleConns = o.MinIdleConns

	fmt.Println("redisOpts", redisOpts)

	// Any setup related commands here

	return &RedisPool{
		Client: redis.NewClient(redisOpts),
	}, nil
}

// Interface adapter for asynq to resuse the same Redis connection pool.
func (r *RedisPool) MakeRedisClient() *redis.Client {
	return r.Client
}

// You can then pass *mypool.RedisPool directly to async.NewServer

func NewAsynqClient(redisPool *RedisPool) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisPool.Client.Options().Addr,
		Password: redisPool.Client.Options().Password,
		DB:       redisPool.Client.Options().DB,
	})
}

func NewAsynqServer(redisPool *RedisPool, concurency int) *asynq.Server {
	return asynq.NewServer(asynq.RedisClientOpt{
		Addr:     redisPool.Client.Options().Addr,
		Password: redisPool.Client.Options().Password,
		DB:       redisPool.Client.Options().DB,
	}, asynq.Config{Concurrency: concurency})
}
