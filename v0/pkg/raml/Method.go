package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

type Method interface {
	yaml.Unmarshaler
	yaml.Marshaler

	// DisplayName returns an option of the contents of the `displayName` facet
	// applied to the current RAML element.
	//
	// If the element does not have a display name applied, the returned option
	// will be empty.
	DisplayName() option.String

	// SetDisplayName sets the value of the displayName property for the current
	// RAML element.
	SetDisplayName(name string) Method

	// UnsetDisplayName removes the displayName property from the current RAML
	// element.
	UnsetDisplayName() Method

	// Description returns an option of the contents of the `description` facet
	// applied to the current RAML element.
	//
	// If the element does not have a description applied, the returned option
	// will be empty.
	Description() option.String

	// SetDescription sets the value of the description property for the current
	// RAML element.
	SetDescription(desc string) Method

	// UnsetDescription removes the description property from the current RAML
	// element.
	UnsetDescription() Method

	// Annotations returns a map of the annotations applied to the current
	// RAML element.
	Annotations() AnnotationMap

	// SetAnnotations replaces all annotations applied to the current RAML element
	// with those defined in the given map.
	//
	// Passing this method a nil value is equivalent to calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) Method

	// UnsetAnnotations removes all annotations applied to the current RAML
	// element.
	UnsetAnnotations() Method

	// Headers returns the map of headers defined on the current RAML element.
	Headers() DataTypeMap

	// SetHeaders replaces the map of headers for the current RAML element with
	// the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetHeaders.
	SetHeaders(headers DataTypeMap) Method

	// UnsetHeaders removes all headers from the current RAML element.
	UnsetHeaders() Method

	// Protocols returns the values of the protocol array set on the current RAML
	// element.
	Protocols() []string

	// SetProtocols sets the protocol array for the current RAML element to the
	// given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetProtocols.
	SetProtocols(protocols []string) Method

	// UnsetProtocols removes the protocols property and all its values from the
	// current RAML element.
	UnsetProtocols() Method

	// SecuredBy returns the value of the securedBy array defined on the current
	// RAML element.
	SecuredBy() []string

	// SetSecuredBy sets the securedBy array for the current RAML element.
	//
	// Passing this method a nil value is equivalent to calling UnsetSecuredBy.
	SetSecuredBy(securedBy []string) Method

	// UnsetSecuredBy removes the securedBy property and all its values from the
	// current RAML element.
	UnsetSecuredBy() Method

	Body() Body

	SetBody(body Body) Method

	UnsetBody() Method
}
