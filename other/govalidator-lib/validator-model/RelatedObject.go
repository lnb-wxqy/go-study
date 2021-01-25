package validator_model


type Related struct {
	RelatedType string `json:"RelatedType,omitempty"` // 关联类型

	RelatedID string `json:"RelatedID,omitempty"` //当关联多个信息时，用“,”号隔开
}

type RelatedList struct {
	RelatedObject []*Related `json:"RelatedObject,omitempty"`
}


