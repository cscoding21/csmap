// Code generated . DO NOT EDIT.
// ################################## DO NOT EDIT THIS FILE ######################################
// Common Sense Coding (https://github.com/cscoding21/csgen)

// Generate Date: 2024-08-22 16:39:03.879067552 -0700 PDT m=+3.547173177
// Implementation Name: csmap
// Developer Note: The contents of this file will be recreated each time its generator is called

// -----------------------------------------------------------------------------------------------

package tests

import (
	"github.com/cscoding21/csmap/tests/diffnames"
	"github.com/cscoding21/csmap/utils"
)

// LocationSourceDiffnamesToDiffnames converts the source object to the target object.
func LocationSourceDiffnamesToDiffnames(r diffnames.LocationSource) diffnames.LocationTarget {
	out := diffnames.LocationTarget{
		Lat: r.Lat,
		Lon: r.Lon,
	}

	return out
}

// LocationSourceDiffnamesToDiffnamesSlice converts the source object slice to the target object slice.
func LocationSourceDiffnamesToDiffnamesSlice(r []*diffnames.LocationSource) []*diffnames.LocationTarget {
	out := []*diffnames.LocationTarget{}

	for _, v := range r {
		out = append(out, utils.ValToRef(LocationSourceDiffnamesToDiffnames(*v)))
	}

	return out
}

// TestSourceDiffnamesToDiffnames converts the source object to the target object.
func TestSourceDiffnamesToDiffnames(r diffnames.TestSource) diffnames.TestTarget {
	out := diffnames.TestTarget{
		ID:       r.ID,
		Name:     &r.Name,
		Age:      *r.Age,
		Location: LocationSourceDiffnamesToDiffnames(r.Location),
	}

	return out
}

// TestSourceDiffnamesToDiffnamesSlice converts the source object slice to the target object slice.
func TestSourceDiffnamesToDiffnamesSlice(r []*diffnames.TestSource) []*diffnames.TestTarget {
	out := []*diffnames.TestTarget{}

	for _, v := range r {
		out = append(out, utils.ValToRef(TestSourceDiffnamesToDiffnames(*v)))
	}

	return out
}
