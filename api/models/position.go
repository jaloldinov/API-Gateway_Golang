package models

type PositionAttribute struct {
	AttributeId string `json:"attribute_id"`
	Value       string `json:"value"`
}

type CreatePositionModel struct {
	Name               string               `json:"name"`
	ProfessionId       string               `json:"profession_id"`
	CompanyId          string               `json:"company_id"`
	PositionAttributes []*PositionAttribute `json:"position_attributes"`
}

type PositionAttributes struct {
	AttributeID string           `json:"attribute_id" `
	Value       string           `json:"value"`
	AttriModel  []AttributeModel `json:"attribute"`
}

type PositionModel struct {
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	ProfessionId  string               `json:"profession_id"`
	CompanyId     string               `json:"company_id"`
	PosAttributes []PositionAttributes `json:"position_attributes"`
}

type GetAllPositionModel struct {
	Positions []PositionModel `json:"positions"`
	Count     int32           `json:"count"`
}
