package models

type CreateAttributeModel struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type AttributeModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"attribute_type"`
}

type GetAllAttributeModel struct {
	Attributes []AttributeModel `json:"attributes"`
	Count      int32            `json:"count"`
}
