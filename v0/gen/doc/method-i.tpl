{{define "method-interface" -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

type Method interface {
	yaml.Unmarshaler
	yaml.Marshaler
	{{template "hasDisplayName" "Method"}}
	{{template "hasDescription" "Method"}}
	{{template "hasAnnotations" "Method"}}
	{{template "hasHeaders"     "Method"}}
	{{template "hasProtocols"   "Method"}}
	{{template "hasSecuredBy"   "Method"}}
	{{template "hasBody"        "Method"}}
}
{{- end}}