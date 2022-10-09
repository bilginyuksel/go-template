package adapter_test

import (
	"context"
	"gotemplate/internal/expense"
	"gotemplate/internal/expense/adapter"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	_database      = "gotemplate"
	_collection    = "expenses"
	_exposedPorts  = "27022:27017"
	_connectionURI = "mongodb://localhost:27022"
	_mongoImage    = "mongo:latest"
)

// MongoTestSuite is a suite for testing the mongo adapter
type MongoTestSuite struct {
	suite.Suite

	mongoContainer testcontainers.Container
	mongoClient    *mongo.Client
}

// SetupSuite is called once before the tests in the suite are run
// It starts the mongo container and creates&connects to mongo client
func (m *MongoTestSuite) SetupSuite() {
	ctx := context.Background()

	m.mongoContainer = m.createMongoContainer(ctx)
	m.mongoClient = m.createMongoClient(ctx, _connectionURI)
}

// TearDownTest is called after every test in the suite
// It drops the database for test isolation
func (m *MongoTestSuite) TearDownTest() {
	_ = m.mongoClient.Database(_database).Drop(context.Background())
}

// TearDownSuite is called once after the tests in the suite have been run
// Disconnects from mongo client
// Terminates the mongo container
func (m *MongoTestSuite) TearDownSuite() {
	ctx := context.Background()

	_ = m.mongoClient.Disconnect(ctx)
	_ = m.mongoContainer.Terminate(ctx)
}

func TestMongoSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping mongo adapter tests")
	}

	suite.Run(t, new(MongoTestSuite))
}

func (m *MongoTestSuite) TestInsert() {
	ctx := context.Background()

	expectedExpense := &expense.Expense{
		ID:          "1",
		Title:       "test",
		Description: "test",
		Price:       1,
		At:          time.Now().Round(time.Minute).UTC(),
	}

	mongoAdapter := adapter.NewMongo(m.coll())
	id, err := mongoAdapter.Insert(ctx, expectedExpense)
	m.Nil(err)

	expectedExpense.ID = id

	m.assertExpense(expectedExpense, m.readFromMongo(ctx, id))
}

func (m *MongoTestSuite) TestFilter() {
	ctx := context.Background()

	expectedExpenses := []expense.Expense{
		{
			Title:       "test-1",
			Description: "test-desc-1",
			Price:       15,
			At:          time.Now().Add(-time.Hour).Round(time.Minute).UTC(),
		},
		{
			Title:       "test-2",
			Description: "test-desc-2",
			Price:       50,
			At:          time.Now().Round(time.Minute).UTC(),
		},
		{
			Title:       "test-3",
			Description: "test-desc-3",
			Price:       19,
			At:          time.Now().Add(time.Hour).Round(time.Minute).UTC(),
		},
	}

	ids := m.insertMockExpenses(ctx, expectedExpenses...)
	expectedExpenses[0].ID = ids[0]
	expectedExpenses[1].ID = ids[1]
	expectedExpenses[2].ID = ids[2]

	mongoAdapter := adapter.NewMongo(m.coll())

	filter := &expense.Filter{
		TitleContains:   "test",
		LowerThanPrice:  19,
		HigherThanPrice: 15,
		After:           time.Now().Add(-time.Hour * 2).UTC(),
		Before:          time.Now().Add(time.Hour * 2).UTC(),
	}

	expenses, err := mongoAdapter.Filter(ctx, filter)

	m.Nil(err)
	m.Len(expenses, 2)
	m.Equal(expectedExpenses[0], expenses[0])
	m.Equal(expectedExpenses[2], expenses[1])
}

func (m *MongoTestSuite) insertMockExpenses(ctx context.Context, expenses ...expense.Expense) []string {
	arr := make(bson.A, 0)
	for _, e := range expenses {
		pbson := bson.M{
			"_id":         primitive.NewObjectID(),
			"title":       e.Title,
			"description": e.Description,
			"price":       e.Price,
			"at":          e.At,
		}
		arr = append(arr, pbson)
	}

	res, err := m.coll().InsertMany(ctx, arr)
	if err != nil {
		panic(err)
	}

	var ids []string
	for _, id := range res.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID).Hex())
	}

	return ids
}

func (m *MongoTestSuite) assertExpense(expected *expense.Expense, actual bson.M) {
	m.Equal(expected.ID, actual["_id"].(primitive.ObjectID).Hex())
	m.Equal(expected.Title, actual["title"])
	m.Equal(expected.Description, actual["description"])
	m.Equal(expected.Price, float32(actual["price"].(float64)))
	m.Equal(expected.At, actual["at"].(primitive.DateTime).Time().UTC())
}

func (m *MongoTestSuite) readFromMongo(ctx context.Context, id string) bson.M {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	var result bson.M
	if err := m.mongoClient.Database(_database).Collection(_collection).FindOne(ctx, bson.M{"_id": objID}).Decode(&result); err != nil {
		panic(err)
	}
	return result
}

func (m *MongoTestSuite) coll() *mongo.Collection {
	return m.mongoClient.Database(_database).Collection(_collection)
}

func (m *MongoTestSuite) createMongoContainer(ctx context.Context) testcontainers.Container {
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        _mongoImage,
			ExposedPorts: []string{_exposedPorts},
			WaitingFor:   wait.ForLog("Waiting for connections"),
			AutoRemove:   true,
		},
		Started: true,
	})
	if err != nil {
		panic(err)
	}

	return container
}

func (m *MongoTestSuite) createMongoClient(ctx context.Context, uri string) *mongo.Client {
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(_connectionURI),
	)
	if err != nil {
		panic(err)
	}
	return client
}
