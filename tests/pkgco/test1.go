package pkgco

import "time"

// PagingSource used for testing
type PagingSource struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Token string `json:"token"`

	PagingSourceEmbedded
}

type PagingSourceEmbedded struct {
	EmbeddedString string
	EmbeddedInt    int
}

// PagingSource used for testing
type PagingSourceTwo struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Token string `json:"token"`

	EmbeddedString string
	EmbeddedInt    int

	ID        string
	CreatedAt time.Time
}
