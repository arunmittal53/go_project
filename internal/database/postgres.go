package database

import (
	"context"
	"errors"
	"fmt"
	"go_project/internal/models"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresClient struct {
	db *gorm.DB
}

func NewPostgresClient() (*PostgresClient, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	DB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	return &PostgresClient{
		db: DB,
	}, nil
}

func (pc *PostgresClient) GormDB() *gorm.DB {
	return pc.db
}

func (pc *PostgresClient) SaveUser(c context.Context, user *models.User) (string, error) {
	err := pc.db.Save(user).Error
	if isDuplicateError(err) {
		return "already exist", err
	}
	if err != nil {
		return "", err
	}
	return "User Saved", nil
}

func (pc *PostgresClient) FetchUser(c context.Context, userId string) (*models.User, error) {
	var user *models.User
	err := pc.db.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (pc *PostgresClient) FetchAllUsers(c context.Context) ([]*models.User, error) {
	var users []*models.User
	err := pc.db.Find(&users).Error
	if err != nil {
		return []*models.User{}, err
	}
	return users, nil
}

func (pc *PostgresClient) DeleteUser(c context.Context, userId string) (string, error) {
	var user models.User
	err := pc.db.Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return "", err
	}
	return "User deleted", nil
}

func isDuplicateError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}
