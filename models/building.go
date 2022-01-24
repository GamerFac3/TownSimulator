package models

type Building struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	BuildingTypeID uint   `json:"building_type_id"`
	Jobs           []Job  `json:"jobs"`
}
