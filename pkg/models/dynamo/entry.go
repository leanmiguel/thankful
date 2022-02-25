package dynamo

import (
	"leanmiguel/thankful/pkg/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const ThankfulEntriesTableName = "thankful_entries"
const UserIdAttribute = "user_id"
const CreatedTimeAttribute = "created_time"

type EntryModel struct {
	DB *dynamodb.DynamoDB
}

func (m *EntryModel) Insert(user, content string) (int, error) {
	return 0, nil
}

func (m *EntryModel) Get(userId string, createdTime string) (*models.Entry, error) {

	keyCondition := expression.KeyEqual(expression.Key(UserIdAttribute), expression.Value(userId)).And(expression.KeyBeginsWith(expression.Key(CreatedTimeAttribute), createdTime))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()

	if err != nil {
		return nil, err
	}

	result, err := m.DB.Query(&dynamodb.QueryInput{
		TableName:                 aws.String(ThankfulEntriesTableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	if err != nil {
		return nil, err
	}

	entries := []models.Entry{}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &entries)

	if err != nil {
		return nil, err
	}

	return &entries[0], nil
}

func (m *EntryModel) Latest(userId string) (*[]models.Entry, error) {
	keyCondition := expression.KeyEqual(expression.Key(UserIdAttribute), expression.Value(userId))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()

	if err != nil {
		return nil, err
	}

	result, err := m.DB.Query(&dynamodb.QueryInput{
		TableName:                 aws.String(ThankfulEntriesTableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	if err != nil {
		return nil, err
	}

	entries := []models.Entry{}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &entries)

	if err != nil {
		return nil, err
	}

	return &entries, nil
}
