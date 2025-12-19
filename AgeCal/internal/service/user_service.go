package service

import (
	"AgeCal/internal/models"
	"AgeCal/internal/repository"
	"context"
	"time"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

func (s *UserService) GetByID(ctx context.Context, id int32) (models.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob,
		Age:  calculateAge(u.Dob),
	}, nil
}

func (s *UserService) List(ctx context.Context) ([]models.User, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	result := []models.User{}
	for _, u := range users {
		result = append(result, models.User{
			ID:   int(u.ID),
			Name: u.Name,
			DOB:  u.Dob,
			Age:  calculateAge(u.Dob),
		})
	}
	return result, nil
}

func (s *UserService) Create(ctx context.Context, name string, dob string) (models.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return models.User{}, err
	}

	u, err := s.repo.Create(ctx, name, parsedDOB)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob,
	}, nil
}

func (s *UserService) Update(ctx context.Context, id int32, name string, dob string) (models.User, error) {
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return models.User{}, err
	}

	u, err := s.repo.Update(ctx, id, name, parsedDOB)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:   int(u.ID),
		Name: u.Name,
		DOB:  u.Dob,
		Age:  calculateAge(u.Dob),
	}, nil
}

func (s *UserService) Delete(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
