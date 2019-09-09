package main

import "github.com/go-redis/redis/v7"

func (s *Service) initDatabase(){
	// Redis
	r := redis.NewClient(&redis.Options{
		Addr:     s.Config.Get("database.redis_address").(string),
		Password: "",
		DB:       0,
	})

	s.Redis = r
	// Test connection
	_, err := s.Redis.Ping().Result()
	if err != nil{
		panic("Can't connect to Redis.")
	}

}
