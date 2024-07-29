package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserFeedbackRepository struct {
	db *pgxpool.Pool
}

func NewUserFeedBackRepository(database *pgxpool.Pool) *UserFeedbackRepository {
	return &UserFeedbackRepository{
		db: database,
	}
}

func (repo *UserFeedbackRepository) CreateNewFeedback(name, email, feedback string) error {
	query := `
	INSERT INTO user_feedbacks (name, email, feedback)
	VALUES ($1, $2, $3)
	RETURNING id
	`
	var id uuid.UUID
	err := repo.db.QueryRow(context.Background(), query, name, email, feedback).Scan(&id)
	if err != nil {
		fmt.Println("The error is: ", err)
	}
	fmt.Println(id)
	return nil
}
