package dynamo

import "github.com/aws/aws-sdk-go/service/dynamodb"

type EntriesModel struct {
	DB *dynamodb.DynamoDB
}

// func (m *EntriesModel)
