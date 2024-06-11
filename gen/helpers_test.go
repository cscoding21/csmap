package gen

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestGetManifestPath(t *testing.T) {
	pwd, _ := os.Getwd()
	testCases := []struct {
		ok   bool
		have string
		want string
	}{
		{ok: true, have: "", want: "csmap.yaml"},
		{ok: true, have: "csmap.yaml", want: fmt.Sprintf("%s/csmap.yaml", pwd)},
		{ok: true, have: "/test/csmap.yaml", want: "/test/csmap.yaml"},
	}

	for _, testCase := range testCases {
		if testCase.have == "" {
			mp := getManifestPath()
			if mp != path.Join(pwd, "csmap.yaml") {
				t.Errorf("getManifestPath failed on zero arguments: got %s", mp)
			}
		} else {
			mp := getManifestPath(testCase.have)
			if mp != testCase.want {
				t.Errorf("getManifestPath failed: expected %s, got %s", testCase.want, mp)
			}
		}
	}
}

func TestInferPackageFromOutputPath(t *testing.T) {
	testCases := []struct {
		ok   bool
		have string
		want string
	}{
		{ok: true, have: "cmd/tests", want: "tests"},
		{ok: true, have: "/usr/blah/foo/bar", want: "bar"},
		{ok: true, have: "./bar", want: "bar"},
		{ok: true, have: "../bar", want: "bar"},
	}

	for _, testCase := range testCases {
		if want := inferPackageFromOutputPath(testCase.have); want != testCase.want {
			t.Errorf("inferPackageFromOutputPath failed: expected %s, got %s", testCase.want, want)
		}
	}
}
