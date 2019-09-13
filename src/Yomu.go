package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/dgryski/dgoogauth"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	uuid "github.com/satori/go.uuid"
	"math"
	"strconv"
	"strings"
	"time"
)

func (s *Service) login(c *gin.Context) (int, interface{}){
	l := new(loginCode)
	err := c.ShouldBindJSON(&l)
	if err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	gaCode := s.Config.Get("key.gacode").(string)
	timestamp := time.Now().Unix() / int64(30)
	val := fmt.Sprintf("%06d", dgoogauth.ComputeCode(gaCode, timestamp))
	if val == l.Code{
		token := uuid.NewV4()
		s.Redis.Set("token", token, -1)
		return s.makeSuccessJSON(token)
	}
	return s.makeErrJSON(403, 40300, "登录失败！")
}

func (s *Service) sign(c *gin.Context) (int, interface{}){
	year := time.Now().Year()
	firstTimeStamp := time.Date(year, 1, 1, 0, 0, 0, 0, time.Now().Location()).Unix()
	nowTimeStamp := time.Now().Unix()

	days := math.Floor(float64((nowTimeStamp - firstTimeStamp) / (24 * 60 * 60)))
	checkStatus, _ := s.Redis.GetBit(fmt.Sprintf("sign_%d", year), int64(days)).Result()
	if checkStatus == 0{
		s.Redis.SetBit(fmt.Sprintf("sign_%d", year), int64(days),1)
		return s.makeSuccessJSON("签到成功！")
	}else{
		return s.makeErrJSON(403, 40300, "你今天已经签到过了！")
	}
}

func (s *Service) signHistory(c *gin.Context) (int, interface{}){
	y, exist := c.GetQuery("year")
	year, err := strconv.Atoi(y)

	if exist == false || err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}
	m, exist := c.GetQuery("month")
	month, err := strconv.Atoi(m)
	if exist == false || err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	if month <= 0 || month > 12{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	yearFirstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.Now().Location()).Unix()
	firstDay := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Now().Location()).Unix()
	lastDay := time.Date(year, time.Month(month + 1), 1, 0, 0, 0, 0, time.Now().Location()).Unix()

	days := uint(math.Floor(float64((firstDay - yearFirstDay) / (24 * 60 * 60))))
	monthCount := uint(math.Floor(float64((lastDay - firstDay) / (24 * 60 * 60))))
	fmt.Println(monthCount)

	resultNum, _ := s.Redis.BitField(fmt.Sprintf("sign_%d", year), "get", fmt.Sprintf("u%d", monthCount), days).Result()
	resultString := fmt.Sprintf("%0" + fmt.Sprintf("%d", monthCount) + "b", resultNum[0])
	result := strings.Split(resultString, "")
	return s.makeSuccessJSON(result)
}

func (s *Service) searchBook(c *gin.Context) (int, interface{}) {
	b:= new(searchBook)
	err := c.ShouldBindJSON(&b)
	if err != nil{
		return s.makeErrJSON(400, 40000, "数据格式错误")
	}

	// Check Redis cache
	cacheData, _ := s.Redis.Get(b.Isbn).Result()
	if cacheData != ""{
		cache, _ := simplejson.NewJson([]byte(cacheData))
		result := cache.MustMap()
		result["cache"] = true
		return s.makeSuccessJSON(result)
	}

	request := gorequest.New()
	_, body, errs := request.
		Get("http://jisuisbn.market.alicloudapi.com/isbn/query?isbn=" + b.Isbn).
		Set("Authorization", "APPCODE " + s.Config.Get("key.appcode").(string)).
		End()
	if errs != nil{
		return s.makeErrJSON(500, 50000, "获取书籍信息失败！")
	}

	data, err := simplejson.NewJson([]byte(body))
	if err != nil{
		return s.makeErrJSON(500, 50001, "ISBN 错误！")
	}

	if data.Get("status").MustInt() != 0{
		return s.makeErrJSON(500, 50001, "ISBN 错误！")
	}
	result := data.Get("result").MustMap()

	// Set cache result
	jsonData, _ := json.Marshal(result)
	s.Redis.Set(b.Isbn, jsonData, -1)

	result["cache"] = false
	return s.makeSuccessJSON(result)
}