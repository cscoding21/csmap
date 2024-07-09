package pkg1

import (
	"time"

	"github.com/cscoding21/csmap/tests/pkgco"
)

// TestKey used for testing
type TestKey string

// Activity used for testing
type Activity struct {
	ID       string    `json:"id"`
	Type     string    `json:"type"`
	Summary  string    `json:"summary"`
	Detail   *string   `json:"detail,omitempty"`
	Context  string    `json:"context"`
	TargetID *string   `json:"targetID,omitempty"`
	Time     time.Time `json:"time"`
	Key      TestKey   `json:"key"`
	//Resource      *Resource      `json:"resource,omitempty"`
	//ControlFields *ControlFields `json:"controlFields"`
}

// ActivityResults used for testing
type ActivityResults struct {
	Paging *pkgco.PagingSource `json:"paging"`
	//Filters *Filters    `json:"filters"`
	Results []*Activity `json:"results,omitempty"`
}
