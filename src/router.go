package main

import "github.com/gin-gonic/gin"

func (s *Service) initRouter(){
	r := gin.Default()

	// Pages
	{
		r.GET("/", func(c *gin.Context){

		})
	}

	// RESTful APIs
	r.POST("/login", func(c *gin.Context) {
		c.JSON(s.login(c))
	})

	r.Use(s.AuthRequired())
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

		r.PUT("/book/:id", func(c *gin.Context) {
			c.JSON(s.editBook(c))
		})

		r.DELETE("/book/:id", func(c *gin.Context) {
			c.JSON(s.deleteBook(c))
		})

		r.GET("/books", func(c *gin.Context) {
			c.JSON(s.getBooks(c))
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

func(s *Service) AuthRequired() gin.HandlerFunc{
	return func(c *gin.Context){
		token := c.GetHeader("Authorization")
		if !s.checkToken(token){
			c.JSON(s.makeErrJSON(403, 40301, "无权访问"))
			c.Abort()
		}
	}
}