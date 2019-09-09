package main

import (
	"github.com/pelletier/go-toml"
	"log"
)

func (s *Service) initConfig(){
	c, err := toml.LoadFile("conf/Yomu.toml")
	if err != nil {
		log.Fatal(err)
	}

	s.Config = c
}
