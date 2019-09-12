package main

import "github.com/gin-gonic/gin"

func(s *Service) addBook(c *gin.Context) (int, interface{}){
	b := new(Book)
	err := c.ShouldBindJSON(&b)
	if err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	return s.makeSuccessJSON("添加成功！")
}
