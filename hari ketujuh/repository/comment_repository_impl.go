package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository commentRepositoryImpl) Insert(ctx context.Context, commant entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments (email, commentar) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, commant.Email, commant.Commentar)

	if err != nil {
		return commant, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return commant, err
	}
	commant.Id = int32(id)

	return commant, nil
}

func (repository commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	var comment entity.Comment
	script := "SELECT id, email, commentar FROM comments WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, script, id)

	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Commentar)
		return comment, nil
	} else {
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository commentRepositoryImpl) GetAll(ctx context.Context) ([]entity.Comment, error) {
	var comments []entity.Comment
	script := "SELECT id, email, commentar FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var comment entity.Comment
		rows.Scan(&comment.Id, &comment.Email, &comment.Commentar)
		comments = append(comments, comment)
	}

	return comments, nil
}
