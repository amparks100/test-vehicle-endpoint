# Demo vehicle endpoint for AWS

What this go project does:

POST
- receives a json of vehicle data
- prints the vin and imei out in response
- saves to RDS PostgreSQL database
- publishes message to SNS queue

GET
- retrieves from the RDS PostgreSQL database
  

## How to put this into a lambda

`go get github.com/aws/aws-lambda-go/lambda`

`$env:GOOS = "linux"`

`go build -o test-vehicle-endpoint test-vehicle-endpoint`

`build-lambda-zip -o test-vehicle-endpoint.zip test-vehicle-endpoint`

upload the outputted zip file into the lambda function. set the handler to `test-vehicle-endpoint`

## Amazon setup:

API Gateway -> Lamda -> RDS & SNS -> SQS

To connect to RDS - Need a VPC and a security group that connects to the lambda subnets

To connect to SNS - Need an Endpoint (found inside VPC dashboard) which connects to the lambda subnets


## References:

how to package:
https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html

database:
https://medium.com/@yunskilic/developing-golang-aws-lambda-functions-with-gorm-amazon-rds-for-postgresql-c5efbcbd0b0d

how to publish to SNS:
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sns-example-publish.html
https://sourcegraph.com/github.com/datacratic/aws-sdk-go/-/blob/service/sns/examples_test.go
