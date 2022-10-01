package adapter

import (
	"context"
	"gotemplate/internal/subscription"
	"time"

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
