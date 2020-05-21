{{define "api-interface" -}}
package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
	"io"
)

// APISpec represents a full root RAML document.
type APISpec interface {
	yaml.Marshaler
	yaml.Unmarshaler
	{{template "hasTitle" "APISpec"}}
	{{template "hasAnnotations" "APISpec"}}
	{{template "hasAnnotationTypes" "APISpec"}}
	{{template "hasDescription" "APISpec"}}
	{{template "hasFacets" "APISpec"}}
	{{template "hasImports" "APISpec"}}
	{{template "hasResourceTypes" "APISpec"}}
	{{template "hasSecuritySchemes" "APISpec"}}
	{{template "hasTraits" "APISpec"}}
	{{template "hasTypes" "APISpec"}}
	{{template "hasDocumentation" "APISpec"}}
	{{template "hasProtocols" "APISpec"}}
	{{template "hasSecuredBy" "APISpec"}}
	// Version returns an option of the declared version for this API spec.
	//
	// If no version has been declared, the returned option will be empty.
	Version() option.String

	// SetVersion sets the version value for this API spec.
	SetVersion(version string) APISpec

	// UnsetVersion removes the version declaration from this API spec.
	UnsetVersion() APISpec

	// BaseURI returns an option of the baseUri value declared on the current
	// RAML element.
	//
	// If no baseUri value has been set, the returned option will be empty.
	BaseURI() option.String

	// SetBaseURI sets the baseUri value for the current RAML element.
	SetBaseURI(baseUri string) APISpec

	// UnsetBaseURI removes the baseUri value from the current RAML element.
	UnsetBaseURI() APISpec

	// MediaTypes returns the values of the mediaTypes array set on the current
	// RAML element.
	//
	// When parsed, if the RAML document has a single string value set for the
	// mediaTypes property, this slice will be populated with that single value.
	//
	// When rendered, if this slice contains a single value, it will be rendered
	// as a single string rather than an array.
	MediaTypes() []string

	// SetMediaTypes sets the mediaTypes array for the current RAML element to the
	// given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetMediaTypes.
	SetMediaTypes(mediaTypes []string) APISpec

	// UnsetMediaTypes removes the mediaTypes property and all its values from the
	// current RAML element.
	UnsetMediaTypes() APISpec

	BaseURIParameters() UntypedMap

	Resources() UntypedMap

	WriteRAML(w io.Writer) error
}
{{- end}}