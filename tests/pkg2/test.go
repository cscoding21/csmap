package pkg2

import (
	"time"

	"github.com/cscoding21/csmap/tests/pkgco"
)

type Activity struct {
	ID       string    `json:"id"`
	Type     string    `json:"type"`
	Summary  string    `json:"summary"`
	Detail   *string   `json:"detail,omitempty"`
	Context  string    `json:"context"`
	TargetID *string   `json:"targetID,omitempty"`
	Time     time.Time `json:"time"`
	Key      string
	//Resource      *Resource      `json:"resource,omitempty"`
	//ControlFields *ControlFields `json:"controlFields"`
}

type ActivityResults struct {
	Paging *pkgco.PagingTarget `json:"paging_target"`
	//Filters *Filters    `json:"filters"`
	Results []*Activity `json:"results,omitempty"`
}
