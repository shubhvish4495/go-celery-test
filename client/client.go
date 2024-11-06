package main

import (
	"github.com/gocelery/gocelery"
	"github.com/gomodule/redigo/redis"
)

func main() {

	redisPool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL("redis://localhost:6379")
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}

	// Initialize Redis broker
	redisBroker := gocelery.NewRedisBroker(redisPool)
	celeryBackend := &gocelery.RedisCeleryBackend{Pool: redisPool}

	// Initialize Celery client with Redis broker
	celeryClient, err :=
		gocelery.NewCeleryClient(redisBroker,
			celeryBackend,
			1)
	if err != nil {
		panic(err)
	}

	// Queue a task
	_, err = celeryClient.Delay("sub", 3, 5)
	if err != nil {
		panic(err)
	}

	_, err = celeryClient.Delay("add", 3, 5)
	if err != nil {
		panic(err)
	}

}
