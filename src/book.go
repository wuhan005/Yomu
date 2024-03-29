package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
	"strconv"
)

func(s *Service) getBook(c *gin.Context) (int, interface{}){
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	b := new(Book)
	s.Mysql.Where(&Book{Model: gorm.Model{ID: uint(bookID)}}).Find(b)
	if b.Title == ""{
		return s.makeErrJSON(404, 40400, "该书籍不存在！")
	}
	return s.makeSuccessJSON(b)
}

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

func(s *Service) editBook(c *gin.Context) (int, interface{}){
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	b := new(Book)
	s.Mysql.Where(&Book{Model: gorm.Model{ID: uint(bookID)}}).Find(b)
	if b.Title == ""{
		return s.makeErrJSON(404, 40400, "该书籍不存在！")
	}

	form := new(Book)
	err = c.ShouldBindJSON(&form)
	if err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	tx := s.Mysql.Begin()
	if tx.Model(&Book{}).Where(&Book{Model: gorm.Model{ID: b.ID}}).Updates(form).RowsAffected != 1{
		tx.Rollback()
		return s.makeErrJSON(500, 50001, "添加书籍失败！")
	}

	tx.Commit()
	return s.makeSuccessJSON("修改成功！")
}

func(s *Service) deleteBook(c *gin.Context) (int,interface{}){
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	b := new(Book)
	s.Mysql.Where(&Book{Model: gorm.Model{ID: uint(bookID)}}).Find(b)
	if b.Title == ""{
		return s.makeErrJSON(404, 40400, "该书籍不存在！")
	}

	tx := s.Mysql.Begin()
	if tx.Model(&Book{}).Where(&Book{Model: gorm.Model{ID: b.ID}}).Delete(&Book{Model: gorm.Model{ID: b.ID}}).RowsAffected != 1{
		tx.Rollback()
		return s.makeErrJSON(500, 50002, "删除书籍失败！")
	}
	tx.Commit()
	return s.makeSuccessJSON("删除成功！")
}

func(s *Service) getBooks(c *gin.Context) (int, interface{}){
	np, exist := c.GetQuery("page")
	if !exist{
		return s.makeErrJSON(400, 40002, "请输入页数")
	}
	nowPage, err := strconv.Atoi(np)
	if err != nil || nowPage <= 0{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	// 0 - all, 1 - reading, 2 - finish
	dt, exist := c.GetQuery("type")
	if !exist{
		return s.makeErrJSON(400, 40002, "请输入查询类型")
	}
	dataType, err := strconv.Atoi(dt)
	if err != nil || dataType < 0 || dataType > 2{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	var count int
	perPage := 8
	result := make(map[string]interface{})

	s.Mysql.Model(&Book{}).Count(&count)
	pages := math.Ceil(float64(count / perPage))

	books := make([]*Book, 0)
	s.Mysql.Where(&Book{Status: dataType}).Offset(nowPage * perPage).Limit(perPage).Find(&books)

	result["pages"] = pages
	result["books"] = books

	return s.makeSuccessJSON(result)
}