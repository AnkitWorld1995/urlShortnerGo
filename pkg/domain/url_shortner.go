package domain

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type UrlShortnerRepository struct {
	AwsSession *dynamodb.DynamoDB
}

func NewShortnerRepo(sess *dynamodb.DynamoDB) *UrlShortnerRepository {
	return &UrlShortnerRepository{
		AwsSession: sess,
	}
}

type UrlRepo interface{}
