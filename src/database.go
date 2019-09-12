package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (s *Service) initDatabase(){
	// MySQL
	m, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.Config.Get("database.mysql_user").(string),
		s.Config.Get("database.mysql_password").(string),
		s.Config.Get("database.mysql_address").(string),
		s.Config.Get("database.mysql_database").(string),
	))

	if err != nil{
		panic("Can't connect to MySQL.")
	}
	s.Mysql = m

	// Create tables
	s.Mysql.AutoMigrate(&Book{}, &ReadHistroy{})

	// Redis
	r := redis.NewClient(&redis.Options{
		Addr:     s.Config.Get("database.redis_address").(string),
		Password: "",
		DB:       0,
	})

	s.Redis = r
	// Test connection
	_, err = s.Redis.Ping().Result()
	if err != nil{
		panic("Can't connect to Redis.")
	}
}
