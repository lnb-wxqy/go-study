package validator_model

//通知结构体

type SubscribeNotificationModel struct {
	SubscribeNotificationListObject *SubscribeNotificationListObject `json:"SubscribeNotificationListObject"`
}

type SubscribeNotificationListObject struct {
	SubscribeNotificationObject []*SubscribeNotification `json:"SubscribeNotificationObject"`
}

type SubscribeNotification struct {
	NotificationID   string       `json:"NotificationID" validate:"required||string=33||regexp=13-14,04"` //通知标识,必填项，长度33，13-14为04 正则匹配：`^\d{12}04\d{19}$`
	SubscribeID      string       `json:"SubscribeID" validate:"required||string=33||regexp=13-14,03"`    //订阅标识,必填项，长度33，13-14为03 正则匹配：`^\d{12}03\d{19}$`
	Title            string       `json:"Title" validate:"required||string=_,256"`                        //订阅标题,必填项,最大长度256
	TriggerTime      string       `json:"TriggerTime" validate:"required||datetime"`                      //触发时间，格式示例：20201011123225,必填项
	InfoIDs          string       `json:"InfoIDs" validate:"required||string=_,1024"`                     //订阅通知的详细信息(人员、人脸、机动车、非机动车、物品、场景)标识集合,必填项,长度不能超过1024
	CaseObjectList   []*Case      `json:"CaseObjectList"`                                                 //视频案事件
	Tollgate         []*Tollgate  `json:"Tollgate"`                                                       //视频卡口信息数据集
	LaneList         []*Lane      `json:"Lane"`                                                           //车道信息数据集
	DeviceList       []*APE       `json:"DeviceList"`                                                     //设备信息数据集
	DeviceStatusList []*APEStatus `json:"DeviceStatusList"`                                               // 该通知针对批量订阅方式
	APSObjectList    []*APS       `json:"APSObjectList"`                                                  //设备网管信息数据集

	APSStatusObjectList       []*APSStatus               `json:"APSStatusObjectList"`       // 该通知针对批量订阅方式
	PersonObjectList          *PersonListObject          `json:"PersonObjectList"`          //人员信息数据集
	FaceObjectList            *FaceListObject            `json:"FaceObjectList"`            //人脸信息数据集
	MotorVehicleObjectList    *MotorVehicleListObject    `json:"MotorVehicleObjectList"`    // 机动车(过车)信息数据集
	NonMotorVehicleObjectList *NonMotorVehicleListObject `json:"NonMotorVehicleObjectList"` //非机动车数据集
	ThingObjectList           []*Thing                   `json:"ThingObjectList"`           //物品列表
	SceneObjectList           []*Scene                   `json:"SceneObjectList"`           //场景列表
	Email                     string                     `json:"email"`                     //测试

}

type Case struct {
	//TODO ...
}

type Tollgate struct {
	//TODO ...
}

type Lane struct {
	//TODO ...
}

type APE struct {
	//TODO ...
}

type APEStatus struct {
	//TODO ...
}

type APS struct {
	//TODO ...
}

type APSStatus struct {
	//TODO ...
}

type Thing struct {
	//TODO ...
}

type Scene struct {
	//TODO ...
}
