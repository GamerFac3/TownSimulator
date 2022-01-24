package models

type Job struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PersonID    uint   `json:"person_id"`
	BuildingID  uint   `json:"building_id"`
}

type CreateJobInput struct {
	Name        string `json:"name" bindings:"required"`
	Description string `json:"description" bindings:"required"`
	PersonID    uint   `json:"person_id"`
	BuildingID  uint   `json:"building_id" bindings:"required"`
}

type UpdateJobInput struct {
	Name        string `json:"name" `
	Description string `json:"description" `
	PersonID    uint   `json:"person_id"`
	BuildingID  uint   `json:"building_id" `
}
