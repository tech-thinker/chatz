package providers

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/tech-thinker/chatz/config"
)

type RedisProvider struct {
    config *config.Config
}

func (agent *RedisProvider) Post(message string) (interface{}, error) {
    options, err := redis.ParseURL(agent.config.ConnectionURL)
	if err != nil {
		return nil, err
	}
    
    rdb := redis.NewClient(options)
    ctx := context.Background()
	
    _, err = rdb.Ping(ctx).Result()
	if err != nil {
        return nil, err
	}

    err = rdb.Publish(ctx, agent.config.ChannelId, message).Err()
	if err != nil {
        return nil, err
	}
    
    return "Published", nil
}

func (agent *RedisProvider) Reply(threadId string, message string) (interface{}, error) {
    fmt.Println("Reply to redis not supported yet.")
    return nil, errors.New("Reply to redis not supported yet.")
}

