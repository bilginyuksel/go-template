package adapter

import (
	"context"
	"gotemplate/internal/expense"

	"github.com/prometheus/client_golang/prometheus"
)

// WrapperMongoMetrics is a wrapper for the mongo adapter that records metrics
type WrapperMongoMetrics struct {
	hist prometheus.Histogram

	m *Mongo
}

// Filter wraps the mongo filter method and records the time it takes to execute
func (w *WrapperMongoMetrics) Filter(ctx context.Context, f *expense.Filter) ([]expense.Expense, error) {
	timer := prometheus.NewTimer(w.hist)
	defer timer.ObserveDuration()

	return w.m.Filter(ctx, f)
}

// Insert wraps the mongo insert method and records the time it takes to execute
func (w *WrapperMongoMetrics) Insert(ctx context.Context, e *expense.Expense) (string, error) {
	timer := prometheus.NewTimer(w.hist)
	defer timer.ObserveDuration()

	return w.m.Insert(ctx, e)
}

func NewWrapperMongoMetrics(hist prometheus.Histogram, m *Mongo) *WrapperMongoMetrics {
	return &WrapperMongoMetrics{
		hist: hist,
		m:    m,
	}
}
