package main

import (
	"fmt"

	"github.com/newit-hieutm/go-backend/internal/jobs/redis_pools"
	"github.com/newit-hieutm/go-backend/internal/jobs/tasks"
)

// func init(){
// 	fmt.Println("RUN INIT")
// 	InitQueues()
// }

func main() {
	redisPool, err := redis_pools.NewRedisPool(redis_pools.RedisPoolOpts{
		DSN:          "redis://localhost:6379/0",
		MinIdleConns: 3,
	})

	if err != nil {
		panic("Can not create redis pools")
	}

	fmt.Println("Dia chi con tro cua redis pool la:", &redisPool)

	asynqClient := redis_pools.NewAsynqClient(redisPool)
	fmt.Println("Created success asynqClient", asynqClient)

	for i := 0; i < 50; i++ {
		t1, err := tasks.NewWelcomeEmailTask(i)

		if err != nil {
			panic(err)
		}
		fmt.Println("Created success t1", t1)
		// Process the task immediately.
		info, err := asynqClient.Enqueue(t1)
		if err != nil {
			panic(err)
		}
		fmt.Printf(" [*] Successfully enqueued task: %+v", info)
	}
}
