package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"nuts/awsgo"
	"os"
)

func main() {
	lambda.Start(EjecLambda)
}

func EjecLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	awsgo.InitAWS()

	if !ValParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error in environment variables. They must include 'SecretName', 'BucketName', 'UrlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}
	SecretModel, err := secretsmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error in reading secrets" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	return res, nil
}

func ValParams() bool {
	_, getParams := os.LookupEnv("SecretName")
	if !getParams {
		return getParams
	}

	_, getParams = os.LookupEnv("BucketName")
	if !getParams {
		return getParams
	}

	_, getParams = os.LookupEnv("UrlPrefix")
	if !getParams {
		return getParams
	}

	return getParams
}
