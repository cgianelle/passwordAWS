# passwordAWS
AWS Lambda function to generate passwords

## Build/Deploy for Lambda
[Go Support for AWS Lambda](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/)
~~~~
GOOS=linux go build -o lambda_handler
zip handler.zip ./lambda_handler
~~~~

handler = lambda_handler
