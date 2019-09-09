package main

import "github.com/gin-gonic/gin"

func (s *Service) initRouter(){
	r := gin.Default()

	r.POST("/sign", func(c *gin.Context){
		c.JSON(s.sign(c))
	})

	s.Router = r
	_ = s.Router.Run(s.Config.Get("server.port").(string))
}