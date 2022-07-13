# Fhir Implementation Guide Converter From Excel to FHIR Resource

## Parameters
- `-template` - Path to excel that contain Resource IG
- `-resource` - Path to json file of Resource
- `-out` - Out Flie Name for result

## Build

```
go build -o fhir-ig-converter
```

## Usage
- `./fhir-ig-converter -template path-to-ig-excel.xlsx -resource path-to-fhir-resource.json`

## Run without build
- `go run main.go -template path-to-ig-excel.xlsx -resource path-to-fhir-resource.json`
