package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/models"
)

type UserRepository struct {
	DB *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// Create a new user
func (r *UserRepository) CreateUser(user *models.User) error {

	user.ID = uuid.New().String()

	query := `
	INSERT INTO users
	(id, full_name, email, password)
	VALUES ($1,$2,$3,$4)
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		user.ID,
		user.FullName,
		user.Email,
		user.Password,
	)

	return err
}

// Find user by email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {

	query := `
	SELECT id, full_name, email, password, created_at, updated_at
	FROM users
	WHERE email=$1
	`

	var user models.User

	err := r.DB.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}