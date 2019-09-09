package main

import "github.com/gin-gonic/gin"

func (s *Service) initRouter(){
	s.Router = gin.Default()

	_ = s.Router.Run(s.Config.Get("server.port").(string))
}
