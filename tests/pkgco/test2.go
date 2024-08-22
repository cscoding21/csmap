package pkgco

import (
	"github.com/cscoding21/csmap/tests/common"
)

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

	common.ControlFields
}
