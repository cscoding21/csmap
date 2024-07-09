package pkgco

// PagingTarget used for testing
type PagingTarget struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Token string `json:"token"`
}
