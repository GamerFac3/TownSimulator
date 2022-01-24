package models

type BuildingType struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	JobNames  []JobName  `json:"job_names"`
	Buildings []Building `json:"buildings"`
}

type JobName struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	BuildingTypeID uint   `json:"building_type_id"`
}
