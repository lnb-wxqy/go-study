package validator_model

type NonMotorVehicleModel struct {
	NonMotorVehicleListObject *NonMotorVehicleListObject `json:"NonMotorVehicleListObject" validate:"required"`

	ProxyType         string `json:"proxyType,omitempty"`         //数据格式标识，0:GAT1400 格式，1:私有格式(扩展GAT1400)
	ProxyManufacturer string `json:"proxyManufacturer,omitempty"` //厂商编码
	ResJson           string `json:"resJson,omitempty"`           //扩展字段
}

type NonMotorVehicleListObject struct {
	NonMotorVehicleObject []*NonMotorVehicleObject `json:"NonMotorVehicleObject" validate:"required"`
}

type NonMotorVehicleObject struct {
	NonMotorVehicleID string                      `json:"NonMotorVehicleID" validate:"required||string=48||regexp=21-22,02"` //车辆标识 -备注：辆全局唯一标识 R -必选 正则匹配：^\d{20}02\d{26}$
	TollgateID        string                      `json:"TollgateID,omitempty"`                                              //关联卡口编码 O -可选
	VehicleType       string                      `json:"VehicleType,omitempty"`                                             //车辆款型
	FaceID            string                      `json:"FaceID,omitempty" validate:"string=48"`                             //人脸标识
	PersonID          string                      `json:"PersonID,omitempty" validate:"string=48"`                           //人员标识
	VehicleClass      string                      `json:"VehicleClass,omitempty"`                                            //车辆类型
	LaneNo            int32                       `json:"LaneNo,omitempty"`                                                  //车道号
	Direction         string                      `json:"Direction,omitempty"`                                               //行驶方向
	BreakRuleMode     string                      `json:"BreakRuleMode,omitempty"`                                           //违章类型
	PersonInfoList    PersonInfoOfVehicleTypeList `json:"PersonInfoList,omitempty"`                                          //人体属性信息
	CyclingType       int                         `json:"CyclingType,omitempty"`                                             //骑车类型
	//IsModified        string                      `json:"IsModified,omitempty"`     //改装标志
	Gender         int32  `json:"Gender,omitempty" validate:"in=0,1,2,9"` //性别
	TraceInfo      string `json:"TraceInfo,omitempty"`                    //非机动车轨迹
	Ethnic         int32  `json:"Ethnic,omitempty"`                       //民族
	NonmotorAngle  int32  `json:"NonmotorAngle,omitempty"`                //非机动车角度
	MoveDirection  int32  `json:"MoveDirection,omitempty"`                //移动方向
	IllageBehavior int32  `json:"IllageBehavior,omitempty"`               //载人情况

	InfoKind      int    `json:"InfoKind,omitempty" validate:"required||in=0,1"`    //信息分类 -备注：人工采集还是自动采集 R -
	SourceID      string `json:"SourceID,omitempty" validate:"required||string=41"` //来源标识 -备注：来源图像标识 R -必选
	DeviceID      string `json:"DeviceID,omitempty" validate:"required||string=20"` //设备编码 R/O -自动采集时必选
	LeftTopX      int32  `json:"LeftTopX,omitempty"`                                //左上角X坐标 车的轮廓外界矩形在画面中的位置，记录坐标 R/O 自动采集时必选
	LeftTopY      int32  `json:"LeftTopY,omitempty"`                                //左上角Y坐标 车的轮廓外界矩形在画面中的位置，记录坐标 R/O 自动采集时必选
	RightBtmX     int32  `json:"RightBtmX,omitempty"`                               //右上角X坐标 车的轮廓外界矩形在画面中的位置，记录坐标 R/O 自动采集时必选
	RightBtmY     int32  `json:"RightBtmY,omitempty"`                               //右上角Y坐标 车的轮廓外界矩形在画面中的位置，记录坐标 R/O 自动采集时必选;
	MarkTime      string `json:"MarkTime,omitempty"`                                //位置标记时间 O 人工采集时有效;
	AppearTime    string `json:"AppearTime,omitempty"`                              //车辆出现时间 O 人工采集时有效
	DisAppearTime string `json:"DisAppearTime,omitempty"`                           //车辆消失时间 O 人工采集时有效;
	HasPlate      string `json:"HasPlate,omitempty" validate:"required"`            //有无车牌号;
	PlateClass    string `json:"PlateClass,omitempty" validate:"required"`          //号牌种类;
	PlateColor    string `json:"PlateColor,omitempty" validate:"required"`          //车牌颜色;
	PlateNo       string `json:"PlateNo,omitempty" validate:"required"`             //车牌号;
	PlateNoAttach string `json:"PlateNoAttach,omitempty"`                           //挂车牌号;
	PlateDescribe string `json:"PlateDescribe,omitempty"`                           //车牌描述;
	//IsDecked            string             `json:"IsDecked,omitempty"`            //是否套牌;
	//IsAltered           string             `json:"IsAltered,omitempty"`           //是否涂改;
	//IsCovered           string             `json:"IsCovered,omitempty"`           //是否遮挡;
	Speed               float64 `json:"Speed,omitempty"`                            //行驶速度;
	DrivingStatusCode   string  `json:"DrivingStatusCode,omitempty"`                //行驶状态代码;
	UsingPropertiesCode int     `json:"UsingPropertiesCode,omitempty"`              //车辆使用性质代码;
	VehicleBrand        string  `json:"VehicleBrand,omitempty"`                     //车辆品牌;
	VehicleModel        string  `json:"VehicleModel,omitempty"`                     //车辆型号;
	VehicleLength       int     `json:"VehicleLength,omitempty"`                    //车辆长度;
	VehicleWidth        int     `json:"VehicleWidth,omitempty"`                     //车辆宽度;
	VehicleHeight       int     `json:"VehicleHeight,omitempty"`                    //车辆高度;
	VehicleColor        string  `json:"VehicleColor,omitempty" validate:"required"` //车身颜色 R 必选;
	VehicleHood         string  `json:"VehicleHood,omitempty"`                      //车前盖;
	VehicleTrunk        string  `json:"VehicleTrunk,omitempty"`                     //车后盖;
	VehicleWheel        string  `json:"VehicleWheel,omitempty"`                     //车轮;
	WheelPrintedPattern string  `json:"WheelPrintedPattern,omitempty"`              //车轮印花纹;
	VehicleWindow       string  `json:"VehicleWindow,omitempty"`                    //车窗;
	VehicleRoof         string  `json:"VehicleRoof,omitempty"`                      //车顶;
	VehicleDoor         string  `json:"VehicleDoor,omitempty"`                      //车门;
	SideOfVehicle       string  `json:"SideOfVehicle,omitempty"`                    //车侧;
	CarOfVehicle        string  `json:"CarOfVehicle,omitempty"`                     //车厢;
	RearviewMirror      string  `json:"RearviewMirror,omitempty"`                   //后视镜;
	VehicleChassis      string  `json:"VehicleChassis,omitempty"`                   //底盘;
	VehicleShielding    string  `json:"VehicleShielding,omitempty"`                 //遮挡;
	FilmColor           string  `json:"FilmColor,omitempty"`                        //贴膜颜色;
	StorageURL          string  `json:"StorageURL,omitempty"`                       //大图（场景图）路径;
	NationalityCode     string  `json:"NationalityCode,omitempty"`                  //NationalityCode;
	TabID               string  `json:"TabID,omitempty"`                            //TabID;
	RelatedType         string  `json:"RelatedType,omitempty"`                      //关联关系类型【海康提供的标准】;
	Longitude           float64 `json:"Longitude,omitempty"`                        //设备经度【固定点位设备可选填，移动设备必填】;
	Latitude            float64 `json:"Latitude,omitempty"`                         //设备纬度【固定点位设备可选填，移动设备必填】
	resJson             string  `json:"resJson,omitempty"`                          //预留扩展字段
	//SubImageList        *base.SubImageList `json:"SubImageList,omitempty"`        //图像列表;
	//FeatureList         *base.FeatureList  `json:"FeatureList,omitempty"`         //FeatureList;
	//RelatedList         *base.RelatedList  `json:"RelatedList,omitempty"`         //关联关系实体;
}

type PersonInfoOfVehicleTypeList struct {
	PersonInfoList []*PersonInfoOfVehicleType `json:"PersonInfoList"`
}

type PersonInfoOfVehicleType struct {
	IDNumber            string `json:"IDNumber"`
	Name                string `json:"Name"`
	UsedName            string `json:"UsedName"`
	CyclingPersonNumber int    `json:"CyclingPersonNumber"`
	PersonDirection     int    `json:"PersonDirection"`
	Things              int    `json:"Things"`
	HeightUpLimit       int    `json:"HeightUpLimit"`
	HeightLowerLimit    int    `json:"HeightLowerLimit"`
	TargetSize          int    `json:"TargetSize"`
	Direction           int    `json:"Direction"`
	Speed               int    `json:"Speed"`
	Backup              int    `json:"Backup"`
	IsBackup            int    `json:"IsBackup"`
	IsGlass             int    `json:"IsGlass"`
	IsCap               int    `json:"IsCap"`
}
