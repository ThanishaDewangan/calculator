package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	db "go-backend/db/sqlc"
	"go-backend/internal/models"
)

type UserRepository struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		queries: db.New(pool),
		pool:    pool,
	}
}

func (r *UserRepository) Create(name string, dob time.Time) (*models.User, error) {
	ctx := context.Background()
	dbUser, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		DOB:       dbUser.Dob,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}, nil
}

func (r *UserRepository) GetByID(id int32) (*models.User, error) {
	ctx := context.Background()
	dbUser, err := r.queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		DOB:       dbUser.Dob,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}, nil
}

func (r *UserRepository) Update(id int32, name string, dob time.Time) (*models.User, error) {
	ctx := context.Background()
	dbUser, err := r.queries.UpdateUser(ctx, db.UpdateUserParams{
		Name: name,
		Dob:  dob,
		ID:   id,
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		DOB:       dbUser.Dob,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}, nil
}

func (r *UserRepository) Delete(id int32) error {
	ctx := context.Background()
	return r.queries.DeleteUser(ctx, id)
}

func (r *UserRepository) List(page, pageSize int) ([]models.User, error) {
	ctx := context.Background()
	
	offset := (page - 1) * pageSize
	dbUsers, err := r.queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	users := make([]models.User, len(dbUsers))
	for i, dbUser := range dbUsers {
		users[i] = models.User{
			ID:        dbUser.ID,
			Name:      dbUser.Name,
			DOB:       dbUser.Dob,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
		}
	}

	return users, nil
}

func (r *UserRepository) Count() (int64, error) {
	ctx := context.Background()
	return r.queries.CountUsers(ctx)
}
