package repository

import (
	db "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestNewComment(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnectionDb())
	comment := entity.Comment{
		Commentar: "ini komentar repository",
		Email:     "comment@gmail.com",
	}
	ctx := context.Background()
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindComment(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnectionDb())
	ctx := context.Background()
	result, err := commentRepository.FindById(ctx, 111)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestGetAllComment(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnectionDb())
	ctx := context.Background()
	result, err := commentRepository.GetAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
