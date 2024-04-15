package main

import (
	"context"
	"errors"
	"fmt"
	"initial/awsgo"
	"initial/bd"
	"initial/models"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()

	if !ParamsValidation() {
		fmt.Println("Error on loading Params. It is needed SecretName")
		err := errors.New("error on loading Params. It is needed SecretName")
		return event, err
	}

	var data models.SingUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("sub = " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()

	if err != nil {
		fmt.Println("Error reading secret on main " + err.Error())
		return event, err
	}

	err = bd.SingUp(data)

	return event, err
}

func ParamsValidation() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam
}
