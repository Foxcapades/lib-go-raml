{{define "overlay-interface" -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

// Overlay represents the contents of a RAML API overlay.
type Overlay interface {
	yaml.Unmarshaler
	yaml.Marshaler
	{{template "hasTitle" "APISpec"}}
	{{template "hasAnnotations" "APISpec"}}
	{{template "hasAnnotationTypes" "APISpec"}}
	{{template "hasDescription" "APISpec"}}
	{{template "hasFacets" "APISpec"}}
	{{template "hasResourceTypes" "APISpec"}}
	{{template "hasSecuritySchemes" "APISpec"}}
	{{template "hasTraits" "APISpec"}}
	{{template "hasTypes" "APISpec"}}
	{{template "hasDocumentation" "APISpec"}}

	BaseURIParameters() UntypedMap

	Resources() UntypedMap
}
{{- end}}