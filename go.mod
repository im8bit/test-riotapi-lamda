module test-riotapi-lambda/main

go 1.17

require (
	github.com/aws/aws-lambda-go v1.28.0
	github.com/aws/aws-sdk-go v1.43.15
	github.com/im8bit/test-riotapi-library v0.0.0-00010101000000-000000000000
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

replace github.com/im8bit/test-riotapi-library => ../test-riotapi-library
