package model

type SubscribeModel struct {
	SubscribeListObject *SubscribeListObject `json:"SubscribeListObject"`
}

type SubscribeListObject struct {
	SubscribeObject []*Subscribe `json:"SubscribeObject"`
}

type SubscribeObject struct {
	SubscribeObject *Subscribe `json:"SubscribeObject"`
}

type Subscribe struct {
	SubscribeID     string `json:"SubscribeID" db:"SubscribeID"`         //订阅标识,必填项
	Title           string `json:"Title" db:"Title"`                     //必填项
	SubscribeDetail string `json:"SubscribeDetail" db:"SubscribeDetail"` //必填项
	/*
	* 资源路径URI，必填项，多个用逗号隔开，长度256，设备国标长度20位，最多每次订阅12个设备
	* (卡口ID、设备ID、采集内容ID、案件ID、目标视图库ID、行政区编号2/4/6位等)
	 */
	ResourceURI           string `json:"ResourceURI" db:"ResourceURI"`                     //必填项
	ApplicantName         string `json:"ApplicantName" db:"ApplicantName"`                 //申请人,必填项
	ApplicantOrg          string `json:"ApplicantOrg" db:"ApplicantOrg"`                   //申请单位,必填项
	BeginTime             string `json:"BeginTime" db:"BeginTime"`                         //开始时间,必填项
	EndTime               string `json:"EndTime" db:"EndTime"`                             //结束时间,必填项
	ReceiveAddr           string `json:"ReceiveAddr" db:"ReceiveAddr"`                     //订阅信息接收地址URL如： http://x.x.x.x/receive1，必填项
	ReportInterval        int32  `json:"ReportInterval" db:"ReportInterval"`               //信息上报时间间隔,<=0表示不限制
	Reason                string `json:"Reason" db:"Reason"`                               //订阅理由
	OperateType           int32  `json:"OperateType" db:"OperateType"`                     //操作类型，0，订阅 1，取消订阅
	SubscribeStatus       int32  `json:"SubscribeStatus" db:"SubscribeStatus"`             //订阅执行状态 0：订阅中 1：已取消订阅 2：订阅到期 9：未订阅 该字段只读
	SubscribeCancelOrg    string `json:"SubscribeCancelOrg" db:"SubscribeCancelOrg"`       //订阅取消单位，仅在取消订阅时使用
	SubscribeCancelPerson string `json:"SubscribeCancelPerson" db:"SubscribeCancelPerson"` //订阅取消人，仅在取消订阅时使用
	CancelTime            string `json:"CancelTime" db:"CancelTime"`                       //订阅取消时间，仅在取消订阅时使用
	CancelReason          string `json:"CancelReason" db:"CancelReason"`                   //订阅取消原因，仅在取消订阅时使用
	ResourceClass         int32  `json:"ResourceClass" db:"ResourceClass"`                 //订阅资源类别  订阅时必选，指定订阅的资源类别  0-卡口 1-设备 2-采集内容 3-案件 4-视图库 5-行政区划
	ResultFeatureDeclare  int32  `json:"ResultFeatureDeclare" db:"ResultFeatureDeclare"`   //是否返回特征值  1需要 -1不需要
	ResultImageDeclare    string `json:"ResultImageDeclare" db:"ResultImageDeclare"`       //返回结果图片类型 01 车辆大图 02 车牌彩色小图 	03 车牌二值化图 04 驾驶员面部特征图 10 人员图 11 人脸图 12 非机动车图 14 场景图
	TabID                 string `json:"TabID" db:"TabID"`                                 //订阅分类标签标识
}

type SubscribeLibListObject struct {
	SubscribeLib []*SubscribeLib `json:"SubscribeLibListObject"`
}
type SubscribeLib struct {
	*Subscribe
	ManagerID   string `json:"ManagerID" db:"ManagerID"`     //视图库ID
	ManagerName string `json:"ManagerName" db:"ManagerName"` //视图库名称
	UpdateTime  string `json:"UpdateTime" db:"UpdateTime"`   //更新时间
	InsertTime  string `json:"InsertTime" db:"InsertTime"`   //插入时间
}
