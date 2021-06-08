package main

import (
	"2019-nCoV-Vaccine-XIAN/vaccine"
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始查询是否存在可预约新冠疫苗")
	// 抓包获取以下参数
	// 用户认证ID
	authId := ""
	// 企业ID
	enterpriseId := ""
	// 企业名称
	enterpriseName := ""
	// 经纬度
	longitude := "108.91616"
	latitude := "34.20484"
	v := vaccine.New(authId, enterpriseId, enterpriseName, longitude, latitude)
	for true {
		v.Start()
		time.Sleep(5 * time.Minute)
	}
}
