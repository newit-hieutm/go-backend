package main

import (
	"fmt"
	"log"

	"github.com/hibiken/asynq"
	"github.com/newit-hieutm/go-backend/internal/jobs/redis_pools"
	"github.com/newit-hieutm/go-backend/internal/jobs/tasks"
)

// workers.go
func main() {
	redisPool, err := redis_pools.NewRedisPool(redis_pools.RedisPoolOpts{
		DSN:          "redis://localhost:6379/0",
		MinIdleConns: 3,
	})

	if err != nil {
		panic("Can not get redis pools")
	}


	fmt.Println("Dia chi con tro cua redis pool la:", &redisPool)


	consumer := redis_pools.NewAsynqServer(redisPool, 3)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeWelcomeEmail, tasks.HandleWelcomeEmailTask)
	mux.HandleFunc(tasks.TypeReminderEmail, tasks.HandleReminderEmailTask)

	if err := consumer.Run(mux); err != nil {
		log.Fatal(err)
	}
}
