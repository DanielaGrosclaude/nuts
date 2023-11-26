package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"nuts/awsgo"
	"nuts/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecret models.Secret

	fmt.Println("> Pido Secreto" + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		secretsId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}
	json.Unmarshal([]byte(*clave.SecretString), &dataSecret)
	fmt.Println("> Lectura de secret ok" + secretName)
	return dataSecret, nil
}
