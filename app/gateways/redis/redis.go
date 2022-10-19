package redis

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	maxIdle     = 3
	idleTimeout = 240
)

func ConnectPool(url string) (*redis.Pool, error) {
	pool := &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			log.Printf("attempting to connect to redis on %s...\n", url)

			c, err := redis.DialURL(url)
			if err != nil {
				return redis.Dial("tcp", url)
			}

			return c, err
		},
	}

	conn := pool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil { // check if redis server is responsive
		return nil, fmt.Errorf("failed to ping redis server: %w", err)
	}
	log.Println("successfully connected to redis server")

	return pool, nil
}
