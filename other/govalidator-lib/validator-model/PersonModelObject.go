package validator_model

type PersonModel struct {
	PersonListObject *PersonListObject `json:"PersonListObject" validate:"required"`

	ProxyType         string `json:"proxyType,omitempty"`         //数据格式标识，0:GAT1400 格式，1:私有格式(扩展GAT1400)
	ProxyManufacturer string `json:"proxyManufacturer,omitempty"` //厂商编码
	ResJson           string `json:"resJson,omitempty"`           //扩展字段
}

type PersonListObject struct {
	PersonObject []*PersonObject `json:"PersonObject" validate:"required"`
}

type PersonObject struct {
	PersonID             string `json:"PersonID" validate:"required||string=48||regexp=21-22,02"` //人员标识 正则匹配：^\d{20}02\d{26}$
	PersonAppearTime     string `json:"PersonAppearTime,omitempty"`                               //人员出现时间
	PersonDisAppearTime  string `json:"PersonDisAppearTime,omitempty"`                            //人员消失时间
	HeightUpLimit        int32  `json:"HeightUpLimit,omitempty"`                                  //身高上限
	HeightLowerLimit     int32  `json:"HeightLowerLimit,omitempty"`                               //身高下限
	BodyType             string `json:"BodyType,omitempty"`                                       //体型
	Gesture              string `json:"Gesture,omitempty"`                                        //姿态
	Status               string `json:"Status,omitempty"`                                         //状态
	BodyFeature          string `json:"BodyFeature,omitempty"`                                    //体表特征
	HabitualMovement     string `json:"HabitualMovement,omitempty"`                               //习惯动作
	Behavior             string `json:"Behavior,omitempty"`                                       //行为
	BehaviorDescription  string `json:"BehaviorDescription,omitempty"`                            //行为描述
	Appendant            string `json:"Appendant,omitempty"`                                      //附属物
	AppendantDescription string `json:"AppendantDescription,omitempty"`                           //附属物描述

	UmbrellaColor string `json:"UmbrellaColor,omitempty"` //伞颜色
	ScarfColor    string `json:"ScarfColor,omitempty"`    //围巾颜色

	BagStyle string `json:"BagStyle,omitempty"` //包款式
	BagColor string `json:"BagColor,omitempty"` //包颜色

	CoatStyle  string `json:"CoatStyle,omitempty"`  //上衣款式
	CoatColor  string `json:"CoatColor,omitempty"`  //上衣颜色
	CoatLength string `json:"CoatLength,omitempty"` //上衣长度

	TrousersStyle string `json:"TrousersStyle,omitempty"` //裤子款式
	TrousersColor string `json:"TrousersColor,omitempty"` //裤子颜色
	TrousersLen   string `json:"TrousersLen,omitempty"`   //裤子长度

	ShoesStyle string `json:"ShoesStyle,omitempty"` //鞋子款式
	ShoesColor string `json:"ShoesColor,omitempty"` //鞋子颜色

	FaceID            int   `json:"FaceID,omitempty"`            //人脸标识
	NonMotorVehicleID int   `json:"NonMotorVehicleID,omitempty"` //非机动车标识
	PersonDirection   int32 `json:"PersonDirection,omitempty"`   //人体朝向属性
	Backpack          int32 `json:"Backpack,omitempty"`          //背包位置
	IsBackpack        int32 `json:"IsBackpack,omitempty"`        //是否背包
	IsGlass           int32 `json:"IsGlass,omitempty"`           //是否戴眼镜
	IsRespirator      int32 `json:"IsRespirator,omitempty"`      //是否戴口罩
	IsCap             int32 `json:"IsCap,omitempty"`             //是否戴帽子
	UpTexture         int32 `json:"UpTexture,omitempty"`         //上身纹理
	DownType          int32 `json:"DownType,omitempty"`          //下身类型
	MoveDirection     int32 `json:"MoveDirection,omitempty"`     //运动方向
	MoveSpeed         int32 `json:"MoveSpeed,omitempty"`         //运动速度
	HeadMarker        int32 `json:"HeadMarker,omitempty"`        //头部标识

	InfoKind                          int     `json:"InfoKind,omitempty" validate:"required||in=0,1"`     //信息分类 -备注：人工采集还是自动采集 R -
	SourceID                          string  `json:"SourceID,omitempty" validate:"required||string=41"`  //来源标识 -备注：来源图像标识 R -必选
	DeviceID                          string  `json:"DeviceID,omitempty" validate:"required||string=20"`  //设备编码
	LeftTopX                          int32   `json:"LeftTopX,omitempty" validate:"required"`             //左上角X坐标
	LeftTopY                          int32   `json:"LeftTopY,omitempty" validate:"required"`             //左上角Y坐标
	RightBtmX                         int32   `json:"RightBtmX,omitempty" validate:"required"`            //右下角X坐标
	RightBtmY                         int32   `json:"RightBtmY,omitempty" validate:"required"`            //右下角Y坐标
	LocationMarkTime                  string  `json:"LocationMarkTime,omitempty"`                         //位置标记时间
	IDType                            string  `json:"IDType,omitempty"`                                   //证件种类
	IDNumber                          string  `json:"IDNumber,omitempty"`                                 //证件号码
	Name                              string  `json:"Name,omitempty"`                                     //姓名
	UsedName                          string  `json:"UsedName,omitempty"`                                 // 曾用名
	Alias                             string  `json:"Alias,omitempty"`                                    //绰号
	GenderCode                        string  `json:"GenderCode,omitempty" validate:"in=0,1,2,9"`         //性别代码
	AgeUpLimit                        int32   `json:"AgeUpLimit,omitempty"`                               //年龄上限
	AgeLowerLimit                     int32   `json:"AgeLowerLimit,omitempty"`                            //年龄下限
	EthicCode                         string  `json:"EthicCode,omitempty"`                                //民族代码
	NationalityCode                   string  `json:"NationalityCode,omitempty"`                          //国籍代码
	NativeCityCode                    string  `json:"NativeCityCode,omitempty"`                           //籍贯省市县代码
	ResidenceAdminDivision            string  `json:"ResidenceAdminDivision,omitempty"`                   //居住地行政区划
	ChineseAccentCode                 string  `json:"ChineseAccentCode,omitempty"`                        //汉语口音代码
	PersonOrg                         string  `json:"PersonOrg,omitempty"`                                //单位名称
	JobCategory                       string  `json:"JobCategory,omitempty"`                              //职业类别代码
	AccompanyNumber                   int32   `json:"AccompanyNumber,omitempty"`                          //同行人数
	SkinColor                         string  `json:"SkinColor,omitempty"`                                //肤色
	HairStyle                         string  `json:"HairStyle,omitempty"`                                //发型
	HairType                          string  `json:"HairType,omitempty"`                                 //
	HairColor                         string  `json:"HairColor,omitempty"`                                //发色
	FaceStyle                         string  `json:"FaceStyle,omitempty"`                                //脸型
	FacialFeature                     string  `json:"FacialFeature,omitempty"`                            //脸部特征
	PhysicalFeature                   string  `json:"PhysicalFeature,omitempty"`                          //体貌特征
	RespiratorColor                   string  `json:"RespiratorColor,omitempty"`                          //口罩颜色
	CapStyle                          string  `json:"CapStyle,omitempty"`                                 //帽子款式
	CapColor                          string  `json:"CapColor,omitempty"`                                 //帽子颜色
	GlassStyle                        string  `json:"GlassStyle,omitempty"`                               //眼镜款式
	GlassColor                        string  `json:"GlassColor,omitempty"`                               //眼镜颜色
	IsDriver                          int32   `json:"IsDriver,omitempty" validate:"in=0,1,2"`             //是否驾驶员
	IsForeigner                       int32   `json:"IsForeigner,omitempty" validate:"in=0,1,2"`          //是否涉外人员
	PassportType                      string  `json:"PassportType,omitempty"`                             //护照证件种类
	ImmigrantTypeCode                 string  `json:"ImmigrantTypeCode,omitempty"`                        //出入境人员类别编码
	IsSuspectedTerrorist              int32   `json:"IsSuspectedTerrorist,omitempty" validate:"in=0,1,2"` //是否涉恐人员
	SuspectedTerroristNumber          string  `json:"SuspectedTerroristNumber,omitempty"`                 //涉恐人员编号
	IsCriminalInvolved                int32   `json:"IsCriminalInvolved,omitempty" validate:"in=0,1,2"`   //是否涉案人员
	IsSuspiciousPerson                int32   `json:"IsSuspiciousPerson,omitempty" validate:"in=0,1,2"`   //是否可疑人
	CriminalInvolvedSpecilisationCode string  `json:"CriminalInvolvedSpecilisationCode,omitempty"`        //涉案人员专长代码
	BodySpeciallMark                  string  `json:"BodySpeciallMark,omitempty"`                         //体表特殊标记
	CrimeMethod                       string  `json:"CrimeMethod,omitempty"`                              //作案手段
	CrimeCharacterCode                string  `json:"CrimeCharacterCode,omitempty"`                       //作案特点代码
	EscapedCriminalNumber             string  `json:"EscapedCriminalNumber,omitempty"`                    //在逃人员编号
	IsDetainees                       int32   `json:"IsDetainees,omitempty"`                              //是否在押人员
	DetentionHouseCode                string  `json:"DetentionHouseCode,omitempty"`                       //看守所编码
	DetaineesIdentity                 string  `json:"DetaineesIdentity,omitempty"`                        //在押人员身份
	DetaineesSpecialIdentity          string  `json:"DetaineesSpecialIdentity,omitempty"`                 //在押人员特殊身份
	MemberTypeCode                    string  `json:"MemberTypeCode,omitempty"`                           //成员类型代码
	IsVictim                          int32   `json:"IsVictim,omitempty"`                                 //是否被害人
	VictimType                        string  `json:"VictimType,omitempty"`                               //被害人种类
	InjuredDegree                     string  `json:"InjuredDegree,omitempty"`                            //受伤害程度
	CorpseConditionCode               string  `json:"CorpseConditionCode,omitempty"`                      //尸体状况代码
	IsSuspiciousFace                  int32   `json:"IsSuspiciousFace,omitempty"`                         //是否可疑人
	StorageURL                        string  `json:"StorageURL,omitempty"`                               //大图（场景图）路径
	TabID                             string  `json:"TabID,omitempty"`                                    //归属分类标签标识
	ResJson                           string  `json:"resJson,omitempty"`                                  //预留扩展字段
	RelatedType                       string  `json:"RelatedType,omitempty"`                              //关联关系类型【海康提供的标准】 01-人员 02-机动车 03-非机动车 04-物品 05-场景 06-人脸 07-视频图像标签 99-其他
	Longitude                         float64 `json:"Longitude,omitempty"`                                //设备经度【固定点位设备可选填，移动设备必填】
	Latitude                          float64 `json:"Latitude,omitempty"`                                 //设备纬度【固定点位设备可选填，移动设备必填】
	//
	//SubImageList *base.SubImageList `json:"SubImageList,omitempty"` //图像列表
	//FeatureList  *base.FeatureList  `json:"FeatureList,omitempty"`  //特征值列表
	//RelatedList  *base.RelatedList  `json:"RelatedList,omitempty"`  //关联关系实体
}
