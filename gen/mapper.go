package gen

import (
	"fmt"

	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/cscoding21/csgen"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// Generate generates the mapping files.
func Generate(manifestPath string) error {
	mp := getManifestPath(manifestPath)
	manifest := LoadManifest(mp)

	outPath := filepath.Join(manifest.ProjectRoot, manifest.GeneratorPath)
	err := os.MkdirAll(outPath, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, m := range manifest.ObjectMaps {
		builder := csgen.NewCSGenBuilderForFile("csmap", manifest.GeneratorPackage)

		for _, targetStruct := range m.TargetObjects {
			sourceObjectName := getSourceObjectName(targetStruct.Name, m.MapOverrides)
			sourceStruct := objectSliceContainsName(sourceObjectName, m.SourceObjects)

			if sourceStruct == nil {
				continue
			}

			log.Printf("Source Object Found: %s\n", sourceStruct.Name)

			params := getConversionFunctionParams(
				manifest.GeneratorPackage,
				targetStruct.Name,
				targetStruct.Package,
				sourceObjectName,
				sourceStruct.Package,
				m.Imports,
				[]string{})

			props := []string{}
			for _, targetField := range targetStruct.Fields {
				sourceField := sourceStruct.GetField(targetField.Name)
				if sourceField == nil {
					log.Printf("XXX No corresponding source field found for target field: %s\n", targetField.Name)
					continue
				}

				log.Printf("----- Corresponding Source Field Found: %s --- %s\n", targetField.Name, sourceField.Name)

				if targetField.IsSlice {
					if targetField.IsPrimitive {
						props = append(props, getSimpleAssignmentSlice(*sourceField, targetField))
					} else {
						props = append(props, getObjectAssignmentSlice(*sourceField, targetField, sourceStruct.Package, targetStruct.Package))
					}
				} else {
					if targetField.IsPrimitive {
						props = append(props, getSimpleAssignment(*sourceField, targetField))
					} else {
						props = append(props, getObjectAssignment(*sourceField, targetField, sourceStruct.Package, targetStruct.Package))
					}
				}
			}

			params.Assignments = props

			contents := getConversionFunction(strings.ToLower(targetStruct.Name), functionTemplate, params)
			sliceContents := getConversionFunction(strings.ToLower(targetStruct.Name), functionSliceTemplate, params)

			builder.WriteString(contents)
			builder.WriteString(sliceContents)
		}

		outputFileName := csgen.GetFileName("csmap", outPath, m.Name)
		outputFileContents := string(builder.String())

		err = csgen.WriteGeneratedGoFile(outputFileName, outputFileContents)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func objectSliceContainsName(name string, graph []csgen.Struct) *csgen.Struct {
	for _, o := range graph {
		if strings.EqualFold(o.Name, name) {
			return &o
		}
	}

	return nil
}

func getSourceObjectName(name string, overrides []MapOverride) string {
	if len(overrides) == 0 {
		return name
	}

	for _, o := range overrides {
		if strings.EqualFold(o.TargetName, name) {
			return o.SourceName
		}
	}

	return name
}

func getSimpleAssignment(source csgen.Field, target csgen.Field) string {
	rawTargetName := csgen.GetRawType(target.Name)

	indicator := csgen.GetFieldIndicator(source, target)

	if source.Type != target.Type && indicator == "" {
		return fmt.Sprintf("%s: %s(r.%s)", rawTargetName, target.Type, rawTargetName)
	}

	return fmt.Sprintf("%s: %sr.%s", rawTargetName, indicator, rawTargetName)
}

func getSimpleAssignmentSlice(source csgen.Field, target csgen.Field) string {
	rawTargetName := csgen.GetRawType(target.Name)

	indicator := csgen.GetFieldIndicator(source, target)

	if source.Type != target.Type && indicator == "" {
		return fmt.Sprintf("%s: %s(r.%s)", rawTargetName, target.Type, rawTargetName)
	}

	return fmt.Sprintf("%s: %sr.%s", rawTargetName, indicator, rawTargetName)
}

func getObjectAssignment(source csgen.Field, target csgen.Field, sourcePackage string, targetPackage string) string {
	targetTypeLocal := target.Type
	sourceTypeLocal := source.Type

	convertToVal := csgen.IsRefType(sourceTypeLocal)

	log.Printf("OBJECT TYPE COMPARE: %s - %s\n", targetTypeLocal, sourceTypeLocal)
	log.Printf("Convert: to val - %v\n", convertToVal)

	coerceType := csgen.IsPrimitive(sourceTypeLocal) && !csgen.IsPrimitive(targetTypeLocal)

	log.Printf("COERCE TYPE: %v\n", coerceType)

	targetPropertyName := csgen.GetRawType(target.Name)
	rawTargetName := csgen.GetRawType(target.Type)
	objectPackage := csgen.ExtractPackageName(rawTargetName)

	if len(objectPackage) > 0 {
		sourcePackage = objectPackage
		targetPackage = objectPackage
	}

	strippedSourceName := csgen.GetRawType(csgen.StripPackageName(source.Type))
	funcName := getConversionFunctionName(strippedSourceName, sourcePackage, targetPackage)

	if coerceType {
		if csgen.IsFullyQualifiedPackage(targetTypeLocal) {
			funcName = targetTypeLocal
		} else {
			funcName = fmt.Sprintf("%s.%s", targetPackage, targetTypeLocal)
		}
	}

	funcInputArg := fmt.Sprintf("r.%s", source.Name)

	if convertToVal {
		funcInputArg = wrapInRefToVal(funcInputArg)
	}

	funcDef := fmt.Sprintf("%s(%s)", funcName, funcInputArg)

	if target.IsPointer {
		funcDef = wrapInValToRef(funcDef)
	}

	return fmt.Sprintf("%s: %s", targetPropertyName, funcDef)
}

func getObjectAssignmentSlice(source csgen.Field, target csgen.Field, sourcePackage string, targetPackage string) string {
	targetTypeLocal := target.Type
	sourceTypeLocal := source.Type

	convertToVal := !csgen.IsRefType(targetTypeLocal) && csgen.IsRefType(sourceTypeLocal)
	convertToRef := csgen.IsRefType(targetTypeLocal) && !csgen.IsRefType(sourceTypeLocal)

	log.Printf("SLICE TYPE COMPARE: %s - %s\n", targetTypeLocal, sourceTypeLocal)
	log.Printf("Convert: to val - %v | to ref - %v\n", convertToVal, convertToRef)

	targetPropertyName := csgen.GetRawType(target.Name)

	strippedSourceName := csgen.GetRawType(csgen.StripPackageName(source.Type))
	funcName := getConversionFunctionNameToSlice(strippedSourceName, sourcePackage, targetPackage)

	if convertToRef {
		return fmt.Sprintf("%s: %s(utils.ValToRefSlice(r.%s))", targetPropertyName, funcName, targetPropertyName)
	} else if convertToVal {
		return fmt.Sprintf("%s: %s(utils.RefToValSlice(r.%s))", targetPropertyName, funcName, targetPropertyName)
	}

	return fmt.Sprintf("%s: %s(r.%s)", targetPropertyName, funcName, targetPropertyName)
}

func getConversionFunctionName(objectName string, sourcePackage string, targetPackage string) string {
	caser := cases.Title(language.English, cases.NoLower)

	return fmt.Sprintf("%s%sTo%s",
		caser.String(objectName),
		caser.String(sourcePackage),
		caser.String(targetPackage))
}

func getConversionFunctionNameToSlice(objectName string, sourcePackage string, targetPackage string) string {
	return fmt.Sprintf("%sSlice", getConversionFunctionName(objectName, sourcePackage, targetPackage))
}

func getConversionFunction(name string, template string, p ConvertFunctionParams) string {
	out := csgen.ExecuteTemplate[ConvertFunctionParams](name, template, p)

	return out
}

// LoadManifest loads the manifest file and returns a slice of ObjectMap structs.
func LoadManifest(path string) Manifest {
	log.Printf("Loading manifest file: %s\n", path)
	yfile, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var manifest Manifest
	err = yaml.Unmarshal(yfile, &manifest)
	if err != nil {
		log.Fatal(err)
	}

	if len(manifest.GeneratorPackage) == 0 {
		manifest.GeneratorPackage = csgen.InferPackageFromOutputPath(manifest.GeneratorPath)
	}

	for i, m := range manifest.ObjectMaps {
		sourcePath := filepath.Join(manifest.ProjectRoot, m.SourcePath)
		targetPath := filepath.Join(manifest.ProjectRoot, m.TargetPath)

		manifest.ObjectMaps[i].SourceObjects, _ = csgen.GetStructs(sourcePath)
		manifest.ObjectMaps[i].TargetObjects, _ = csgen.GetStructs(targetPath)
	}

	return manifest
}

var functionTemplate = `
// {{.FunctionName}} converts the source object to the target object.
func {{.FunctionName}}(r {{.FQSourceObjectName}}) {{.FQTargetObjectName}} {
	out := {{.FQTargetObjectName}} { 
		{{range .Assignments}}{{.}},
		{{end}}
	}

	return out
}

`

var functionSliceTemplate = `
// {{.FunctionNameForSlice}} converts the source object slice to the target object slice.
func {{.FunctionNameForSlice}} (r []*{{.FQSourceObjectName}}) []*{{.FQTargetObjectName}} {
	out := []*{{.FQTargetObjectName}}{}

	for _, v := range r {
		out = append(out, utils.ValToRef({{.FunctionName}}(*v)))
	}

	return out
}

`
