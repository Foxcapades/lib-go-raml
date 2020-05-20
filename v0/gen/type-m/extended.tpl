{{- /* gotype: github.com/Foxcapades/lib-go-raml-types/v0/tools/gen/type.extTypeProps */ -}}
package raml
{{ define "extended" }}
import (
{{if not (eq .Name "Array") -}}
	"github.com/Foxcapades/goop/v1/pkg/option"
{{- end}}
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/assign"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// New{{ .Name }}Type returns a new internal implementation
// of the raml.{{ .Name }}Type interface.
//
// Generated @ {{ .Time }}
func New{{ .Name }}Type() *{{ .Name }}Type {
	out := &{{.Name}}Type{
		examples: New{{.Name}}ExampleMap(),
	}
	{{if eq .Name "Object" -}}
		{{template "object-constructor" $}}
	{{- else if eq .Name "Array" -}}
		{{template "array-constructor" $}}
	{{- else if eq .Name "String" -}}
		{{template "string-constructor" $}}
	{{- else if eq .Name "Number" -}}
		{{template "number-constructor" $}}
	{{- else if eq .Name "Integer" -}}
		{{template "integer-constructor" $}}
	{{- else if eq .Name "File" -}}
		{{template "file-constructor" $}}
	{{- else if eq .Name "Datetime" -}}
		{{template "datetime-constructor" $}}
	{{- end }}
	out.ExtendedDataType = NewExtendedDataType(rmeta.Type{{.Name}}, out)

	return out
}

// {{.Name}}Type is a default generated implementation of
// the raml.{{.Name}}Type interface
//
// Generated @ {{ .Time }}
type {{.Name}}Type struct {
	*ExtendedDataType

	def      {{if .IsDefPtr}}*{{end}}{{.DefType}}
	example  raml.{{.Name}}Example
	examples raml.{{.Name}}ExampleMap
	enum     []{{.EnumType}}
	{{if eq .Name "Object" -}}
		{{template "object-props" $}}
	{{- else if eq .Name "Array" -}}
		{{template "array-props" $}}
	{{- else if eq .Name "String" -}}
		{{template "string-props" $}}
	{{- else if eq .Name "Number" -}}
		{{template "number-props" $}}
	{{- else if eq .Name "Integer" -}}
		{{template "integer-props" $}}
	{{- else if eq .Name "File" -}}
		{{template "file-props" $}}
	{{- else if eq .Name "Datetime" -}}
		{{template "datetime-props" $}}
	{{- end }}
}

func (o *{{.Name}}Type) SetType(s string) raml.{{.Name}}Type {
	o.schema = s
	return o
}

func (o *{{.Name}}Type) Default() {{if .DefIsOpt}}option.{{end}}{{.DefTypeName}} {
	{{ if .DefIsOpt -}}
		return option.NewMaybe{{.DefTypeName}}(o.def)
	{{- else -}}
		return o.def
	{{- end}}
}

func (o *{{.Name}}Type) SetDefault(i {{.DefType}}) raml.{{.Name}}Type {
	o.def = {{if .IsDefPtr}}&{{end}}i
	return o
}

func (o *{{.Name}}Type) UnsetDefault() raml.{{.Name}}Type {
	o.def = nil
	return o
}

func (o *{{.Name}}Type) Example() raml.{{.Name}}Example {
	return o.example
}

func (o *{{.Name}}Type) SetExample(ex raml.{{.Name}}Example) raml.{{.Name}}Type {
	o.example = ex
	return o
}

func (o *{{.Name}}Type) UnsetExample() raml.{{.Name}}Type {
	o.example = nil
	return o
}

func (o *{{.Name}}Type) Examples() raml.{{.Name}}ExampleMap {
	return o.examples
}

func (o *{{.Name}}Type) SetExamples(examples raml.{{.Name}}ExampleMap) raml.{{.Name}}Type {
	if examples == nil {
		return o.UnsetExamples()
	}
	o.examples = examples
	return o
}

func (o *{{.Name}}Type) UnsetExamples() raml.{{.Name}}Type {
	o.examples = New{{.Name}}ExampleMap()
	return o
}

func (o *{{.Name}}Type) SetDisplayName(s string) raml.{{.Name}}Type {
	o.displayName = &s
	return o
}

func (o *{{.Name}}Type) UnsetDisplayName() raml.{{.Name}}Type {
	o.displayName = nil
	return o
}

func (o *{{.Name}}Type) SetDescription(s string) raml.{{.Name}}Type {
	o.description = &s
	return o
}

func (o *{{.Name}}Type) UnsetDescription() raml.{{.Name}}Type {
	o.description = nil
	return o
}

func (o *{{.Name}}Type) SetAnnotations(annotations raml.AnnotationMap) raml.{{.Name}}Type {
	if annotations == nil {
		return o.UnsetAnnotations()
	}
	o.hasAnnotations.mp = annotations
	return o
}

func (o *{{.Name}}Type) UnsetAnnotations() raml.{{.Name}}Type {
	o.hasAnnotations.mp = NewAnnotationMap()
	return o
}

func (o *{{.Name}}Type) SetFacetDefinitions(facets raml.FacetMap) raml.{{.Name}}Type {
	if facets == nil {
		return o.UnsetFacetDefinitions()
	}
	o.facets = facets
	return o
}

func (o *{{.Name}}Type) UnsetFacetDefinitions() raml.{{.Name}}Type {
	o.facets = NewFacetMap()
	return o
}

func (o *{{.Name}}Type) SetXml(x raml.Xml) raml.{{.Name}}Type {
	o.xml = x
	return o
}

func (o *{{.Name}}Type) UnsetXml() raml.{{.Name}}Type {
	o.xml = nil
	return o
}

func (o *{{.Name}}Type) Enum() []{{.EnumType}} {
	return o.enum
}

func (o *{{.Name}}Type) SetEnum(i []{{.EnumType}}) raml.{{.Name}}Type {
	o.enum = i
	return o
}

func (o *{{.Name}}Type) UnsetEnum() raml.{{.Name}}Type {
	o.enum = nil
	return o
}

func (o *{{.Name}}Type) SetExtraFacets(facets raml.AnyMap) raml.{{.Name}}Type {
	if facets == nil {
		return o.UnsetExtraFacets()
	}
	o.hasExtra.mp = facets
	return o
}

func (o *{{.Name}}Type) UnsetExtraFacets() raml.{{.Name}}Type {
	o.hasExtra.mp = NewAnyMap()
	return o
}

func (o *{{.Name}}Type) SetRequired(b bool) raml.{{.Name}}Type {
	o.required = b
	return o
}

{{if eq .Name "Object" -}}
	{{template "object-methods" $}}
{{- else if eq .Name "Array" -}}
	{{template "array-methods" $}}
{{- else if eq .Name "String" -}}
	{{template "string-methods" $}}
{{- else if eq .Name "Number" -}}
	{{template "number-methods" $}}
{{- else if eq .Name "Integer" -}}
	{{template "integer-methods" $}}
{{- else if eq .Name "File" -}}
	{{template "file-methods" $}}
{{- else if eq .Name "Datetime" -}}
	{{template "datetime-methods" $}}
{{- end -}}

func (o *{{.Name}}Type) marshal(out raml.AnyMap) error {
	logrus.Trace("internal.{{.Name}}Type.marshal")
	out.PutNonNil(rmeta.KeyDefault, o.def)

	if err := o.ExtendedDataType.marshal(out); err != nil {
		return err
	}
	{{if eq .Name "Object" -}}
	{{template "object-marshal" $}}
	{{- else if eq .Name "Array" -}}
	{{template "array-marshal" $}}
	{{- else if eq .Name "String" -}}
	{{template "string-marshal" $}}
	{{- else if eq .Name "Number" -}}
	{{template "number-marshal" $}}
	{{- else if eq .Name "Integer" -}}
	{{template "integer-marshal" $}}
	{{- else if eq .Name "File" -}}
	{{template "file-marshal" $}}
	{{- else if eq .Name "Datetime" -}}
	{{template "datetime-marshal" $}}
	{{- end}}
	out.PutNonNil(rmeta.KeyEnum, o.enum).
		PutNonNil(rmeta.KeyExample, o.example)

	if o.examples.Len() > 0 {
		out.PutNonNil(rmeta.KeyExamples, o.examples)
	}

	return nil
}

func (o *{{.Name}}Type) assign(key, val *yaml.Node) error {
	logrus.Trace("internal.{{.Name}}Type.assign")
	switch key.Value {
	case rmeta.KeyExample:
		if ex, err := ExampleSortingHat(o.kind); err != nil {
			return err
		} else if err := ex.UnmarshalRAML(val); err != nil {
			return err
		} else {
			o.example = ex.(raml.{{.Name}}Example)
		}
		return nil
	case rmeta.KeyExamples:
		return o.examples.UnmarshalRAML(val)
	case rmeta.KeyEnum:
		return xyml.ForEachList(val, func(cur *yaml.Node) error {
			{{if eq .DefTypeName "Bool" "Float64" "Int64" "String" -}}
			if val, err := xyml.To{{.DefTypeName}}(cur); err != nil {
				return err
			} else {
				o.enum = append(o.enum, val)
			}
			{{- else if eq .DefTypeName "Untyped" -}}
			o.enum = append(o.enum, val)
			{{- end}}

			return nil
		})
		return nil
	case rmeta.KeyRequired:
		return assign.AsBool(val, &o.required)
	}
	{{if eq .Name "Object" -}}
		{{template "object-assign" $}}
	{{- else if eq .Name "Array" -}}
		{{template "array-assign" $}}
	{{- else if eq .Name "String" -}}
		{{template "string-assign" $}}
	{{- else if eq .Name "Number" -}}
		{{template "number-assign" $}}
	{{- else if eq .Name "Integer" -}}
		{{template "integer-assign" $}}
	{{- else if eq .Name "File" -}}
		{{template "file-assign" $}}
	{{- else if eq .Name "Datetime" -}}
		{{template "datetime-assign" $}}
	{{- end}}
	return o.ExtendedDataType.assign(key, val)
}
{{end}}