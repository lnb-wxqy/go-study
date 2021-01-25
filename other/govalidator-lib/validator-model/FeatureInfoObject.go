package validator_model

type FeatureInfo struct {
	Vendor           string `json:"Vendor,omitempty"`           //.厂家
	AlgorithmVersion string `json:"AlgorithmVersion,omitempty"` //算法版本
	FeatureData      string `json:"FeatureData,omitempty"`      //特征值数据
}

type FeatureList struct {
	FeatureInfoObject []*FeatureInfo `json:"FeatureInfoObject,omitempty"`
}

func Build1400FeatureList(vendor ,algorithmVersion ,featureData string,featureList *FeatureList)  *FeatureList {
	featureList.FeatureInfoObject = append(featureList.FeatureInfoObject,&FeatureInfo{
		Vendor: vendor,
		AlgorithmVersion:algorithmVersion,
		FeatureData:featureData,
	} )
	return featureList
}