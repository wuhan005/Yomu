package main

import "github.com/gin-gonic/gin"

func (s *Service) initRouter(){
	r := gin.Default()

	// RESTful APIs
	{
		r.POST("/search", func(c *gin.Context){
			c.JSON(s.searchBook(c))
		})

		r.GET("/book/:id", func(c *gin.Context){
			c.JSON(s.getBook(c))
		})

		r.POST("/book", func(c *gin.Context) {
			c.JSON(s.addBook(c))
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