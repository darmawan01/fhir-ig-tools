package models

import "time"

type Meta struct {
	VersionID   string    `json:"versionId"`
	LastUpdated time.Time `json:"lastUpdated"`
}

type Extension struct {
	URL          string `json:"url"`
	ValueString  string `json:"valueString,omitempty"`
	ValueCode    string `json:"valueCode,omitempty"`
	ValueInteger int    `json:"valueInteger,omitempty"`
}

type Telcom struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

type Contact struct {
	Telecom []Telcom `json:"telecom"`
}

type Mapping struct {
	Identity string `json:"identity"`
	URI      string `json:"uri"`
	Name     string `json:"name"`
}

type Discriminator struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type Slicing struct {
	Discriminator []Discriminator `json:"discriminator"`
	Rules         string          `json:"rules"`
}

type Type struct {
	Code    string   `json:"code"`
	Profile []string `json:"profile"`
}

type Binding struct {
	Strength string `json:"strength"`
	ValueSet string `json:"valueSet"`
}

type Element struct {
	ID         string  `json:"id"`
	Path       string  `json:"path"`
	Slicing    Slicing `json:"slicing,omitempty"`
	Min        int     `json:"min,omitempty"`
	SliceName  string  `json:"sliceName,omitempty"`
	Max        string  `json:"max,omitempty"`
	Type       []Type  `json:"type,omitempty"`
	Short      string  `json:"short,omitempty"`
	Definition string  `json:"definition,omitempty"`
	Binding    Binding `json:"binding,omitempty"`
	FixedURI   string  `json:"fixedUri,omitempty"`
	Comment    string  `json:"comment,omitempty"`
}

type Differential struct {
	Element []Element `json:"element"`
}

type BaseModel struct {
	ResourceType   string       `json:"resourceType"`
	ID             string       `json:"id"`
	Meta           Meta         `json:"meta"`
	Extension      []Extension  `json:"extension"`
	URL            string       `json:"url"`
	Version        string       `json:"version"`
	Name           string       `json:"name"`
	Status         string       `json:"status"`
	Date           time.Time    `json:"date"`
	Publisher      string       `json:"publisher"`
	Contact        []Contact    `json:"contact"`
	Description    string       `json:"description"`
	Purpose        string       `json:"purpose"`
	FhirVersion    string       `json:"fhirVersion"`
	Mapping        []Mapping    `json:"mapping"`
	Kind           string       `json:"kind"`
	Abstract       bool         `json:"abstract"`
	Type           string       `json:"type"`
	BaseDefinition string       `json:"baseDefinition"`
	Derivation     string       `json:"derivation"`
	Differential   Differential `json:"differential"`
}
