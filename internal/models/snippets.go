package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnipetModel struct {
	DB *sql.DB
}

func (m *SnipetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (m *SnipetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

func (m *SnipetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
