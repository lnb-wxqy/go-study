package govalidator_lib

import (
	"errors"
	"fmt"
	"github.com/smokezl/govalidators"
	validator_model "golang-study/other/govalidator-lib/validator-model"
	"reflect"
	"regexp"
	"strings"
)

const (
	NotificationIDRegexStr = `^\d{12}04\d{19}$`
	SubscribeIDRegexStr    = `^\d{12}03\d{19}$`
	DataIDRegexStr         = `^\d{20}02\d{26}$`

	NOTIFICATIONID    = "NotificationID"
	SUBSCRIBEID       = "SubscribeID"
	FACEID            = "FaceID"
	PERSONID          = "PersonID"
	MOTORVEHICLEID    = "MotorVehicleID"
	NONMOTORVEHICLEID = "NonMotorVehicleID"
)

type RegexpValidator struct {
	EMsg string
	//govalidators.Range
}

func (rv *RegexpValidator) Validate(params map[string]interface{}, val reflect.Value, args ...string) (bool, error) {

	eMsg := "[name]的第[args0]位必须是[args1]"
	eParamsMap := map[string]string{
		"name": params["name"].(string),
		"args0": args[0],
		"args1": args[1],
	}
	if rv.EMsg != "" {
		eMsg = rv.EMsg
	}

	v := val.String()
	var b bool
	switch eParamsMap["name"] {
	case NOTIFICATIONID:
		b, _ = regexp.MatchString(NotificationIDRegexStr, v)
	case SUBSCRIBEID:
		b, _ = regexp.MatchString(SubscribeIDRegexStr, v)
	case FACEID, PERSONID, MOTORVEHICLEID, NONMOTORVEHICLEID:
		b, _ = regexp.MatchString(DataIDRegexStr, v)
	}
	if !b {
		return false, formatError(eMsg, eParamsMap)
	}
	return true, nil
}

func formatError(format string, eParamsMap map[string]string) error {
	var params []string
	for k, v := range eParamsMap {
		params = append(params, "["+k+"]", v)
	}
	replacer := strings.NewReplacer(params...)
	return errors.New(replacer.Replace(format))
}

type Class struct {
	Cid       int64  `validate:"required||integer=10000,_"`
	Cname     string `validate:"required||string=1,5||unique"`
	BeginTime string `validate:"required||datetime=H:i"`
}

type Student struct {
	Uid string `validate:"string=1,15"`

	//Uid          string    `validate:"required||integer=10000,_"`
	Name         string   `validate:"required||string=1,15"`
	Age          int64    `validate:"required||integer=10,30"`
	Sex          string   `validate:"required||in=male,female"`
	Email        string   `validate:"email||user||vm"`
	PersonalPage string   `validate:"url"`
	Hobby        []string `validate:"array=_,2||unique||in=swimming,running,drawing"`
	CreateTime   string   `validate:"datetime"`
	Class        []Class  `validate:"array=1,3"`
}

func ValidatorStudent() {

	validator := govalidators.New()
	student := &Student{
		Uid:          "33333333333333333333333333333333",
		Name:         "",
		Age:          31,
		Sex:          "male1",
		Email:        "@qq.com",
		PersonalPage: "www.abcd.com",
		Hobby:        []string{"swimming", "singing"},
		CreateTime:   "2018-03-03 05:60:00",
		Class: []Class{
			Class{
				Cid:       12345678,
				Cname:     "语文",
				BeginTime: "13:00",
			},
			Class{
				Cid:       22345678,
				Cname:     "数学",
				BeginTime: "13:00",
			},
			Class{
				Cid:       32345678,
				Cname:     "数学",
				BeginTime: "13:60",
			},
		},
	}
	validatorMap := make(map[string]interface{}, 0)
	validatorMap = map[string]interface{}{
		"regexp": &RegexpValidator{
			EMsg: "[name]嘻嘻嘻",
		},
		//"vm": validationMethod,
		"string": &govalidators.StringValidator{
			EMsg: "[name] 长度必须为[args]",
			Range: govalidators.Range{
				RangeEMsg: map[string]string{
					"between":  "[name] 长度必须在 [min] 和 [max] 之间",
					"lessThan": "",
					"equal":    "",
					"atLeast":  "",
				},
			},
		},
		"integer": &govalidators.IntegerValidator{
			Range: govalidators.Range{
				RangeEMsg: map[string]string{
					"between": "[name] 的值必须在 [min] 和 [max] 之间",
				},
			},
		},
		"in": &govalidators.InValidator{
			EMsg:     "[name] 的值必须为 [args] 中的一个",
			TypeEMsg: "[name]的[args]",
		},
		"email": &govalidators.EmailValidator{
			EMsg: "[name] 不是一个有效的email地址",
		},
		"url": &govalidators.UrlValidator{
			EMsg: "[name] 不是一个有效的url地址",
		},
		"datetime": &govalidators.DateTimeValidator{
			EMsg: "[name] 不是一个有效的日期",
		},
		"unique": &govalidators.UniqueValidator{
			EMsg: "[name] 不是唯一的",
		},
	}
	validator.SetValidators(validatorMap)
	errList := validator.Validate(student)
	if errList != nil {
		for _, err := range errList {
			fmt.Println("err:", err)
		}
	}

}

func Validator() {

	var faceListObject validator_model.FaceListObject
	var faceObject = make([]*validator_model.FaceObject, 0)
	face := &validator_model.FaceObject{
		FaceID:                            "123",
		FaceAppearTime:                    "",
		FaceDisAppearTime:                 "",
		Attitude:                          "",
		Similaritydegree:                  0,
		EyebrowStyle:                      "",
		NoseStyle:                         "",
		MustacheStyle:                     "",
		LipStyle:                          "",
		WrinklePouch:                      "",
		AcneStain:                         "",
		FreckleBirthmark:                  "",
		ScarDimple:                        "",
		OtherFeature:                      "",
		PersonID:                          "",
		MotorVehicleID:                    "",
		NonMotorVehicleID:                 "",
		ShotTime:                          "",
		Emotion:                           0,
		IsGlass:                           0,
		IsRespirator:                      0,
		IsCap:                             0,
		HeadMarker:                        0,
		MoveDirection:                     0,
		MoveSpeed:                         0,
		InfoKind:                          0,
		SourceID:                          "",
		DeviceID:                          "",
		LeftTopX:                          0,
		LeftTopY:                          -1,
		RightBtmX:                         0,
		RightBtmY:                         0,
		LocationMarkTime:                  "",
		IDType:                            "",
		IDNumber:                          "",
		Name:                              "",
		UsedName:                          "",
		Alias:                             "",
		GenderCode:                        "",
		AgeUpLimit:                        0,
		AgeLowerLimit:                     0,
		EthicCode:                         "",
		NationalityCode:                   "",
		NativeCityCode:                    "",
		ResidenceAdminDivision:            "",
		ChineseAccentCode:                 "",
		PersonOrg:                         "",
		JobCategory:                       "",
		AccompanyNumber:                   0,
		SkinColor:                         "",
		HairStyle:                         "",
		HairType:                          "",
		HairColor:                         "",
		FaceStyle:                         "",
		FacialFeature:                     "",
		PhysicalFeature:                   "",
		RespiratorColor:                   "",
		CapStyle:                          "",
		CapColor:                          "",
		GlassStyle:                        "",
		GlassColor:                        "",
		IsDriver:                          0,
		IsForeigner:                       0,
		PassportType:                      "",
		ImmigrantTypeCode:                 "",
		IsSuspectedTerrorist:              0,
		SuspectedTerroristNumber:          "",
		IsCriminalInvolved:                0,
		IsSuspiciousPerson:                0,
		CriminalInvolvedSpecilisationCode: "",
		BodySpeciallMark:                  "",
		CrimeMethod:                       "",
		CrimeCharacterCode:                "",
		EscapedCriminalNumber:             "",
		IsDetainees:                       0,
		DetentionHouseCode:                "",
		DetaineesIdentity:                 "",
		DetaineesSpecialIdentity:          "",
		MemberTypeCode:                    "",
		IsVictim:                          0,
		VictimType:                        "",
		InjuredDegree:                     "",
		CorpseConditionCode:               "",
		IsSuspiciousFace:                  0,
		StorageURL:                        "",
		TabID:                             "",
		ResJson:                           "",
		RelatedType:                       "",
		Longitude:                         0,
		Latitude:                          0,
	}
	faceObject = append(faceObject, face)
	faceListObject.FaceObject = faceObject

	var subscribeNotificationModel validator_model.SubscribeNotificationModel
	var subscribeNotificationListObject validator_model.SubscribeNotificationListObject
	var subscribeNotificationObject = make([]*validator_model.SubscribeNotification, 0)
	subscribeNotification := &validator_model.SubscribeNotification{
		NotificationID:            "310116120013040400333101161200131",
		SubscribeID:               "310116120013040400333101161200131",
		Title:                     "1111111111111111111111111111111111111111111111111111111111111",
		TriggerTime:               "20201011123225",
		InfoIDs:                   "222222222222222",
		CaseObjectList:            nil,
		Tollgate:                  nil,
		LaneList:                  nil,
		DeviceList:                nil,
		DeviceStatusList:          nil,
		APSObjectList:             nil,
		APSStatusObjectList:       nil,
		PersonObjectList:          nil,
		FaceObjectList:            &faceListObject,
		MotorVehicleObjectList:    nil,
		NonMotorVehicleObjectList: nil,
		ThingObjectList:           nil,
		SceneObjectList:           nil,
	}

	subscribeNotificationObject = append(subscribeNotificationObject, subscribeNotification)
	subscribeNotificationListObject.SubscribeNotificationObject = subscribeNotificationObject
	subscribeNotificationModel.SubscribeNotificationListObject = &subscribeNotificationListObject

	validator := govalidators.New()
	validatorMap := make(map[string]interface{}, 0)
	validatorMap = map[string]interface{}{
		"regexp": &RegexpValidator{
			EMsg: "[name]的第[args0]位必须是[args1]",
		},
		//"regexp2": validationMethod,
		"required": &govalidators.RequiredValidator{
			EMsg: "[name] 不能为空",
		},
		"string": &govalidators.StringValidator{
			Range: govalidators.Range{
				RangeEMsg: map[string]string{
					//lessThan,equal,atLeast,between
					"equal":    "[name] 的长度必须是[min]",
					"between":  "[name] 的长度必须在 [min] 和 [max] 之间",
					"lessThan": "[name] 的长度不能大于[max]",
					"atLeast":  "[name] 的长度不能小于[min]",
				},
			},
		},
		"in": &govalidators.InValidator{
			EMsg:     "[name] 的值必须为 [args] 中的一个",
			TypeEMsg: "[name] 的值必须为 [args] 中的一个个",
		},
		"email": &govalidators.EmailValidator{
			EMsg: "[name] 不是一个有效的email地址",
		},
		"url": &govalidators.UrlValidator{
			EMsg: "[name] 不是一个有效的url地址",
		},
		"datetime": &govalidators.DateTimeValidator{
			FmtStr: "YmdHis", //Y-m-d H:i:s
			EMsg:   "[name] 不是一个有效的日期",
		},
	}
	validator.SetValidators(validatorMap)
	//errList := validator.Validate(subscribeNotificationModel)
	errList := validator.Validate(face)
	//errList := append(errs, errs2...)
	if errList != nil {
		for _, err := range errList {
			fmt.Println("err: ", err.Error())
		}
	}
}
