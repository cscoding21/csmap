package gen

import (
	"github.com/cscoding21/csgen"
)

// Manifest strongly typed respresentation of the manifest file.
type Manifest struct {
	ProjectRoot      string         `yaml:"project_root"`
	GeneratorPath    string         `yaml:"generator_path"`
	GeneratorPackage string         `yaml:"generator_package"`
	ObjectMaps       []ObjectMap    `yaml:"maps"`
	CommonStructs    []CommonStruct `yaml:"common_structs"`
}

// MapOverride struct that represents an explicitly specified object mapping, overriding the defailt name matching
type MapOverride struct {
	SourceName string `yaml:"source_name"`
	TargetName string `yaml:"target_name"`
}

// CommonStruct a container for structs that have global application when creating maps
type CommonStruct struct {
	PackageName string       `yaml:"package_name"`
	StructName  string       `yaml:"struct_name"`
	Path        string       `yaml:"path"`
	Struct      csgen.Struct `yaml:"struct"`
}

// GetCommonStruct return a common struct by name if it exists and nil otherwise
func (m *Manifest) GetCommonStruct(name string) *CommonStruct {
	for _, s := range m.CommonStructs {
		if s.StructName == name {
			return &s
		}
	}

	return nil
}

// ObjectMap is a struct that contains the source and destination paths for a single mapping.
type ObjectMap struct {
	Name          string `yaml:"name"`
	SourcePath    string `yaml:"source_path"`
	SourceObjects []csgen.Struct
	TargetPath    string `yaml:"target_path"`
	TargetObjects []csgen.Struct
	Imports       []string      `yaml:"imports"`
	MapOverrides  []MapOverride `yaml:"map_overrides"`
}

// ConvertFunctionParams is a struct that contains the parameters that feed into the template.
type ConvertFunctionParams struct {
	ConverterPackage     string
	FunctionName         string
	FunctionNameForSlice string
	SourcePackage        string
	SourceObjectName     string
	FQSourceObjectName   string
	TargetPackage        string
	TargetObjectName     string
	FQTargetObjectName   string
	Assignments          []string
	Imports              []string
}

func getConversionFunctionParams(
	converterPackage string,
	targetObjectName string,
	targetPackageName string,
	sourceObjectName string,
	sourcePackageName string,
	imports []string,
	assignments []string) ConvertFunctionParams {

	params := ConvertFunctionParams{
		ConverterPackage:     converterPackage,
		FunctionName:         getConversionFunctionName(sourceObjectName, sourcePackageName, targetPackageName),
		FunctionNameForSlice: getConversionFunctionNameToSlice(sourceObjectName, sourcePackageName, targetPackageName),
		SourcePackage:        sourcePackageName,
		SourceObjectName:     sourceObjectName,
		FQSourceObjectName:   getFQObjectName(converterPackage, sourcePackageName, sourceObjectName),
		TargetPackage:        targetPackageName,
		TargetObjectName:     targetObjectName,
		FQTargetObjectName:   getFQObjectName(converterPackage, targetPackageName, targetObjectName),
		Imports:              imports,
		Assignments:          assignments,
	}

	return params
}
