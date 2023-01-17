package sts

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"os"
)

type GetCallerIdentityAPI interface {
	GetCallerIdentity(ctx context.Context,
		params *sts.GetCallerIdentityInput,
		optFns ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error)
}

func GetCallerIdentity(ctx context.Context, client GetCallerIdentityAPI, params *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	return client.GetCallerIdentity(ctx, params)
}

func Pr() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	client := sts.NewFromConfig(cfg)

	result, err := GetCallerIdentity(context.TODO(), client, &sts.GetCallerIdentityInput{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(os.Stderr, "Account ID: %s\n", *result.Account)
	fmt.Fprintf(os.Stderr, "User ID: %s\n", *result.UserId)
	fmt.Fprintf(os.Stderr, "ARN: %s\n", *result.Arn)

}
