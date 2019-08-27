package models

//VehicleMessageModel model used for messages sent to SNS/SQS queue
type VehicleMessageModel struct {
	MessageType string `json:"MessageType"`
	Vin         string `json:"vin"`
	Imei        string `json:"imei"`
}
