package vaccine

import (
	"crypto/tls"
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Vaccine struct {
	Request  *gorequest.SuperAgent
	Headers  *http.Header
	Url1     string
	Url2     string
	ReqData1 *ReqData1
	ReqData2 *ReqData2

	ResData1 *ResData1
	ResData2 *ResData2

	Date           string
	EnterpriseId   string
	EnterpriseName string
}

var header = map[string]string{
	"Host":         "yqpt.xa.gov.cn",
	"User-Agent":   "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36 MicroMessenger/7.0.9.501 NetType/WIFI MiniProgramEnv/Windows WindowsWechat",
	"content-type": "application/json",
	"Referer":      "https://servicewechat.com/wx45ffe0eb50d02e59/36/page-frame.html",
}

var url1 = "https://yqpt.xa.gov.cn/neusoft-appt/appt-vfic/app/getApptProjectList?orderBy=&orderSeq=&projectType=1&longitude=%s&latitude=%s&collectLocationName=&pageNum=1&pageSize=100"
var url2 = "https://yqpt.xa.gov.cn/neusoft-appt/appt-vfic/app/getProjectInfoById"

func New(authId, enterpriseId, enterpriseName, longitude, latitude string) *Vaccine {
	headerTmp := http.Header{}
	for index, item := range header {
		headerTmp.Set(index, item)
	}
	headerTmp.Set("Authorization", authId)
	return &Vaccine{
		// 不配置会清除header
		// see https://github.com/parnurzeal/gorequest/issues/257
		Request: gorequest.New().SetDoNotClearSuperAgent(true),
		Headers: &headerTmp,
		Url1:    fmt.Sprintf(url1, longitude, latitude),
		Url2:    url2,
		ReqData1: &ReqData1{
			OrderBy:             "",
			OrderSeq:            "",
			ProjectType:         1,
			Longitude:           longitude,
			Latitude:            latitude,
			CollectLocationName: "",
			PageNum:             1,
			PageSize:            100,
		},
		ReqData2: &ReqData2{
			Date:      time.Now().Format("2006-01-02"),
			ProjectId: "",
		},
		ResData1:       new(ResData1),
		ResData2:       new(ResData2),
		Date:           time.Now().Format("2006-01-02"),
		EnterpriseId:   enterpriseId,
		EnterpriseName: enterpriseName,
	}
}

func (v *Vaccine) Start() *Vaccine {
	fmt.Println("开始本次查询...")
	v.Request.Header = *v.Headers
	// 关闭ssl校验
	v.Request.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	_, s, errors := v.Request.Post(v.Url1).SendStruct(*v.ReqData1).End()
	if errors != nil {
		fmt.Println(errors)
		return v
	}
	err := gjson.DecodeTo(s, v.ResData1)
	if err != nil {
		fmt.Println(errors)
		return v
	}
	v.findVaccine()
	fmt.Println("本次查询结束")
	return v
}

func (v *Vaccine) findVaccine() {
	listData := v.ResData1.Data
	for _, s := range listData {
		collectLocationName := s.CollectLocationName
		projectId := s.ProjectId
		data := ReqData2{
			Date:      v.ReqData2.Date,
			ProjectId: projectId,
		}
		_, s2, errors := v.Request.Post(v.Url2).SendStruct(data).End()
		if errors != nil {
			fmt.Println(s2)
			fmt.Println(errors)
			continue
		}
		res := new(ResData2)
		if err := gjson.DecodeTo(s2, &res); err != nil {
			fmt.Println(s2)
			fmt.Println(err)
			continue
		}
		distance, err := strconv.ParseFloat(s.Distance, 64)
		if err != nil {
			fmt.Println("距离转换失败")
			continue
		}
		if distance > 4 {
			//fmt.Println("距离太远, pass")
			continue
		}
		Dtime, _ := time.Parse("2006-01-02 15:04:05", v.Date+" 18:00:00")
		for _, itemList := range res.Data.DayList {
			for _, item := range itemList.Time {
				if item.SumCount-item.ApptCount > 0 {
					fmt.Printf("%s: %s 有疫苗(%d)\n", collectLocationName, item.StartTime, item.SumCount-item.ApptCount)
					timeTmp, _ := time.Parse("2006-01-02 15:04:05", v.Date+" "+item.StartTime+":00")
					if Dtime.Unix() <= timeTmp.Unix() {
						fmt.Println("尝试预约!")
						b := v.appt(item.StartTime, item.EndTime, s.Address, s.CollectLocationName, s.CollectLocationId, s.DeptId,
							v.EnterpriseId, v.EnterpriseName, item.ProjectApptId, s.ProjectId, "1")
						if b {
							os.Exit(-1)
						}
					}
				} else {
					fmt.Printf("%s: %s时已无疫苗\n", collectLocationName, item.StartTime)
				}
			}
		}
		time.Sleep(time.Second)
	}
}

// appt relation可能为第几次接种
func (v *Vaccine) appt(stime, etime, caddress, cname, cid, deptid, eid, ename, papptid, pid, relation string) bool {
	url := "https://yqpt.xa.gov.cn/neusoft-appt/appt-vfic/app/personAppt"
	book := map[string]interface{}{
		"askAboutJson": "{\"content\":[{\"q\":\"01.近一周有发热等不舒服吗？\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"02.是否对药物、食物、疫苗等过敏？\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"03.是否曾经在接种疫苗后出现过严重反应？\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"04.是否有癫痫、脑或其他神经系统疾病？\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"05.是否患有癌症、白血病、艾滋病或其他免疫系统疾病？\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"06.在过去三个月内，是否使用过可的松、强的松、其他类固醇或抗肿瘤药物，或进行过放射性治疗？\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"07.有哮喘、肺部疾病、心脏疾病、肾脏疾病、代谢性疾病（如糖尿病）或血液系统疾病吗？\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"08.在过去的一年内，是否接受过输血或血液制品、或使用过免疫球蛋白？\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"09.是否怀孕或有可能3个月内怀孕？（仅需询问育龄妇女）\",\"o\":[\"是\",\"否\"],\"a\":\"否\",\"m\":\"\"},{\"q\":\"10.其他\",\"o\":[\"是\",\"否\"],\"a\":\"\",\"m\":\"\",\"t\":\"input\"}],\"confirm\":true}",
		"knowJson":     "{\"content\":\"\",\"confirm\":true}",
		"promisesJson": "{\"content\":\"\",\"confirm\":true}",
	}
	data := ReqAppt{
		ApptDate:               v.Date,
		ApptTime:               stime + "-" + etime,
		CollectLocationAddress: caddress,
		CollectLocationName:    cname,
		CollectLocationId:      cid,
		DeptId:                 deptid,
		EnterpriseId:           eid,
		EnterpriseName:         ename,
		ProjectApptId:          papptid,
		ProjectId:              pid,
		Relation:               relation,
		Book:                   book,
	}
	_, s, errors := v.Request.Post(url).SendStruct(data).End()
	if errors != nil {
		fmt.Println(errors)
		return false
	}
	res := new(ResAppt)
	if err := gjson.DecodeTo(s, &res); err != nil {
		fmt.Println(s)
		fmt.Println(err)
		return false
	}
	fmt.Println(s)
	if res.Msg == "预约成功,请务必携带身份证去进行接种" {
		fmt.Println("预约成功!")
		return true
	}
	return false
}
