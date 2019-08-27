package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"

	"github.com/amparks100/test-vehicle-endpoint/db"
	"github.com/amparks100/test-vehicle-endpoint/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	lambda.Start(HandleRequest)
}

//HandleRequest API Gateway request handler
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Request Method: %s", request.HTTPMethod)
	log.Printf("Request body: %s", request.Body)
	var response events.APIGatewayProxyResponse

	postgresConnector := db.PostgresConnector{}
	database, err := postgresConnector.GetConnection()
	defer database.Close()

	if err != nil {
		response.StatusCode = 500
		response.Body = fmt.Sprintf("Could not open DB")
		return response, err
	}

	if request.HTTPMethod == "POST" {
		response, err = PostVehicle(request, database)
	}
	if request.HTTPMethod == "GET" {
		response, err = GetVehicle(request, database)
	}
	return response, err
}

//PostVehicle - handles post requests and saves to DB
func PostVehicle(request events.APIGatewayProxyRequest, database *gorm.DB) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse
	vehicle := models.VehicleRequest{}
	err := json.Unmarshal([]byte(request.Body), &vehicle)

	if err != nil {
		log.Fatal(err)
		response.StatusCode = 400
		response.Body = fmt.Sprintf("Could not unmarshal request: %s", request.Body)
		return response, err
	}
	database.AutoMigrate(&models.VehicleDataModel{})
	vehicleData := &models.VehicleDataModel{}
	vehicleData.Vin = vehicle.Vin
	vehicleData.Imei = vehicle.Imei
	database.Create(vehicleData)
	log.Printf("Saved Vehicle VIN: %s to database", vehicleData.Vin)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		response.StatusCode = 500
		response.Body = fmt.Sprintf("Unable to create session")
		return response, err
	}
	log.Printf("Created Session")
	client := sns.New(sess)
	log.Printf("Created client")

	message := fmt.Sprintf("%v", request.Body)

	pubReq, resp := client.PublishRequest(&sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String("arn:aws:sns:us-east-1:*:vehicle-test-queue"),
	})
	err = pubReq.Send()
	if err != nil {
		response.StatusCode = 500
		response.Body = fmt.Sprintf("Unable to publish message, %v", resp)
		return response, err
	}
	log.Printf("Published message to SNS, messageId: %v", resp.MessageId)

	response.StatusCode = 200
	response.Body = fmt.Sprintf("HALLO vehicle! VIN: %s, IMEI: %s", vehicle.Vin, vehicle.Imei)
	return response, nil
}

//GetVehicle - handles get requests and retrieves from DB
func GetVehicle(request events.APIGatewayProxyRequest, database *gorm.DB) (events.APIGatewayProxyResponse, error) {
	var response events.APIGatewayProxyResponse

	queryVin := request.QueryStringParameters["VIN"]
	log.Printf("Received query parameter for vin: %s", queryVin)

	var vehicle models.VehicleDataModel
	database.Where("vin = ?", queryVin).First(&vehicle)

	if len(vehicle.Vin) != 0 {
		response.StatusCode = 200
		response.Body = fmt.Sprintf("Found vehicle! VIN: %s, IMEI: %s", vehicle.Vin, vehicle.Imei)
	} else {
		response.StatusCode = 404
	}
	return response, nil
}
