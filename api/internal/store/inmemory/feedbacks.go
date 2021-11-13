package inmemory

import (
	"api/internal/models"
	"context"
	"fmt"
	"sync"
	"time"
)

type FeedbackRepo struct {
	data map[int]*models.Feedback

	mu *sync.RWMutex
}

func (db *FeedbackRepo) Create(ctx context.Context, feedback *models.Feedback) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	feedback.CreatedAt = time.Now()
	db.data[feedback.ID] = feedback
	return nil
}

func (db *FeedbackRepo) All(ctx context.Context) ([]*models.Feedback, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	feedbacks := make([]*models.Feedback, 0, len(db.data))
	for _, feedback := range db.data {
		feedbacks = append(feedbacks, feedback)
	}

	return feedbacks, nil
}

func (db *FeedbackRepo) ByID(ctx context.Context, userId int) ([]*models.Feedback, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	feedbacks := make([]*models.Feedback, 0, len(db.data))
	for _, feedback := range db.data {
		if feedback == db.data[userId] {
			fmt.Println(feedback)
			feedbacks = append(feedbacks, feedback)
		}
	}

	return feedbacks, nil
}

func (db *FeedbackRepo) Update(ctx context.Context, feedback *models.Feedback) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.data[feedback.ID] = feedback
	return nil
}

func (db *FeedbackRepo) Delete(ctx context.Context, id int) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, id)
	return nil
}
