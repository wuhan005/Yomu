package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title 		string 		`json:"title"`
	Author 		string		`json:"author"`
	Cover 		string 		`json:"cover"`
	Isbn 		string 		`json:"isbn"`
	TotalPage	string		`json:"total_page"`
}

type ReadHistroy struct {
	gorm.Model

}

func (s *Service) makeErrJSON(httpStatusCode int, errCode int, msg interface{}) (int, interface{}) {
	return httpStatusCode, map[string]interface{}{"error": errCode, "msg": fmt.Sprint(msg)}
}

func (s *Service) makeSuccessJSON(data interface{}) (int, interface{}) {
	return 200, map[string]interface{}{"error": 0, "msg": "success", "data": data}
}