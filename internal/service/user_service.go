package service

import (
	"go-backend/internal/models"
	"go-backend/internal/repository"
	"time"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CalculateAge calculates age from date of birth
// Returns the age in years based on whether the birthday has occurred this year
func CalculateAge(dob time.Time) int {
	now := time.Now()
	years := now.Year() - dob.Year()
	
	// Check if birthday hasn't occurred this year
	// Compare month and day to handle leap years correctly
	nowMonth, nowDay := now.Month(), now.Day()
	dobMonth, dobDay := dob.Month(), dob.Day()
	
	if nowMonth < dobMonth || (nowMonth == dobMonth && nowDay < dobDay) {
		years--
	}
	
	return years
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Create(req.Name, dob)
	if err != nil {
		return nil, err
	}

	user.Age = CalculateAge(user.DOB)
	return user, nil
}

func (s *UserService) GetUserByID(id int32) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.Age = CalculateAge(user.DOB)
	return user, nil
}

func (s *UserService) UpdateUser(id int32, req models.UpdateUserRequest) (*models.User, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Update(id, req.Name, dob)
	if err != nil {
		return nil, err
	}

	user.Age = CalculateAge(user.DOB)
	return user, nil
}

func (s *UserService) DeleteUser(id int32) error {
	return s.repo.Delete(id)
}

func (s *UserService) ListUsers(page, pageSize int) ([]models.User, int64, error) {
	users, err := s.repo.List(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}

	// Calculate age for each user
	for i := range users {
		users[i].Age = CalculateAge(users[i].DOB)
	}

	return users, total, nil
}
