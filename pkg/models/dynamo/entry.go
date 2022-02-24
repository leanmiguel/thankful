package dynamo

import (
	"leanmiguel/thankful/pkg/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type EntryModel struct {
	DB *dynamodb.DynamoDB
}

func (m *EntryModel) Insert(user, content string) (int, error) {
	return 0, nil
}

func (m *EntryModel) Get(id string, sortKey string) (*models.Entry, error) {
	return nil, nil
}

func (m *EntryModel) Latest() ([]*models.Entry, error) {
	return nil, nil
}
