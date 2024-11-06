package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

// Define a sample task function
func add(a int, b int) int {
	fmt.Println("job is running")
	return a + b
}

func sub(a, b int) int {
	fmt.Println("sub func is running")
	return a - b
}

func main() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}

	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	// Initialize Redis broker
	redisBroker := gocelery.NewRedisBroker(redisPool)
	celeryBackend := &gocelery.RedisCeleryBackend{Pool: redisPool}

	// Initialize Celery worker
	celeryWorker := gocelery.NewCeleryWorker(redisBroker, celeryBackend, 1)

	// Register task with the worker
	celeryWorker.Register("sub", sub)

	celeryWorker.Register("add", add)

	// Start the worker
	fmt.Println("Starting GoCelery worker...")
	defer celeryWorker.StopWorker()
	celeryWorker.StartWorker()

	time.Sleep(300 * time.Second)
}
