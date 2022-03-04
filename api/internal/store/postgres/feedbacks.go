package postgres

import (
	"api/internal/models"
	"api/internal/store"
	"context"
	"github.com/jmoiron/sqlx"
)

func (db *DB) Feedback() store.FeedbackRepository {
	if db.feedbacks == nil {
		db.feedbacks = NewFeedbacksRepository(db.conn)
	}
	return db.feedbacks
}

type FeedbacksRepository struct {
	conn *sqlx.DB
}

func (f FeedbacksRepository) Create(ctx context.Context, user *models.Feedback) error {
	panic("implement me")
}

func (f FeedbacksRepository) All(ctx context.Context) ([]*models.Feedback, error) {
	panic("implement me")
}

func (f FeedbacksRepository) ByID(ctx context.Context, id int) ([]*models.Feedback, error) {
	panic("implement me")
}

func (f FeedbacksRepository) Update(ctx context.Context, user *models.Feedback) error {
	panic("implement me")
}

func (f FeedbacksRepository) Delete(ctx context.Context, userId int) error {
	panic("implement me")
}

func NewFeedbacksRepository(conn *sqlx.DB) store.FeedbackRepository {
	return &FeedbacksRepository{conn: conn}
}
