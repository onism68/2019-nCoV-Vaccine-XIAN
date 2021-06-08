package vaccine

type ReqData1 struct {
	OrderBy             string `json:"orderBy"`
	OrderSeq            string `json:"orderSeq"`
	ProjectType         int    `json:"projectType"`
	Longitude           string `json:"longitude"`
	Latitude            string `json:"latitude"`
	CollectLocationName string `json:"collectLocationName"`
	PageNum             int    `json:"pageNum"`
	PageSize            int    `json:"pageSize"`
}

type ReqData2 struct {
	Date      string `json:"date"`
	ProjectId string `json:"projectId"`
}

type ResData1 struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data []struct {
		SearchValue interface{} `json:"searchValue"`
		CreateBy    string      `json:"createBy"`
		CreateTime  string      `json:"createTime"`
		UpdateBy    string      `json:"updateBy"`
		UpdateTime  *string     `json:"updateTime"`
		Remark      string      `json:"remark"`
		Params      struct {
		} `json:"params"`
		ProjectId           string      `json:"projectId"`
		CollectLocationId   string      `json:"collectLocationId"`
		CollectLocationName string      `json:"collectLocationName"`
		NickName            interface{} `json:"nickName"`
		FullName            string      `json:"fullName"`
		TypeId              interface{} `json:"typeId"`
		TypeName            interface{} `json:"typeName"`
		AppointmentSum      string      `json:"appointmentSum"`
		Address             string      `json:"address"`
		Longitude           float64     `json:"longitude"`
		Latitude            float64     `json:"latitude"`
		Detail              interface{} `json:"detail"`
		Thumbnail           string      `json:"thumbnail"`
		WorkTime            string      `json:"workTime"`
		WorkDate            string      `json:"workDate"`
		DeptId              string      `json:"deptId"`
		DeptName            *string     `json:"deptName"`
		UserId              interface{} `json:"userId"`
		IsDeleted           string      `json:"isDeleted"`
		Distance            string      `json:"distance"`
		VaccinateTime       int         `json:"vaccinateTime"`
		IntervalTime        int         `json:"intervalTime"`
		VaccineInfo         interface{} `json:"vaccineInfo"`
		LimitDeptFlag       string      `json:"limitDeptFlag"`
		LimitEnterpriseFlag interface{} `json:"limitEnterpriseFlag"`
		EnterpriseIds       interface{} `json:"enterpriseIds"`
	} `json:"data"`
}

type ResData2 struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		ProjectId string `json:"projectId"`
		DayList   []struct {
			ProjectId string `json:"projectId"`
			Year      string `json:"year"`
			Month     string `json:"month"`
			Day       string `json:"day"`
			Time      []struct {
				StartTime     string `json:"startTime"`
				EndTime       string `json:"endTime"`
				ProjectApptId string `json:"projectApptId"`
				SumCount      int    `json:"sumCount"`
				ApptCount     int    `json:"apptCount"`
			} `json:"time"`
		} `json:"dayList"`
	} `json:"data"`
}

type ReqAppt struct {
	ApptDate               string      `json:"apptDate"`
	ApptTime               string      `json:"apptTime"`
	CollectLocationAddress string      `json:"collectLocationAddress"`
	CollectLocationName    string      `json:"collectLocationName"`
	CollectLocationId      string      `json:"collectLocationId"`
	DeptId                 string      `json:"deptId"`
	EnterpriseId           string      `json:"enterpriseId"`
	EnterpriseName         string      `json:"enterpriseName"`
	ProjectApptId          string      `json:"projectApptId"`
	ProjectId              string      `json:"projectId"`
	Relation               string      `json:"relation"`
	Book                   interface{} `json:"book"`
}

type ResAppt struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		SearchValue interface{} `json:"searchValue"`
		CreateBy    interface{} `json:"createBy"`
		CreateTime  interface{} `json:"createTime"`
		UpdateBy    interface{} `json:"updateBy"`
		UpdateTime  interface{} `json:"updateTime"`
		Remark      interface{} `json:"remark"`
		Params      struct {
		} `json:"params"`
		PersonApptId           string      `json:"personApptId"`
		PersonId               string      `json:"personId"`
		PersonName             string      `json:"personName"`
		PersonPhone            string      `json:"personPhone"`
		IdCard                 string      `json:"idCard"`
		ProjectApptId          string      `json:"projectApptId"`
		ApptDate               string      `json:"apptDate"`
		ApptTime               string      `json:"apptTime"`
		DeptId                 string      `json:"deptId"`
		DeptName               interface{} `json:"deptName"`
		EnterpriseId           string      `json:"enterpriseId"`
		EnterpriseName         string      `json:"enterpriseName"`
		OrderNum               string      `json:"orderNum"`
		CollectLocationId      string      `json:"collectLocationId"`
		CollectLocationName    string      `json:"collectLocationName"`
		CollectLocationAddress string      `json:"collectLocationAddress"`
		HeaderPersonId         string      `json:"headerPersonId"`
		Relation               string      `json:"relation"`
		Status                 int         `json:"status"`
		ProcessType            interface{} `json:"processType"`
		AddTime                string      `json:"addTime"`
		DelFlag                interface{} `json:"delFlag"`
		Code                   interface{} `json:"code"`
		ProcessStatus          interface{} `json:"processStatus"`
		Book                   struct {
			BookId                     string      `json:"bookId"`
			PersonId                   string      `json:"personId"`
			PersonApptId               string      `json:"personApptId"`
			PromisesJson               string      `json:"promisesJson"`
			KnowJson                   string      `json:"knowJson"`
			AskAboutJson               string      `json:"askAboutJson"`
			AddTime                    string      `json:"addTime"`
			PromisesPdf                interface{} `json:"promisesPdf"`
			PromisesQiyuesuoDocumentId interface{} `json:"promisesQiyuesuoDocumentId"`
			PromisesQiyuesuoContractId interface{} `json:"promisesQiyuesuoContractId"`
			KnowPdf                    interface{} `json:"knowPdf"`
			KnowQiyuesuoDocumentId     interface{} `json:"knowQiyuesuoDocumentId"`
			KnowQiyuesuoContractId     interface{} `json:"knowQiyuesuoContractId"`
			AskAboutPdf                interface{} `json:"askAboutPdf"`
			AskAboutQysDocumentId      interface{} `json:"askAboutQysDocumentId"`
			AskAboutQysContractId      interface{} `json:"askAboutQysContractId"`
		} `json:"book"`
		ObservationEndTime interface{} `json:"observationEndTime"`
		CurrentTime        interface{} `json:"currentTime"`
		LastInclude        interface{} `json:"lastInclude"`
		AskAboutPdf        interface{} `json:"askAboutPdf"`
		Project            interface{} `json:"project"`
	} `json:"data"`
}
