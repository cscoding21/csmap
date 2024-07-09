
<p align="center"><img src="https://github.com/cscoding21/cscoding/blob/main/assets/csc-banner.png?raw=true" width=728></p>

<p align="center">
    <a href="https://github.com/cscoding21/csmap"><img src="https://img.shields.io/badge/built_with-Go-29BEB0.svg?style=flat-square"></a>&nbsp;
    <a href="https://goreportcard.com/report/github.com/cscoding21/csmap"><img src="https://goreportcard.com/badge/github.com/cscoding21/csmap?style=flat-square"></a>&nbsp;
 <a href="https://pkg.go.dev/mod/github.com/cscoding21/csmap"><img src="https://pkg.go.dev/badge/mod/github.com/cscoding21/csmap"></a>&nbsp;
    <a href="https://github.com/cscoding21/csmap/" alt="Stars">
        <img src="https://img.shields.io/github/stars/cscoding21/csmap?color=0052FF&labelColor=090422" /></a>&nbsp;
    <a href="https://github.com/cscoding21/csmap/pulse" alt="Activity">
        <img src="https://img.shields.io/github/commit-activity/m/cscoding21/csmap?color=0052FF&labelColor=090422" /></a>
    <br />
    <!-- <a href="https://discord.gg/BjV88Bys" alt="Discord">
        <img src="https://img.shields.io/discord/1196192809120710779" /></a>&nbsp; -->
    <a href="https://www.youtube.com/@CommonSenseCoding-ge5dn" alt="YouTube">
        <img src="https://img.shields.io/badge/youtube-watch_videos-red.svg?color=0052FF&labelColor=090422&logo=youtube" /></a>&nbsp;
    <a href="https://twitter.com/cscoding21" alt="YouTube">
        <img src="https://img.shields.io/twitter/follow/cscoding21" /></a>&nbsp;
</p>



# CSMap
CSVal is a Golang package that generates functions that map values from one object to another.  The primary use-case is to allow objects from the one tier (e.g. [DTOs](https://en.wikipedia.org/wiki/Data_transfer_object)) to be mapped to objects at the business/service layer.  Because the mapping code is generated pre-compilation, reflection is not required.

## Installation
To install the CSMap package...run the following command:

    go get github.com/cscoding21/csmap

Additionally, the runner can be installed on the target machine.  This is a wrapper executable that accepts a file location for the manifest file and processes the defined maps.

    go install github.com/cscoding21/csmap

## How it Works
CSMap uses Golang's AST package (abstracted by [CSGen](http://github.com/cscoding21/csgen)) to create mapping functions between two different objects.  The developer defines the mapping rules in the manifest file and the generator will create the mapping functions.  The developer can then use the mapping functions to map objects from one tier to another.

The manifest file contains an array of files to map.  Each item contains two files...a source and a destination, along with some optional additional data to specify behavior.  

CSMap reads the structs contained in each of the files and looks for definitions with the same name.  Alternately, if an override is specified, structs with different names can be mapped.

The map generator creates one file for each source/destination pair in the manifest.  Within the file, there will be two functions for each mapped struct: one for a single object and the other for a slice.  For each field in the source...a field with the same name in the destination will be mapped __provided they are of a like data type__.  The mapping code will attempt to handle the following edge-cases for individual fields:

- Mapping of an address/pointer to a value field and vice-versa
- Mapping by calling a separate mapping function created by the mapper __in the same file__
- Mapping of a type to a compatible type (e.g. string to custom def of type string)

## Usage
Once the mapping functions have been generated, they can be called in the following manner:

    # Define the source object
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

    # Calling the generated function maps all of the connected fields
	testDest := tests.ActivityPkg1ToPkg2(testSource)


## Manifest File
The manifest file is a YAML file containing configuration parameters as well as a list of object maps to process.  The developer can pass in the name and location of the manifest file.  If nothing is specified, the generator will look for a file called __csmap.yaml__ in the current directory (from where the process is being run).

Below is a sample manifest file:

    # The root of the project
    project_root: "/home/jeph/projects/cscoding21/csmap"

    # The relative path to the root where generated output will go
    generator_path: "tests"

    # An optional value for the package name of generated files.  If not specified, the name of the directory containing the generated output will be used.
    # generator_package: "tests"

    # source and destination paths are relative to project root
    maps:
    # A name for the map.  This is used as the root of the file name
    - name: "source_data1"

      # A path to the file containing the source object
      source_path: "tests/diffnames/source_data1.go"

      # A path to the file containing the target object
      target_path: "tests/diffnames/target_data1.go"

      # A list of imports for the generated file.  The process calls __goimports__ during file generation, so this shouldn't ever be necessary 
      # imports: 
      # - "github.com/cscoding21/csgen"

      # An array of object mapping names.  This is only necessary if the name of the source and target are different.  
      map_overrides:
      - source_name: "TestSource"
        target_name: "TestTarget"
      - source_name: "LocationSource"
        target_name: "LocationTarget"