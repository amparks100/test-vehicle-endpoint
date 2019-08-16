package models

//VehicleRequest struct of the request coming into api
type VehicleRequest struct {
	Vin  string `json:"VIN"`
	Imei string `json:"IMEI,omitempty"`
}
