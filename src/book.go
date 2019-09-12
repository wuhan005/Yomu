package main

import "github.com/gin-gonic/gin"

func(s *Service) addBook(c *gin.Context) (int, interface{}){
	b := new(Book)
	err := c.ShouldBindJSON(&b)
	if err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	// Check whether repeat
	b1 := new(Book)
	s.Mysql.Where(&Book{Isbn: b.Isbn}).Find(&b1)
	if b1.Title != ""{
		return s.makeErrJSON(400, 40001, "该书籍已存在！")
	}

	tx := s.Mysql.Begin()
	if tx.Create(b).RowsAffected != 1{
		tx.Rollback()
		return s.makeErrJSON(500, 50001, "添加书籍失败！")
	}
	tx.Commit()
	return s.makeSuccessJSON("添加成功！")
}
