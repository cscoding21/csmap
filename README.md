
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
    <a href="https://discord.gg/BjV88Bys" alt="Discord">
        <img src="https://img.shields.io/discord/1196192809120710779" /></a>&nbsp;
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

    go install github.com/cscoding21/csval

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
    # Test of 2 objects with different names 
    - name: "source_data1"
      source_path: "tests/diffnames/source_data1.go"
      target_path: "tests/diffnames/target_data1.go"
      # imports: 
      # - "github.com/cscoding21/csgen"
      map_overrides:
      - source_name: "TestSource"
        target_name: "TestTarget"
      - source_name: "LocationSource"
        target_name: "LocationTarget"