package models

type Building struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	Name           string `json:"name"`
	BuildingTypeID uint   `json:"building_type_id"`
	Jobs           []Job  `json:"jobs"`
}

type CreateBuildingInput struct {
	Name           string `json:"name" bindings:"required"`
	BuildingTypeID uint   `json:"building_type_id" bindings:"required"`
	Jobs           []Job  `json:"jobs"`
}

type UpdateBuildingInput struct {
	Name           string `json:"name"`
	BuildingTypeID uint   `json:"building_type_id"`
	Jobs           []Job  `json:"jobs"`
}
