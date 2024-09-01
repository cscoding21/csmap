// Code generated . DO NOT EDIT.
// ################################## DO NOT EDIT THIS FILE ######################################
// Common Sense Coding (https://github.com/cscoding21/csgen)

// Generate Date: 2024-09-01 16:09:08.730962745 -0700 PDT m=+49.319163767
// Implementation Name: csmap
// Developer Note: The contents of this file will be recreated each time its generator is called

// -----------------------------------------------------------------------------------------------

package tests

import (
	"github.com/cscoding21/csmap/tests/common"
	"github.com/cscoding21/csmap/tests/pkgco"
	"github.com/cscoding21/csmap/utils"
)

// PagingSourcePkgcoToPkgco converts the source object to the target object.
func PagingSourcePkgcoToPkgco(r pkgco.PagingSource) pkgco.PagingTarget {
	out := pkgco.PagingTarget{
		Page:  r.Page,
		Size:  r.Size,
		Token: r.Token,

		// ---Embedded Structs
		PagingTargetEmbedded: pkgco.PagingTargetEmbedded{
			EmbeddedString: r.EmbeddedString,
			EmbeddedInt:    r.EmbeddedInt,
		},
	}

	return out
}

// PagingSourcePkgcoToPkgcoSlice converts the source object slice to the target object slice.
func PagingSourcePkgcoToPkgcoSlice(r []*pkgco.PagingSource) []*pkgco.PagingTarget {
	out := []*pkgco.PagingTarget{}

	for _, v := range r {
		out = append(out, utils.ValToRef(PagingSourcePkgcoToPkgco(*v)))
	}

	return out
}

// PagingSourceTwoPkgcoToPkgco converts the source object to the target object.
func PagingSourceTwoPkgcoToPkgco(r pkgco.PagingSourceTwo) pkgco.PagingTargetTwo {
	out := pkgco.PagingTargetTwo{
		Page:  r.Page,
		Size:  r.Size,
		Token: r.Token,

		// ---Embedded Structs
		PagingTargetEmbedded: pkgco.PagingTargetEmbedded{
			EmbeddedString: r.EmbeddedString,
			EmbeddedInt:    r.EmbeddedInt,
		},

		// ---Embedded Structs
		ControlFields: common.ControlFields{
			ID:        r.ID,
			CreatedAt: r.CreatedAt,
		},
	}

	return out
}

// PagingSourceTwoPkgcoToPkgcoSlice converts the source object slice to the target object slice.
func PagingSourceTwoPkgcoToPkgcoSlice(r []*pkgco.PagingSourceTwo) []*pkgco.PagingTargetTwo {
	out := []*pkgco.PagingTargetTwo{}

	for _, v := range r {
		out = append(out, utils.ValToRef(PagingSourceTwoPkgcoToPkgco(*v)))
	}

	return out
}
