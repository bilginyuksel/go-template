package adapter

import (
	"context"
	"gotemplate/internal/subscription"
	"gotemplate/pkg/errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// Mongo is an adapter for subscription that uses MongoDB as the storage
type Mongo struct {
	coll *mongo.Collection
}

// NewMongo returns a new Mongo adapter for subscription
func NewMongo(coll *mongo.Collection) *Mongo {
	return &Mongo{
		coll: coll,
	}
}

// Insert inserts a new subscription into the database and returns id of the inserted subscription
func (m *Mongo) Insert(ctx context.Context, subs *subscription.Subscription) (string, error) {
	res, err := m.coll.InsertOne(ctx, newMongoSubscription(subs))
	if err != nil {
		return "", errors.Wrap(err, "mongo_subscription: failed to insert subscription")
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)

	zap.L().Debug("mongo subscription inserted", zap.String("id", oid.Hex()))

	return oid.Hex(), nil
}

// Filter filters subscriptions based on the given filter
// If the filter is empty, all subscriptions will be returned
func (m *Mongo) Filter(ctx context.Context, f subscription.Filter) ([]subscription.Subscription, error) {
	filter := make(bson.M)
	if f.Status != "" {
		filter["status"] = string(f.Status)
	}

	if !f.NoticeAt.IsZero() {
		filter["noticeAt"] = f.NoticeAt
	}

	zap.L().Debug("mongo subscription filter built", zap.Any("filter", filter))

	cursor, err := m.coll.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "mongo_subscription: failed to get all subscriptions")
	}

	subscriptions := make([]subscription.Subscription, 0)
	for cursor.Next(ctx) {
		var mongoSubs mongoSubscription
		if err := cursor.Decode(&mongoSubs); err != nil {
			return nil, errors.Wrap(err, "mongo_subscription: failed to decode subscription")
		}

		subscriptions = append(subscriptions, *mongoSubs.toSubscription())
	}

	return subscriptions, nil
}

// UpdateNoticeTime updates the notice time of the subscription
func (m *Mongo) UpdateNoticeTime(ctx context.Context, id string, noticeAt time.Time) error {
	oid, _ := primitive.ObjectIDFromHex(id)

	_, err := m.coll.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": bson.M{"notice_at": noticeAt}})
	return errors.Wrap(err, "mongo_subscription: failed to update notice time")
}

type mongoSubscription struct {
	ID               primitive.ObjectID `bson:"_id"`
	Company          string             `bson:"company"`
	Service          string             `bson:"service"`
	Price            float32            `bson:"price"`
	Status           string             `bson:"status"`
	StartedAt        time.Time          `bson:"started_at"`
	MonthlyPayday    int                `bson:"monthly_payday"`
	Notify           bool               `bson:"notify"`
	NoticeBeforeDays int                `bson:"notice_before_days"`
	NoticeAt         time.Time          `bson:"notice_at"`
	CreatedAt        time.Time          `bson:"created_at"`
}

func newMongoSubscription(subs *subscription.Subscription) *mongoSubscription {
	return &mongoSubscription{
		ID:               primitive.NewObjectID(),
		Company:          subs.Company,
		Service:          subs.Service,
		Status:           string(subs.Status),
		Price:            subs.Price,
		StartedAt:        subs.StartedAt,
		MonthlyPayday:    subs.MonthlyPayday,
		Notify:           subs.Settings.Notify,
		NoticeBeforeDays: subs.Settings.BeforeDays,
		NoticeAt:         subs.NoticeAt,
		CreatedAt:        time.Now().UTC(),
	}
}

func (ms *mongoSubscription) toSubscription() *subscription.Subscription {
	return &subscription.Subscription{
		ID:            ms.ID.Hex(),
		Company:       ms.Company,
		Service:       ms.Service,
		Price:         ms.Price,
		StartedAt:     ms.StartedAt,
		Status:        subscription.Status(ms.Status),
		MonthlyPayday: ms.MonthlyPayday,
		Settings: subscription.Settings{
			Notify:     ms.Notify,
			BeforeDays: ms.NoticeBeforeDays,
		},
		NoticeAt: ms.NoticeAt,
	}
}
