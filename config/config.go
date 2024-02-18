package config

import (
	"github.com/urlShortnerGo/pkg/awsclient"
	"os"
)

type AppConfig struct {
	GeneralConfig *GeneralConfig
	AwsConfigure  *awsclient.AwsConfig
}

type GeneralConfig struct {
	AppPort     string
	Environment string
}

var (
	appConfiguration *AppConfig
	awsRegion        = os.Getenv("AWS_REGION")
)

func Init() *AppConfig {
	appConfiguration = &AppConfig{
		GeneralConfig: &GeneralConfig{
			AppPort:     os.Getenv("APP_PORT"),
			Environment: os.Getenv("APP_ENV"),
		},
		AwsConfigure: &awsclient.AwsConfig{
			AwsRegion: awsRegion,
			Dynamodb: &awsclient.DynamoDBConfig{
				Port:    os.Getenv("DYNAMODB_PORT"),
				Address: os.Getenv("DYNAMODB_URL"),
			},
		},
	}
	return appConfiguration
}

func GetConfig() *AppConfig {
	if appConfiguration == nil {
		return Init()
	}
	return appConfiguration
}
