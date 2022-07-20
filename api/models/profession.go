package models

type CreateProfession struct {
	Name string `json:"name" binding:"required"`
}

type Profession struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
}

type GetAllProfessionResponse struct {
	Professions []Profession `json:"professions"`
	Count       int32        `json:"count"`
}
