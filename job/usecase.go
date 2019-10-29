package job

import (
	"context"

	"github.com/kyteproject/jumpjobs/models"
)

// Usecase represent the Job's usecases
type Usecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]*models.Job, string, error)
	GetByID(ctx context.Context, id int64) (*models.Job, error)
	Update(ctx context.Context, ar *models.Job) error
	GetByTitle(ctx context.Context, title string) (*models.Job, error)
	Store(context.Context, *models.Job) error
	Delete(ctx context.Context, id int64) error
}
