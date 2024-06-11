package gen

import (
	"github.com/cscoding21/csgen"
)

// Manifest strongly typed respresentation of the manifest file.
type Manifest struct {
	ProjectRoot      string      `yaml:"project_root"`
	GeneratorPath    string      `yaml:"generator_path"`
	GeneratorPackage string      `yaml:"generator_package"`
	ObjectMaps       []ObjectMap `yaml:"maps"`
}

// MapOverride struct that represents an explicitly specified object mapping, overriding the defailt name matching
type MapOverride struct {
	SourceName string `yaml:"source_name"`
	TargetName string `yaml:"target_name"`
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
