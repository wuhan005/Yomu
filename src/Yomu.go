package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
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