package gen

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerate(t *testing.T) {
	err := Generate("csmap.yaml")
	if err != nil {
		t.Error(err)
	}

	manifest := LoadManifest("csmap.yaml")
	expectedFiles := []string{
		filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "source_data1_csmap.gen.go"),
		filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "source_data2_csmap.gen.go"),
		filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath, "source_pkgco_csmap.gen.go"),
	}

	for _, file := range expectedFiles {
		_, err := os.Stat(file)
		if err != nil {
			t.Error(err)
		}
	}

	//cleanup(expectedFiles)
}

func cleanup(files []string) {
	for _, f := range files {
		os.Remove(f)
	}
}

func TestLoadManifest(t *testing.T) {
	manifest := LoadManifest("csmap.yaml")
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
