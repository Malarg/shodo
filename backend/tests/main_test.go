package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"shodo/internal/config"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type APITestSuite struct {
	suite.Suite
	Config *config.Config
	db     *mongo.Client
}

func (s *APITestSuite) SetupSuite() {
	var err error
	s.Config, err = config.Init("main")
	if err != nil {
		s.FailNow("Failed to init config file", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoUrl := fmt.Sprintf("mongodb://admin:password@%s:27017", s.Config.MongoHost)
	fmt.Println(mongoUrl)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		s.FailNow("Failed to connect to mongo", err)
	}

	s.db = client

	err = s.db.Database(s.Config.DbName).Drop(ctx)
	if err != nil {
		s.FailNow("Failed to drop database", err)
	}
}

func (s *APITestSuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.db.Database(s.Config.DbName).Drop(ctx)
	if err != nil {
		s.FailNow("Failed to drop database", err)
	}
}

func (s *APITestSuite) TearDownSuite() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.db.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(APITestSuite))
}
