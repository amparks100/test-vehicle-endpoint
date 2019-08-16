# Demo vehicle endpoint for AWS

What this go project does:

POST
- receives a json of vehicle data
- prints the vin and imei out in response
- saves to RDS PostgreSQL database

GET
- retrieves from the RDS PostgreSQL database
  

## How to put this into a lambda

`go get github.com/aws/aws-lambda-go/lambda`

`$env:GOOS = "linux"`

`go build -o test-vehicle-endpoint test-vehicle-endpoint`

`build-lambda-zip -o test-vehicle-endpoint.zip test-vehicle-endpoint`

upload the outputted zip file into the lambda function. set the handler to `test-vehicle-endpoint`


## References:

how to package:
https://docs.aws.amazon.com/lambda/latest/dg/lambda-go-how-to-create-deployment-package.html

database:
https://medium.com/@yunskilic/developing-golang-aws-lambda-functions-with-gorm-amazon-rds-for-postgresql-c5efbcbd0b0d
