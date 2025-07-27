package repository

import (
	"database/sql"
	"fmt"
	"any-api/internal/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type ItemRepository interface {
	CreateItem(item *models.Item) (*models.Item, error)
}

type PostgresItemRepository struct {
	db *sql.DB
}

func NewPostgresItemRepository(connStr string) (*PostgresItemRepository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Create items table if it doesn't exist
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS items (
		id TEXT PRIMARY KEY,
		name TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &PostgresItemRepository{db: db}, nil
}

func (r *PostgresItemRepository) CreateItem(item *models.Item) (*models.Item, error) {
	item.ID = uuid.New().String()
	_, err := r.db.Exec("INSERT INTO items(id, name) VALUES($1, $2)", item.ID, item.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to insert item: %w", err)
	}
	return item, nil
}