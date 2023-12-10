package handlers

import (
	"context"
	"fmt"
	"github.com/DanielaGrosclaude/nuts/models"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) models.RespApi {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + ">" + ctx.Value(models.Key("method")).(string))

	var r models.RespApi
	r.Status = 400

	isOK, statusCode, msg, claim := validoAuthorization(ctx, request)

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

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)
	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, todoOk, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOk {
		if err != nil {
			fmt.Println("Error en el token" + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en el token" + msg)
			return false, 401, msg, models.Claim{}
		}
	}
	fmt.Println("Token ok")
	return true, 200, msg, *claim
}
