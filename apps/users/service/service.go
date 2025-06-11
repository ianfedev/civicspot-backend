package usecase

import (
	"context"

	"github.com/ianfedev/civicspot-backend/apps/users/domain"
)

// UserService defines application use cases related to the User domain.
type UserService struct {
	repo domain.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// RegisterIfNotExists creates a user if they don't exist by document type and ID.
func (s *UserService) RegisterIfNotExists(ctx context.Context, u *domain.User) error {
	existing, _ := s.repo.GetByDocument(ctx, u.DocumentType, u.DocumentID)
	if existing != nil {
		return nil
	}
	return s.repo.Create(ctx, u)
}

// GetByID retrieves a user by its ID.
func (s *UserService) GetByID(ctx context.Context, id string) (*domain.User, error) {
	return s.repo.GetByID(ctx, id)
}

// GetByDocument retrieves a user by document type and document ID.
func (s *UserService) GetByDocument(ctx context.Context, docType domain.DocumentType, docID string) (*domain.User, error) {
	return s.repo.GetByDocument(ctx, docType, docID)
}

// Deactivate disables the user, either via soft delete or status change.
func (s *UserService) Deactivate(ctx context.Context, id string) error {
	return s.repo.Deactivate(ctx, id)
}
