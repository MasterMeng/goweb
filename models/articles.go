package models

import "time"

type Atricles struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Abstract   string    `json:"abstract"`
	CreateDate time.Time `json:"createdate"`
	ModityDate time.Time `json:"moditydate"`
	Content    string    `json:"content"`
}
