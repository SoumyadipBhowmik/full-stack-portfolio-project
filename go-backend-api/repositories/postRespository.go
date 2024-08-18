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
	query := `INSERT INTO post(user_id, image_url, description)
	VALUES ($1, $2, $3)
	RETURNING id, user_id, image_url, description, reaction, created_at, updated_at`
	var postDTO dto.PostDTO
	err := repo.Db.QueryRow(context.Background(), query, userId, imageUrl, description).Scan(
		&postDTO.Id,
		&postDTO.UserId,
		&postDTO.ImageUrl,
		&postDTO.Description,
		&postDTO.Reactions,
		&postDTO.CreatedAt,
		&postDTO.UpdatedAt,
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

func (repo *PostRepository) FetchPost(id uuid.UUID) *db.Post {
	query := `
		SELECT id, user_id, image_url, description, reaction, created_at, updated_at
		FROM post
		WHERE id = $1`

	var post db.Post
	err := repo.Db.QueryRow(context.Background(), query, id).Scan(
		&post.Id,
		&post.UserId,
		&post.ImageUrl,
		&post.Description,
		&post.Reactions,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		log.Fatalf("error scanning from the database: %v", err.Error())
	}
	return &post
}

func (repo *PostRepository) FetchAll() *[]db.Post {

	var posts []db.Post
	query :=
		`SELECT *
		FROM post
		`
	rows, err := repo.Db.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("error in database: %v", err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var post db.Post
		err := rows.Scan(
			&post.Id,
			&post.UserId,
			&post.ImageUrl,
			&post.Description,
			&post.Reactions,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Fatalf("error scanning rows %v", err.Error())
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err.Error())
	}
	return &posts
}

func (repo *PostRepository) FetchUserPosts(userId uuid.UUID) *[]db.Post {
	query := `
	SELECT * FROM post
	WHERE user_id = $1
	`
	var posts []db.Post
	rows, err := repo.Db.Query(context.Background(), query, userId)
	if err != nil {
		log.Fatal(err.Error())
	}
	for rows.Next() {
		var post db.Post
		err := rows.Scan(
			&post.Id,
			&post.UserId,
			&post.ImageUrl,
			&post.Description,
			&post.Reactions,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			log.Fatalf("error scanning rows: %v", err.Error())
		}
		posts = append(posts, post)
	}
	if rows.Err() != nil {
		log.Fatalf("there's a error in the rows: %v", err.Error())
	}
	return &posts
}
