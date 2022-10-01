package adapter

import (
	"context"
	"gotemplate/internal/subscription"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type Mongo struct {
	coll *mongo.Collection
}

func NewMongo(coll *mongo.Collection) *Mongo {
	return &Mongo{
		coll: coll,
	}
}

// Insert inserts a new subscription into the database and returns id of the inserted subscription
func (m *Mongo) Insert(ctx context.Context, subs *subscription.Subscription) (string, error) {
	res, err := m.coll.InsertOne(ctx, newMongoSubscription(subs))
	if err != nil {
		zap.L().Error("failed to insert subscription", zap.Error(err))
		return "", err
	}

	oid, _ := res.InsertedID.(primitive.ObjectID)
	return oid.Hex(), err
}

func (m *Mongo) Filter(ctx context.Context, f subscription.Filter) ([]subscription.Subscription, error) {
	filter := make(bson.M)
	if f.Status != "" {
		filter["status"] = f.Status
	}

	if !f.NoticeAt.IsZero() {
		filter["noticeAt"] = f.NoticeAt
	}

	cursor, err := m.coll.Find(ctx, bson.M{"status": string(f.Status)})
	if err != nil {
		zap.L().Error("failed to get all subscriptions", zap.Error(err))
		return nil, err
	}

	subscriptions := make([]subscription.Subscription, 0)
	for cursor.Next(ctx) {
		var mongoSubs mongoSubscription
		if err := cursor.Decode(&mongoSubs); err != nil {
			zap.L().Error("failed to decode subscription", zap.Error(err))
			return nil, err
		}

		subscriptions = append(subscriptions, *mongoSubs.toSubscription())
	}

	return subscriptions, nil
}

type mongoSubscription struct {
	ID               primitive.ObjectID `bson:"_id"`
	Company          string             `bson:"company"`
	Service          string             `bson:"service"`
	Price            float32            `bson:"price"`
	Description      string             `bson:"description"`
	Start            time.Time          `bson:"start"`
	End              time.Time          `bson:"end"`
	PaidInstallments int                `bson:"paid_installments"`
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
		Price:            subs.Price,
		Description:      subs.Description,
		Start:            subs.Start,
		End:              subs.End,
		PaidInstallments: subs.PaidInstallments,
		MonthlyPayday:    subs.MonthlyPayday,
		Notify:           subs.Settings.Notify,
		NoticeBeforeDays: subs.Settings.BeforeDays,
		NoticeAt:         subs.NoticeAt,
		CreatedAt:        time.Now().UTC(),
	}
}

func (ms *mongoSubscription) toSubscription() *subscription.Subscription {
	return &subscription.Subscription{
		ID:               ms.ID.Hex(),
		Company:          ms.Company,
		Service:          ms.Service,
		Price:            ms.Price,
		Description:      ms.Description,
		Start:            ms.Start,
		End:              ms.End,
		PaidInstallments: ms.PaidInstallments,
		MonthlyPayday:    ms.MonthlyPayday,
		Settings: subscription.Settings{
			Notify:     ms.Notify,
			BeforeDays: ms.NoticeBeforeDays,
		},
		NoticeAt: ms.NoticeAt,
	}
}
