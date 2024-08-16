package gen

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"

	"github.com/cscoding21/csgen"
	"github.com/cscoding21/csmap/tests"
	"github.com/cscoding21/csmap/tests/pkg1"
	"golang.org/x/tools/go/packages"
)

const (
	ManifestPath = "csmap.yaml"
)

func TestGenerate(t *testing.T) {
	mcfg := getTestModuleConfig()
	err := Generate(ManifestPath, mcfg)
	if err != nil {
		t.Error(err)
	}

	manifest := LoadManifest(ManifestPath, mcfg)
	expectedFiles := []string{
		filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "z_source_data1_csmap.gen.go"),
		filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "z_source_data2_csmap.gen.go"),
		filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "z_source_pkgco_csmap.gen.go"),
	}

	for _, file := range expectedFiles {
		_, err := os.Stat(file)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestFunctionsCreated(t *testing.T) {
	mcfg := getTestModuleConfig()
	manifest := LoadManifest(ManifestPath, mcfg)
	err := Generate(ManifestPath, mcfg)
	if err != nil {
		t.Error(err)
	}

	testCases := []struct {
		ok   bool
		file string
		want []string
	}{
		{ok: true, file: filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "z_source_data1_csmap.gen.go"), want: []string{"TestSourceDiffnamesToDiffnames", "TestSourceDiffnamesToDiffnamesSlice", "LocationSourceDiffnamesToDiffnames", "LocationSourceDiffnamesToDiffnamesSlice"}},
		{ok: true, file: filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "z_source_data2_csmap.gen.go"), want: []string{"ActivityPkg1ToPkg2", "ActivityPkg1ToPkg2Slice", "ActivityResultsPkg1ToPkg2", "ActivityResultsPkg1ToPkg2Slice"}},
		{ok: true, file: filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "z_source_pkgco_csmap.gen.go"), want: []string{"PagingSourcePkgcoToPkgco", "PagingSourcePkgcoToPkgcoSlice"}},
	}

	for _, file := range testCases {
		//---get a list of functions in the generated file
		functions, err := csgen.GetFunctions(file.file)
		if err != nil {
			t.Error(err)
			continue
		}

		fns := []string{}
		for _, funcs := range functions {
			fns = append(fns, funcs.Name)
		}

		//---the comparison needs elements to be in the same order
		slices.Sort(fns)
		slices.Sort(file.want)

		//---ensure all of the functions that are expected have been created
		if slices.Equal(fns, file.want) == false {
			t.Errorf("Count of functions in %s is incorrect: expected %v, got %v", file.file, len(file.want), len(functions))
		}
	}
}

func TestMappingFunctions(t *testing.T) {
	mcfg := getTestModuleConfig()
	err := Generate(ManifestPath, mcfg)
	if err != nil {
		t.Error(err)
	}

	targetID := "121b"
	testSource := pkg1.Activity{
		ID:       "121e",
		Type:     "update",
		Summary:  "Test activity",
		Detail:   nil,
		Context:  "Test context",
		TargetID: &targetID,
		Time:     time.Now(),
		Key:      "Test key",
	}

	testDest := tests.ActivityPkg1ToPkg2(testSource)

	if testDest.ID != testSource.ID {
		t.Errorf("Mapped ID is incorrect: %s", testDest.ID)
	}

	if testDest.Type != testSource.Type {
		t.Errorf("Mapped Type is incorrect: %s", testDest.Type)
	}

	if testDest.Summary != testSource.Summary {
		t.Errorf("Mapped Summary is incorrect: %s", testDest.Summary)
	}

	if testDest.Detail != testSource.Detail {
		t.Errorf("Mapped Detail is incorrect: %v", testDest.Detail)
	}

	if testDest.Context != testSource.Context {
		t.Errorf("Mapped Context is incorrect: %s", testDest.Context)
	}

	if testDest.TargetID != testSource.TargetID {
		t.Errorf("Mapped TargetID is incorrect: %v", testDest.TargetID)
	}

	if testDest.Time != testSource.Time {
		t.Errorf("Mapped Time is incorrect: %s", testDest.Time)
	}

	if testDest.Key != string(testSource.Key) {
		t.Errorf("Mapped Key is incorrect: %s", testDest.Key)
	}
}

func TestLoadManifest(t *testing.T) {
	mcfg := getTestModuleConfig()

	manifest := LoadManifest("csmap.yaml", mcfg)
	expectedMapCount := 3

	if manifest.ProjectRoot != "/home/jeph/projects/cscoding21/csmap" {
		t.Errorf("ProjectRoot is incorrect: %s", manifest.ProjectRoot)
	}

	if manifest.GeneratorPath != "tests" {
		t.Errorf("GeneratorPath is incorrect: %s", manifest.GeneratorPath)
	}

	if manifest.GeneratorPackage != "tests" {
		t.Errorf("GeneratorPackage is incorrect: %s", manifest.GeneratorPackage)
	}

	if len(manifest.ObjectMaps) != expectedMapCount {
		t.Errorf("Count of ObjectMaps is incorrect: expected %v, got %v", expectedMapCount, len(manifest.ObjectMaps))
	}
}

func getTestModuleConfig() *packages.Config {
	cfg := csgen.GetDefaultPackageConfig()
	cfg.Dir = "../."
	cfg.Tests = true
	return cfg
}
