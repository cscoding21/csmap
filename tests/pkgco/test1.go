package pkgco

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
