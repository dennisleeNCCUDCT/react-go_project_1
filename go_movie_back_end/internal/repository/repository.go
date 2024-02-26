package repository

import (
	"backend/internal/models"
	"database/sql"
)

type DatabaseRepo interface{
Connection() *sql.DB
	ALLMovies()([]*models.Movie, error)
}