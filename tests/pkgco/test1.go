package pkgco

type PagingSource struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Token string `json:"token"`
}
