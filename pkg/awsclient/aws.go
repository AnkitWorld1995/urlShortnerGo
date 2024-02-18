package awsclient

type AwsConfig struct {
	AwsRegion string
	Dynamodb  *DynamoDBConfig
}

type DynamoDBConfig struct {
	Port    string
	Address string
}
