package main

import "encoding/json"

type DrafterResult struct {
	AST Blueprint
	Sourcemap BlueprintSourceMap
	Error Annotation
	Warnings []Annotation
}

type Blueprint struct {
	Version string `json:"_version"`
	Metadata []Metadata
	Name string
	Description string
	Element string
	ResourceGroup []ResourceGroup
	Content []Element
}

type Metadata struct {
	Name string
	Value string
}

type ResourceGroup struct {
	Name string
	Description string
	Resources []Resource
}

type Resource struct {
	Element string
	Name string
	Description string
	UriTemplate string
	Model Payload
	Parameters []Parameter
	Actions []Action
	Content []DataStructure
}

type Payload struct {
	Reference Reference
	Name string
	Description string
	Headers []Header
	Body string
	Schema string
	Content []DataStructure | Asset
}

type Parameter struct {
	Name string
	Description string
	Type string
	Required bool
	Default string
	Example string
	Values []ParameterValue
}

type ParameterValue struct {
	Value string
}

type Action struct {
	Name string
	Description string
	Method string
	Parameters []Parameter
	Attributes ActionAttributes
	Relation string
	UriTemplate string
	Content []DataStructure
	Examples []TransactionExample
}

type ActionAttributes struct {
	Relation string
	UriTemplate string
}

type TransactionExample struct {
	Name string
	Description string
	Requests []Payload
	Responses []Payload
}

type Reference struct {
	Id string
}

type Header struct {
	Name string
	Value string
}

type Annotation struct {
	Message string
	Code int
	Location []Location
}

type Location struct {
	Index int
	Length int
}

type Asset struct {
	Element string
	Attributes AssetAttributes
	Content string
}

type AssetAttributes struct {
	Role string
}

type DataStructure struct {
	Element string
	Content []DataStructureElement
}

type DataStructureElement struct {
	Element string
	Meta DataStructureElementMetadata
	Content []DataStructureElementType
}

type DataStructureElementMetadata struct {
	Id string
	Description string
}

type DataStructureElementType struct {
	Element string
	Meta DataStructureElementTypeMetadata
	Attributes DataStructureElementTypeAttributes
	Content DataStructureElementTypeProperty
}

type DataStructureElementTypeMetadata struct {
	Description string
}

type DataStructureElementTypeAttributes struct {
	TypeAttribute []string
	Samples []interface{}
	Default interface{}
}

type DataStructureElementTypeProperty struct {
	Key DataStructureElementTypePropertyKey
	Value DataStructureElementTypePropertyValue
}

type DataStructureElementTypePropertyKey struct {
	Element string
	Content string
}

type DataStructureElementTypePropertyValue struct {
	Element string
	Content string
}

type Element struct {
	Element string
	Attributes ElementAttributes
	Content interface{}
	rawContent []byte
}

func (e *Element) getRawContent() []byte {
	if e.rawContent == nil {
		e.rawContent, _ = json.Marshal(e.Content)
	}
	return e.rawContent
}

func (e *Element) ContentAsCopyElement() string {
	var res string
	_ := json.Unmarshal(e.getRawContent(), &res)
	return res
}

func (e *Element) ContentAsCategoryElement() []Element {
	var res []Element
	_ := json.Unmarshal(e.getRawContent(), &res)
	return res
}

func (e *Element) ContentAsDataStructureElement() DataStructure {
	var res DataStructure
	_ := json.Unmarshal(e.getRawContent(), &res)
	return res
}

func (e *Element) ContentAsResourceElement() Resource {
	var res Resource
	_ := json.Unmarshal(e.getRawContent(), &res)
	return res
}

type ElementAttributes struct {
	Name string
}

type BlueprintSourceMap interface{}
