package repository

import (
	"belajar-golang-database/entity"
	"context"
)

type CommentRepository interface {
	Insert(ctx context.Context, commant entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	GetAll(ctx context.Context) ([]entity.Comment, error)
}
