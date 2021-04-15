package models

// AtricleType
type AtricleType int

const ()

type Atricle struct {
	ID         string      `json:"id" db:"id"`
	Owner      string      `json:"owner" db:"owner"`
	Title      string      `json:"title" db:"title"`
	Type       AtricleType `json:"type" db:"type"`
	Abstract   string      `json:"abstract" db:"abstract"`
	CreateDate int64       `json:"createdate" db:"created"`
	ModityDate int64       `json:"moditydate" db:"modity"`
	Content    string      `json:"content" db:"content"`
}

type Message struct {
	ID         string `json:"id" db:"id"`
	AtricleID  string `json:"atricle" db:"atricle"`
	UserID     string `json:"user" db:"user"`
	CreateDate int64  `json:"createdate" db:"created"`
	Content    string `json:"content" db:"content"`
}
