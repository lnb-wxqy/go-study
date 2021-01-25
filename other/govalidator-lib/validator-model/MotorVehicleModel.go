package validator_model

type MotorVehicleModel struct {
	MotorVehicleListObject *MotorVehicleListObject `json:"MotorVehicleListObject" validate:"required"`

	ProxyType         string `json:"proxyType,omitempty"`         //数据格式标识，0:GAT1400 格式，1:私有格式(扩展GAT1400)
	ProxyManufacturer string `json:"proxyManufacturer,omitempty"` //厂商编码
	ResJson           string `json:"resJson,omitempty"`           //扩展字段
}

type MotorVehicleListObject struct {
	MotorVehicleObject []*MotorVehicleObject `json:"MotorVehicleObject" validate:"required"`
}

type MotorVehicleObject struct {
	MotorVehicleID    string `json:"MotorVehicleID" validate:"required||string=48||regexp=21-22,02"` //车辆标识 -备注：辆全局唯一标识 R -必选 正则匹配：^\d{20}02\d{26}$
	TollgateID        string `json:"TollgateID,omitempty"`                                           //关联卡口编码 O -可选
	StorageURL1       string `json:"StorageUrl1"  validate:"required||url"`                                                    //近景照片 -卡口相机所拍照片，自动采集必选，图像访问路径采用URI命名规范 R -必选
	StorageURL2       string `json:"StorageUrl2,omitempty"`                                          //车辆照片 O -可选
	StorageURL3       string `json:"StorageUrl3,omitempty"`                                          //远景照片 全景相机所拍照片 O 可选
	StorageURL4       string `json:"StorageUrl4,omitempty"`                                          //合成图 O 可选
	StorageURL5       string `json:"StorageUrl5,omitempty"`                                          //缩略图 O 可选
	LaneNo            int    `json:"LaneNo,omitempty"`                                               //车道号 O 可选
	Direction         string `json:"Direction,omitempty"`                                            //行驶方向
	VehicleClass      string `json:"VehicleClass,omitempty"`                                         //车辆类型
	VehicleColorDepth string `json:"VehicleColorDepth,omitempty"`                                    //颜色深浅
	HitMarkInfo       string `json:"HitMarkInfo,omitempty"`                                          //撞痕信息
	VehicleBodyDesc   string `json:"VehicleBodyDesc,omitempty"`                                      //车身描述
	VehicleFrontItem  string `json:"VehicleFrontItem,omitempty"`                                     //车前部物品
	VehicleRearItem   string `json:"VehicleRearItem,omitempty"`                                      //车前部物品描述
	DescOfRearItem    string `json:"DescOfRearItem,omitempty"`                                       //车后部物品
	NumOfPassenger    int32  `json:"NumOfPassenger,omitempty"`                                       //车内人数
	PassTime          string `json:"PassTime,omitempty" validate:"required||datetime"`               //经过时间
	NameOfPassedRoad  string `json:"NameOfPassedRoad,omitempty"`                                     //经过道路名称
	//IsSuspicious   string `json:"IsSuspicious,omitempty"`         //是否可疑车
	Sunvisor             int32  `json:"Sunvisor,omitempty"`             //遮阳板状态
	SafetyBelt           int32  `json:"SafetyBelt,omitempty"`           //安全带状态
	Calling              int32  `json:"Calling,omitempty"`              //打电话状态
	PlateReliability     string `json:"PlateReliability,omitempty"`     //号牌识别可信度
	PlateCharReliability string `json:"PlateCharReliability,omitempty"` //每位号牌号码可信度
	BrandReliability     string `json:"BrandReliability,omitempty"`     //品牌标识识别可信度
	DriverFaceID         string `json:"DriverFaceID,omitempty"`         //主驾人脸标识
	CopilotFaceID        string `json:"CopilotFaceID,omitempty"`        //副驾人脸标识
	VehicleStyles        string `json:"VehicleStyles,omitempty"`        //车辆年款
	//IsModified           string `json:"IsModified,omitempty"`           //改装标志
	AnnualInspectionNum int32 `json:"AnnualInspectionNum,omitempty"` //车辆年检标数
	MarkerType          int32 `json:"MarkerType,omitempty"`          //标志物
	VehicleHeadend      int32 `json:"VehicleHeadend,omitempty"`      //车头车尾
	ShieldFace          int32 `json:"ShieldFace,omitempty"`          //是否遮挡面部
	PendantsNum         int32 `json:"PendantsNum,omitempty"`         //挂件个数
	OrnamentsNum        int32 `json:"OrnamentsNum,omitempty"`        //摆件个数
	TissueBoxNum        int32 `json:"TissueBoxNum,omitempty"`        //纸巾盒数

	SpecialCar int32 `json:"SpecialCar,omitempty"` //特殊车辆 平谷添加

	DescOfFrontItem string `json:"DescOfFrontItem,omitempty"`
	ResJSON         string `json:"resJson,omitempty"`

	InfoKind      int    `json:"InfoKind,omitempty" validate:"required||in=0,1"`    //信息分类 -备注：人工采集还是自动采集 R -
	SourceID      string `json:"SourceID,omitempty" validate:"required||string=41"` //来源标识 -备注：来源图像标识 R -必选
	DeviceID      string `json:"DeviceID,omitempty" validate:"required||string=20"` //设备编码 R/O -自动采集时必选
	LeftTopX      int32  `json:"LeftTopX,omitempty" validate:"required"`            //左上角X坐标 车的轮廓外界矩形在画面中的位置，记录坐标 R/O 自动采集时必选
	LeftTopY      int32  `json:"LeftTopY,omitempty" validate:"required"`            //左上角Y坐标 车的轮廓外界矩形在画面中的位置，记录坐标 R/O 自动采集时必选
	RightBtmX     int32  `json:"RightBtmX,omitempty" validate:"required"`           //右上角X坐标 车的轮廓外界矩形在画面中的位置，记录坐标 R/O 自动采集时必选
	RightBtmY     int32  `json:"RightBtmY,omitempty" validate:"required"`           //右上角Y坐标 车的轮廓外界矩形在画面中的位置，记录坐标 R/O 自动采集时必选;
	MarkTime      string `json:"MarkTime,omitempty"`                                //位置标记时间 O 人工采集时有效;
	AppearTime    string `json:"AppearTime,omitempty"`                              //车辆出现时间 O 人工采集时有效
	DisAppearTime string `json:"DisAppearTime,omitempty"`                           //车辆消失时间 O 人工采集时有效;
	HasPlate      string `json:"HasPlate,omitempty"`                                //有无车牌号;
	PlateClass    string `json:"PlateClass,omitempty"`                              //号牌种类;
	PlateColor    string `json:"PlateColor,omitempty"`                              //车牌颜色;
	PlateNo       string `json:"PlateNo,omitempty"`                                 //车牌号;
	PlateNoAttach string `json:"PlateNoAttach,omitempty"`                           //挂车牌号;
	PlateDescribe string `json:"PlateDescribe,omitempty"`                           //车牌描述;
	//IsDecked            string             `json:"IsDecked,omitempty"`            //是否套牌;
	//IsAltered    string             `json:"IsAltered,omitempty"`           //是否涂改;
	//IsCovered    string             `json:"IsCovered,omitempty"`           //是否遮挡;
	Speed               float64 `json:"Speed,omitempty"`                               //行驶速度;
	DrivingStatusCode   string  `json:"DrivingStatusCode,omitempty"`                   //行驶状态代码;
	UsingPropertiesCode int     `json:"UsingPropertiesCode,omitempty"`                 //车辆使用性质代码;
	VehicleBrand        string  `json:"VehicleBrand,omitempty"`                        //车辆品牌;
	VehicleModel        string  `json:"VehicleModel,omitempty"`                        //车辆型号;
	VehicleLength       int     `json:"VehicleLength,omitempty"`                       //车辆长度;
	VehicleWidth        int     `json:"VehicleWidth,omitempty"`                        //车辆宽度;
	VehicleHeight       int     `json:"VehicleHeight,omitempty"`                       //车辆高度;
	VehicleColor        string  `json:"VehicleColor,omitempty" validate:"required"`    //车身颜色 R 必选;
	VehicleHood         string  `json:"VehicleHood,omitempty"`                         //车前盖;
	VehicleTrunk        string  `json:"VehicleTrunk,omitempty"`                        //车后盖;
	VehicleWheel        string  `json:"VehicleWheel,omitempty"`                        //车轮;
	WheelPrintedPattern string  `json:"WheelPrintedPattern,omitempty"`                 //车轮印花纹;
	VehicleWindow       string  `json:"VehicleWindow,omitempty"`                       //车窗;
	VehicleRoof         string  `json:"VehicleRoof,omitempty"`                         //车顶;
	VehicleDoor         string  `json:"VehicleDoor,omitempty"`                         //车门;
	SideOfVehicle       string  `json:"SideOfVehicle,omitempty"`                       //车侧;
	CarOfVehicle        string  `json:"CarOfVehicle,omitempty"`                        //车厢;
	RearviewMirror      string  `json:"RearviewMirror,omitempty"`                      //后视镜;
	VehicleChassis      string  `json:"VehicleChassis,omitempty"`                      //底盘;
	VehicleShielding    string  `json:"VehicleShielding,omitempty"`                    //遮挡;
	FilmColor           string  `json:"FilmColor,omitempty"`                           //贴膜颜色;
	StorageURL          string  `json:"StorageURL,omitempty"` //大图（场景图）路径;
	NationalityCode     string  `json:"NationalityCode,omitempty"`                     //NationalityCode;
	TabID               string  `json:"TabID,omitempty"`                               //TabID;
	RelatedType         string  `json:"RelatedType,omitempty"`                         //关联关系类型【海康提供的标准】;
	Longitude           float64 `json:"Longitude,omitempty"`                           //设备经度【固定点位设备可选填，移动设备必填】;
	Latitude            float64 `json:"Latitude,omitempty"`                            //设备纬度【固定点位设备可选填，移动设备必填】
	resJson             string  `json:"resJson,omitempty"`                             //预留扩展字段
	//SubImageList        *base.SubImageList `json:"SubImageList,omitempty"`        //图像列表;
	//FeatureList         *base.FeatureList  `json:"FeatureList,omitempty"`         //FeatureList;
	//RelatedList         *base.RelatedList  `json:"RelatedList,omitempty"`         //关联关系实体;
}
