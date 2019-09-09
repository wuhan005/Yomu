package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/pelletier/go-toml"
)

type Service struct {
	Config		*toml.Tree
	Router		*gin.Engine
	Redis		*redis.Client
}

func (s *Service) init(){
	s.initConfig()
	s.initDatabase()
	s.initRouter()
}