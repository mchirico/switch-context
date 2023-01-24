package sts

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	sessionV1 "github.com/aws/aws-sdk-go/aws/session"
	stsV1 "github.com/aws/aws-sdk-go/service/sts"
	"os"
)

type GetCallerIdentityAPI interface {
	GetCallerIdentity(ctx context.Context,
		params *sts.GetCallerIdentityInput,
		optFns ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error)
}

type GetCallerIdentityV1API interface {
	GetCallerIdentity(params *stsV1.GetCallerIdentityInput) (*stsV1.GetCallerIdentityOutput, error)
}

func GetCallerIdentity(ctx context.Context, client GetCallerIdentityAPI, params *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	return client.GetCallerIdentity(ctx, params)
}

func GetCallerIdentityV1(client GetCallerIdentityV1API, params *stsV1.GetCallerIdentityInput) (*stsV1.GetCallerIdentityOutput, error) {
	return client.GetCallerIdentity(params)
}

func V1() error {

	sess := sessionV1.Must(sessionV1.NewSessionWithOptions(sessionV1.Options{
		SharedConfigState: sessionV1.SharedConfigEnable,
	}))
	sess.Config.Region = aws.String("us-east-2")

	svc := stsV1.New(sess)

	result, err := GetCallerIdentityV1(svc, &stsV1.GetCallerIdentityInput{})
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "\nVersion 1\n")
	fmt.Fprintf(os.Stderr, "Account ID: %s\n", *result.Account)
	fmt.Fprintf(os.Stderr, "User ID: %s\n", *result.UserId)
	fmt.Fprintf(os.Stderr, "ARN: %s\n", *result.Arn)
	fmt.Fprintf(os.Stderr, "\n")
	return nil

}

func V1NoShared() error {

	sess := sessionV1.Must(sessionV1.NewSessionWithOptions(sessionV1.Options{}))
	sess.Config.Region = aws.String("us-east-2")

	svc := stsV1.New(sess)

	result, err := GetCallerIdentityV1(svc, &stsV1.GetCallerIdentityInput{})
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "\nNO AWS SSO/Identity: Version 1 NoShared\n")
	fmt.Fprintf(os.Stderr, "Account ID: %s\n", *result.Account)
	fmt.Fprintf(os.Stderr, "User ID: %s\n", *result.UserId)
	fmt.Fprintf(os.Stderr, "ARN: %s\n", *result.Arn)
	fmt.Fprintf(os.Stderr, "\n")
	return nil

}

func V2() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	client := sts.NewFromConfig(cfg)

	result, err := GetCallerIdentity(context.TODO(), client, &sts.GetCallerIdentityInput{})
	if err != nil {

		return err
	}
	fmt.Fprintf(os.Stderr, "\nVersion 2\n")
	fmt.Fprintf(os.Stderr, "Account ID: %s\n", *result.Account)
	fmt.Fprintf(os.Stderr, "User ID: %s\n", *result.UserId)
	fmt.Fprintf(os.Stderr, "ARN: %s\n", *result.Arn)
	fmt.Fprintf(os.Stderr, "\n")

	return nil
}

func Pr() {
	err := V1()
	if err != nil {
		fmt.Println(err)
	}
	err = V2()
	if err != nil {
		fmt.Println(err)
	}

	err = V1NoShared()
	if err != nil {
		fmt.Printf("You appear to be picking up AWS SSO\n\n")
		fmt.Println(err)
	}
	fmt.Fprintf(os.Stderr, "\naws sts get-caller-identity\n\n")
}
