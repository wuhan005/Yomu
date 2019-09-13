package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type searchBook struct {
	Isbn		string		`json:"isbn"`
}

type loginCode struct {
	Code 		string		`json:"code"`
}

type Book struct {
	gorm.Model
	Title 		string 		`json:"title"`
	Author 		string		`json:"author"`
	Cover 		string 		`json:"cover"`
	Isbn 		string 		`json:"isbn"`
	Summary  	string		`json:"summary"`
	Publisher 	string		`json:"publisher"`
	TotalPage	string		`json:"total_page"`
	Status 		int			`json:"status"`
}

type ReadHistroy struct {
	gorm.Model
	BookId		uint		`json:"book_id"`
	Pages		string		`json:"pages"`
}

func (s *Service) makeErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, map[string]interface{}{"error": errCode, "msg": fmt.Sprint(msg)}
}

func (s *Service) makeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, map[string]interface{}{"error": 0, "msg": "success", "data": data}
}