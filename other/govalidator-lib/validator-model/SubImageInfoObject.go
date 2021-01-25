package validator_model

type SubImageInfo struct {
	ImageID string `json:"ImageID,omitempty"`

	InfoKind int `json:"InfoKind,omitempty"`

	ImageSource string `json:"ImageSource,omitempty"`

	SourceVideoID string `json:"SourceVideoID,omitempty"`

	OriginImageID string `json:"OriginImageID,omitempty"`

	EventSort int32 `json:"EventSort,omitempty"`

	DeviceID string `json:"DeviceID,omitempty"`

	StoragePath string `json:"StoragePath,omitempty"`

	FileHash string `json:"FileHash,omitempty"`

	FileFormat string `json:"FileFormat,omitempty" validate:"required"`

	ShotTime string `json:"ShotTime,omitempty"`

	Title string `json:"Title,omitempty"`

	TitleNote string `json:"TitleNote,omitempty"`

	SpecialIName string `json:"SpecialIName,omitempty"`

	Keyword string `json:"Keyword,omitempty"`

	ContentDescription string `json:"ContentDescription,omitempty" validate:"required"`

	SubjectCharacter string `json:"SubjectCharacter,omitempty"`

	ShotPlaceCode string `json:"ShotPlaceCode,omitempty"`

	ShotPlaceFullAdress string `json:"ShotPlaceFullAdress,omitempty" validate:"required"`

	ShotPlaceLongitude string `json:"ShotPlaceLongitude,omitempty"`

	ShotPlaceLatitude string `json:"ShotPlaceLatitude,omitempty"`

	HorizontalShotDirection string `json:"HorizontalShotDirection,omitempty"`

	VerticalShotDirection string `json:"VerticalShotDirection,omitempty"`

	SecurityLevel string `json:"SecurityLevel,omitempty" validate:"required"`

	Width int32 `json:"Width"`

	Height int32 `json:"Height"`

	CameraManufacturer string `json:"CameraManufacturer,omitempty"`

	CameraVersion string `json:"CameraVersion,omitempty"`

	ApertureValue int `json:"ApertureValue,omitempty"`

	ISOSensitivity int `json:"ISOSensitivity,omitempty"`

	FocalLength int `json:"FocalLength,omitempty"`

	QualityGrade string `json:"QualityGrade,omitempty"`

	CollectorName string `json:"CollectorName,omitempty"`

	CollectorOrg string `json:"CollectorOrg,omitempty"`

	CollectorIDType string `json:"CollectorIDType,omitempty"`

	CollectorID string `json:"CollectorID,omitempty"`

	EntryClerk string `json:"EntryClerk,omitempty"`

	EntryClerkOrg string `json:"EntryClerkOrg,omitempty"`

	EntryClerkIDType string `json:"EntryClerkIDType,omitempty"`

	EntryClerkID string `json:"EntryClerkID,omitempty"`

	EntryTime string `json:"EntryTime,omitempty"`

	ImageProcFlag int `json:"ImageProcFlag,omitempty"`

	FileSize int `json:"FileSize,omitempty"`

	Data string `json:"Data,omitempty"`

	Type string `json:"Type,omitempty" validate:"required"`
}

type SubImageList struct {
	SubImageInfoObject []*SubImageInfo `json:"SubImageInfoObject"`
}

func (subImageList *SubImageList) BuildSubImageList(subImageInfo SubImageInfo) {
	subImageList.SubImageInfoObject = append(subImageList.SubImageInfoObject, &subImageInfo)
}


