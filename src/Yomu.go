package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
	"strings"
	"time"
)

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