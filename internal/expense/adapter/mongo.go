package adapter

import (
	"context"
	"gotemplate/internal/expense"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Mongo is the mongo adapter for expense repository
type Mongo struct {
	coll *mongo.Collection
}

// NewMongo creates a new mongo adapter for expense repository
func NewMongo(coll *mongo.Collection) *Mongo {
	return &Mongo{
		coll: coll,
	}
}

// Insert inserts a new expense
func (m *Mongo) Insert(ctx context.Context, e *expense.Expense) (string, error) {
	res, err := m.coll.InsertOne(ctx, newMongoExpense(e))
	if err != nil {
		return "", err
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), nil
}

// Filter filters expenses by given filter
func (m *Mongo) Filter(ctx context.Context, f *expense.Filter) ([]expense.Expense, error) {
	filter := buildExpenseFilter(f)

	cursor, err := m.coll.Find(ctx, filter)
	if err != nil {
		zap.L().Error("failed to find expenses", zap.Error(err))
		return nil, err
	}

	expenses := make([]expense.Expense, 0)
	for cursor.Next(ctx) {
		var e mongoExpense
		if err := cursor.Decode(&e); err != nil {
			return nil, err
		}

		expenses = append(expenses, *e.toExpense())
	}

	return expenses, nil
}

func buildExpenseFilter(f *expense.Filter) bson.M {
	filter := make(bson.M)

	if f.TitleContains != "" {
		filter["title"] = bson.M{"$regex": ".*" + f.TitleContains + ".*"}
	}

	if f.LowerThanPrice != 0 {
		filter["price"] = bson.M{"$lte": f.LowerThanPrice}
	}

	if f.HigherThanPrice != 0 {
		if _, ok := filter["price"]; ok {
			filter["price"].(bson.M)["$gte"] = f.HigherThanPrice
		} else {
			filter["price"] = bson.M{"$gte": f.HigherThanPrice}
		}
	}

	if !f.Before.IsZero() {
		filter["at"] = bson.M{"$lte": f.Before}
	}

	if !f.After.IsZero() {
		if _, ok := filter["at"]; ok {
			filter["at"].(bson.M)["$gte"] = f.After
		} else {
			filter["at"] = bson.M{"$gte": f.After}
		}
	}

	return filter
}

type mongoExpense struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       float32            `bson:"price"`
	At          time.Time          `bson:"at"`
}

func (m *mongoExpense) toExpense() *expense.Expense {
	return &expense.Expense{
		ID:          m.ID.Hex(),
		Title:       m.Title,
		Description: m.Description,
		Price:       m.Price,
		At:          m.At,
	}
}

func newMongoExpense(e *expense.Expense) *mongoExpense {
	return &mongoExpense{
		ID:          primitive.NewObjectID(),
		Title:       e.Title,
		Description: e.Description,
		Price:       e.Price,
		At:          e.At,
	}
}
