{{define "response-interface" -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

type Response interface {
	yaml.Unmarshaler
	yaml.Marshaler
	{{template "hasDescription" "Response"}}
	{{template "hasAnnotations" "Response"}}
	{{template "hasHeaders"     "Response"}}
	{{template "hasBody"        "Response"}}
}
{{- end}}