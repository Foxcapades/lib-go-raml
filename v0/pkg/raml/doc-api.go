package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

type ApiSpec interface {
	yaml.Marshaler
	yaml.Unmarshaler

	hasAnnotations
	hasAnnotationTypes
	hasDescription
	hasFacets
	hasImports
	hasResourceTypes
	hasSecuritySchemes
	hasTraits
	hasTypes

	Title() string
	Version() option.String
	BaseUri() option.String

	Protocols() []string
	MediaTypes() []string
	Documentation() []interface{}
	SecuredBy() []string
	BaseUriParameters() UntypedMap

	Resources() UntypedMap
}
