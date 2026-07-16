package services

import (
	"errors"

	"github.com/jackc/pgx/v5"

	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/models"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/repository"
	"github.com/Surbhibagad/cloud-native-ai-microservice-platform/services/auth-service/internal/utils"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: repo,
	}
}

func (s *AuthService) Register(req *models.RegisterRequest) error {

	// Check if user already exists
	_, err := s.UserRepo.GetUserByEmail(req.Email)

	if err == nil {
		return errors.New("user already exists")
	}

	// If the error is something other than "no rows", return it.
	if err != pgx.ErrNoRows {
		return err
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	// Create the user model
	user := &models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Save to database
	return s.UserRepo.CreateUser(user)
}

func (s *AuthService) Login(req *models.LoginRequest, jwtSecret string) (string, error) {

	// Find user by email
	user, err := s.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Verify password
	if !utils.CheckPassword(req.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID, jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) GetProfile(userID string) (*models.User, error) {

	user, err := s.UserRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}