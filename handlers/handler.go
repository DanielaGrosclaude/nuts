package handlers

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"nuts/models"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + ">" + ctx.Value(models.Key("method")).(string))

	var r models.RespApi

	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		}
	}
	r.Message = "Method invalid"
	return r

}