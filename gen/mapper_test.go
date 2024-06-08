package gen

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	err := Generate("csmap.yaml")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadManifest(t *testing.T) {
	manifest := LoadManifest("manifest.yaml")

	if manifest.ProjectRoot != "/home/jeph/projects/cscoding21/csmap" {
		t.Errorf("ProjectRoot is incorrect: %s", manifest.ProjectRoot)
	}

	if manifest.GeneratorPath != "tests" {
		t.Errorf("GeneratorPath is incorrect: %s", manifest.GeneratorPath)
	}

	if manifest.GeneratorPackage != "tests" {
		t.Errorf("GeneratorPackage is incorrect: %s", manifest.GeneratorPackage)
	}

	if len(manifest.ObjectMaps) != 1 {
		t.Errorf("Count of ObjectMaps is incorrect: expected 2, got %v", len(manifest.ObjectMaps))
	}
}
