package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"nuts/awsgo"
	"nuts/bd"
	"nuts/handlers"
	"nuts/models"
	"nuts/secretmanager"

	"os"
	"strings"
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
	secretName := os.Getenv("SecretName")
	SecretModel, err := secretmanager.GetSecret(secretName)
	//SecretModel, err := secretsmanager.GetSecret(os.Getenv("SecretName"))
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

	path := strings.Replace(request.PathParameters["nuts"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	//Check bd
	bd.ConectionBD(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "error in the connection with the database" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	respAPI := handlers.Handler(awsgo.Ctx, request)

	if respAPI.CustomResp == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		return respAPI.CustomResp, nil
	}
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
