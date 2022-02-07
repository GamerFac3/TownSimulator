package models

type JobName struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	BuildingTypeID uint   `json:"building_type_id"`
	Priority       int    `json:"priority"`
}
type CreateJobNameInput struct {
	Name           string `json:"name" bindings:"required"`
	Priority       int    `json:"priority" bindings:"required"`
	BuildingTypeID uint   `json:"building_type_id" bindings:"required"`
}

type UpdateJobNameInput struct {
	Name           string `json:"name"`
	BuildingTypeID uint   `json:"building_type_id"`
	Priority       int    `json:"priority"`
}
