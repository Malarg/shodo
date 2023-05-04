package tests

import (
	"shodo/internal/config"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type APITestSuite struct {
	suite.Suite
	Config   *config.Config
	db       *mongo.Client
	testData TestData
}
