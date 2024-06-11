package gen

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func wrapInValToRef(input string) string {
	return fmt.Sprintf("utils.ValToRef(%s)", input)
}

func wrapInRefToVal(input string) string {
	return fmt.Sprintf("utils.RefToVal(%s)", input)
}

func getFQObjectName(outPackage string, objectPackage string, name string) string {
	if strings.EqualFold(objectPackage, outPackage) {
		return name
	}

	return fmt.Sprintf("%s.%s", objectPackage, name)
}

func getManifestPath(manifestPath ...string) string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	mp := "csmap.yaml"
	if len(manifestPath) > 0 {
		//---user passed in a manifestPath.  Use that instead of the default
		mp = manifestPath[0]

		//---ensure the manifest file has a valid yaml extension
		if !strings.HasSuffix(mp, ".yaml") && !strings.HasSuffix(mp, ".yml") {
			log.Fatalf("Manifest file must be a YAML file. %s is not a YAML file.", mp)
		}

		//---if the manifest file is a relative path, don't prepend the present working directory
		if path.IsAbs(mp) {
			pwd = ""
		}
	} else {
	}

	return path.Join(pwd, mp)
}

func inferPackageFromOutputPath(outPath string) string {
	root, dir := path.Split(outPath)

	log.Println(root, dir)

	return dir
}
