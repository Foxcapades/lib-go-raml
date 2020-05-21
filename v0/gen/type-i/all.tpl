package x

type name interface {
{{define "all"}}
	// SetType sets the parent type name, this does not change
	// the underlying kind of the DataType.
	SetType(string) {{.Name}}Type

	{{if .DefIsOpt -}}
	// Default returns an option of the current default value
	// which will be empty if the value is not set.
	{{- else -}}
	// Default returns the current default value or nil of no
	// default value is set.
	{{- end}}
	Default() {{if .DefIsOpt}}option.{{end}}{{.DefTypeName}}

	// SetDefault sets the default value for the current
	// DataType definition.
	SetDefault({{.DefType}}) {{.Name}}Type

	// UnsetDefault removes the default value definition from
	// the current DataType definition.
	//
	// If no default was previously set, this method does
	// nothing.
	UnsetDefault() {{.Name}}Type

	// Example returns the singular example value assigned to
	// the current DataType definition.
	Example() {{.Name}}Example

	// SetExample sets the singular example value assigned to
	// the current DataType definition.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExample.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExample(ex {{.Name}}Example) {{.Name}}Type

	// UnsetExample removes the singular example value from
	// the current DataType definition.
	UnsetExample() {{.Name}}Type

	// Examples returns a mutable map of examples for the
	// current DataType definition.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	Examples() {{.Name}}ExampleMap

	// SetExamples replaces the current DataType definition's
	// example map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// ass calling UnsetExamples.
	//
	// Note: It is illegal in RAML to have both the "example"
	// and "examples" keys set.
	SetExamples(examples {{.Name}}ExampleMap) {{.Name}}Type

	// UnsetExamples clears the current DataType definition's
	// example map.
	UnsetExamples() {{.Name}}Type

	// SetDisplayName sets the current DataType definition's
	// displayName value.
	SetDisplayName(string) {{.Name}}Type

	// UnsetDisplayName removes the current DataType
	// definition's displayName value.
	UnsetDisplayName() {{.Name}}Type

	// SetDescription sets the current DataType definition's
	// description value.
	SetDescription(string) {{.Name}}Type

	// UnsetDescription removes the current DataType
	// definition's description value.
	UnsetDescription() {{.Name}}Type

	// SetAnnotations replaces the current DataType
	// definition's annotation map with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) {{.Name}}Type

	// UnsetAnnotations clears the current DataType
	// definition's applied annotation map.
	UnsetAnnotations() {{.Name}}Type

	// SetFacetDefinitions replaces the current DataType
	// definition's custom facet definitions with the given
	// value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetFacets.
	SetFacetDefinitions(facets FacetMap) {{.Name}}Type

	// UnsetFacetDefinitions removes all custom facets defined
	// on the current DataType definition.
	UnsetFacetDefinitions() {{.Name}}Type

	// SetXML sets the current DataType definition's xml serialization settings
	// object to the given value.
	//
	// Passing this method a nil value is effectively the same as calling
	// UnsetXML.
	SetXML(XML) {{.Name}}Type

	// UnsetXML removes the xml serialization settings object from the current
	// DataType definition.
	UnsetXML() {{.Name}}Type

	// Enum returns a slice of the enum values assigned to the
	// current DataType definition.
	Enum() []{{.EnumType}}

	// SetEnum replaces the current DataType definition's enum
	// slice with the given value.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetEnum.
	SetEnum([]{{.EnumType}}) {{.Name}}Type

	// UnsetEnum clears the enum values from the current
	// DataType definition.
	UnsetEnum() {{.Name}}Type

	// SetExtraFacets replaces the facets applied to the
	// current DataType definition with the given values.
	//
	// Passing this method a nil value is effectively the same
	// as calling UnsetExtraFacets.
	SetExtraFacets(facets AnyMap) {{.Name}}Type

	// UnsetExtraFacets removes all extra facets from the
	// current DataType definition.
	UnsetExtraFacets() {{.Name}}Type

	SetRequired(bool) {{.Name}}Type
{{end}}
}