package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//VehicleDataModel data model used for postgreSQL
type VehicleDataModel struct {
	gorm.Model
	Vin  string `json:"vin"`
	Imei string `json:"imei,omitempty"`
}
