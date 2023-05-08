package ds

import (
	"context"
	"loan-back-services/conf"
	"loan-back-services/pkg/logger"
	"net"
	"time"

	"github.com/go-redis/redis/v8"
)

func LoadRDB() (*redis.Client, error) {
	//host := os.Getenv("REDIS_HOST")
	//port := os.Getenv("REDIS_PORT")
	//user := os.Getenv("REDIS_USER")
	//pass := os.Getenv("REDIS_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(conf.Redis().Host, conf.Redis().Port),
		Username: conf.Redis().Username,
		Password: conf.Redis().Password,
		DB:       0,
	})

	logger.Sugar.Info("Successfully connected to redis")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	return rdb, nil
}
