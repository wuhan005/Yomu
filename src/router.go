package main

import "github.com/gin-gonic/gin"

func (s *Service) initRouter(){
	r := gin.Default()

	// RESTful APIs
	{
		r.GET("/search", func(c *gin.Context){
			
		})

		r.GET("/book", func(c *gin.Context){

		})

		r.POST("/book", func(c *gin.Context) {

		})

		r.PUT("/book", func(c *gin.Context) {

		})

		r.DELETE("/book", func(c *gin.Context) {

		})

		r.GET("/books", func(c *gin.Context) {

		})

		r.POST("/sign", func(c *gin.Context) {
			c.JSON(s.sign(c))
		})

		r.GET("/sign", func(c *gin.Context) {
			c.JSON(s.signHistory(c))
		})
	}

	s.Router = r
	_ = s.Router.Run(s.Config.Get("server.port").(string))
}