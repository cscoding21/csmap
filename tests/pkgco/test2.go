package pkgco

// PagingTarget used for testing
type PagingTarget struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Token string `json:"token"`

	PagingTargetEmbedded
}

type PagingTargetEmbedded struct {
	EmbeddedString string
	EmbeddedInt    int
}

type PagingTargetTwo struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Token string `json:"token"`

	PagingTargetEmbedded
}
