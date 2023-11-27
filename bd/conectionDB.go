package bd

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"nuts/models"
)

var MongoCN *mongo.Client
var DatabaseName string

func ConectionBD(ctx context.Context) error {

	user := ctx.Value(models.Key("user")).(string)
	passwd := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, passwd, host)

	var clientOptions = options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Successful connection to the database")
	MongoCN = client
	db := ctx.Value(models.Key("database")).(string)
	DatabaseName = db
	return nil
}

func ConnectedB() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	return err == nil
}