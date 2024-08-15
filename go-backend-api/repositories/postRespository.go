package repositories

import (
	"context"
	"log"

	"github.com/SoumyadipBhowmik/go-backend/models/db"
	"github.com/SoumyadipBhowmik/go-backend/models/dto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jinzhu/copier"
)

type PostRepository struct {
	Db *pgxpool.Pool
}

func NewPostRepository(Db *pgxpool.Pool) *PostRepository {
	return &PostRepository{Db: Db}
}

func (repo *PostRepository) CreatePost(userId uuid.UUID, imageUrl, description string) *db.Post {
	query := `INSERT INTO POST(user_id, image_url, description)
	VALUES ($1, $2, $3, $4)
	RETURNING user_id, image_url, description`
	var postDTO dto.PostDTO
	err := repo.Db.QueryRow(context.Background(), query, userId, imageUrl, description).Scan(
		&postDTO.UserId,
		&postDTO.ImageUrl,
		&postDTO.Description,
	)
	if err != nil {
		log.Fatalf("error in create post query %v", err)
	}
	var post db.Post
	dtoToDb := copier.Copy(&post, postDTO)
	if dtoToDb != nil {
		log.Fatalf("error copying from dto to db: %v", err)
	}
	return &post
}
