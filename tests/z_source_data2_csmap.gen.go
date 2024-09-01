// Code generated . DO NOT EDIT.
// ################################## DO NOT EDIT THIS FILE ######################################
// Common Sense Coding (https://github.com/cscoding21/csgen)

// Generate Date: 2024-09-01 16:08:27.386099042 -0700 PDT m=+7.974300066
// Implementation Name: csmap
// Developer Note: The contents of this file will be recreated each time its generator is called

// -----------------------------------------------------------------------------------------------

package tests

import (
	"github.com/cscoding21/csmap/tests/pkg1"
	"github.com/cscoding21/csmap/tests/pkg2"
	"github.com/cscoding21/csmap/utils"
)

// ActivityPkg1ToPkg2 converts the source object to the target object.
func ActivityPkg1ToPkg2(r pkg1.Activity) pkg2.Activity {
	out := pkg2.Activity{
		ID:       r.ID,
		Type:     r.Type,
		Summary:  r.Summary,
		Detail:   r.Detail,
		Context:  r.Context,
		TargetID: r.TargetID,
		Time:     r.Time,
		Key:      string(r.Key),
	}

	return out
}

// ActivityPkg1ToPkg2Slice converts the source object slice to the target object slice.
func ActivityPkg1ToPkg2Slice(r []*pkg1.Activity) []*pkg2.Activity {
	out := []*pkg2.Activity{}

	for _, v := range r {
		out = append(out, utils.ValToRef(ActivityPkg1ToPkg2(*v)))
	}

	return out
}

// ActivityResultsPkg1ToPkg2 converts the source object to the target object.
func ActivityResultsPkg1ToPkg2(r pkg1.ActivityResults) pkg2.ActivityResults {
	out := pkg2.ActivityResults{
		Paging:  utils.ValToRef(PagingSourcePkg1ToPkgco(utils.RefToVal(r.Paging))),
		Results: ActivityPkg1ToPkg2Slice(r.Results),
	}

	return out
}

// ActivityResultsPkg1ToPkg2Slice converts the source object slice to the target object slice.
func ActivityResultsPkg1ToPkg2Slice(r []*pkg1.ActivityResults) []*pkg2.ActivityResults {
	out := []*pkg2.ActivityResults{}

	for _, v := range r {
		out = append(out, utils.ValToRef(ActivityResultsPkg1ToPkg2(*v)))
	}

	return out
}
