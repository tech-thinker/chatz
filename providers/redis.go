package providers

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/tech-thinker/chatz/config"
	"github.com/tech-thinker/chatz/constants"
	"github.com/tech-thinker/chatz/models"
)

type RedisProvider struct {
	config *config.Config
}

func (agent *RedisProvider) Post(message string, option models.Option) (any, error) {
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

func (agent *RedisProvider) Reply(threadId string, message string, option models.Option) (any, error) {
	fmt.Println("Reply to redis not supported yet.")
	return nil, errors.New("Reply to redis not supported yet.")
}

func (agent *RedisProvider) setup(env *config.Config) error {
	if env == nil {
		return fmt.Errorf("config is required")
	}
	if env.ConnectionURL == "" {
		return fmt.Errorf("connection_url is required for Redis provider")
	}
	if env.ChannelId == "" {
		return fmt.Errorf("channel_id is required for Redis provider")
	}
	agent.config = env
	return nil
}

func init() {
	RegisterProvider(constants.PROVIDER_REDIS, new(RedisProvider))
}
