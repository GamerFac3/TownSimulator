package models

type BuildingType struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	JobNames  []JobName  `json:"job_names"`
	Buildings []Building `json:"buildings" `
}

type CreateBuildingTypeInput struct {
	Name      string     `json:"name" bindings:"required"`
	JobNames  []JobName  `json:"job_names"`
	Buildings []Building `json:"buildings" `
}

type UpdateBuildingTypeInput struct {
	Name      string     `json:"name"`
	JobNames  []JobName  `json:"job_names"`
	Buildings []Building `json:"buildings" `
}
