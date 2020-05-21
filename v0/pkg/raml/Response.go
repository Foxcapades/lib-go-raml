package raml

import (
	"github.com/Foxcapades/goop/v1/pkg/option"
	"gopkg.in/yaml.v3"
)

type Response interface {
	yaml.Unmarshaler
	yaml.Marshaler

	// Description returns an option of the contents of the `description` facet
	// applied to the current RAML element.
	//
	// If the element does not have a description applied, the returned option
	// will be empty.
	Description() option.String

	// SetDescription sets the value of the description property for the current
	// RAML element.
	SetDescription(desc string) Response

	// UnsetDescription removes the description property from the current RAML
	// element.
	UnsetDescription() Response

	// Annotations returns a map of the annotations applied to the current
	// RAML element.
	Annotations() AnnotationMap

	// SetAnnotations replaces all annotations applied to the current RAML element
	// with those defined in the given map.
	//
	// Passing this method a nil value is equivalent to calling UnsetAnnotations.
	SetAnnotations(annotations AnnotationMap) Response

	// UnsetAnnotations removes all annotations applied to the current RAML
	// element.
	UnsetAnnotations() Response

	// Headers returns the map of headers defined on the current RAML element.
	Headers() DataTypeMap

	// SetHeaders replaces the map of headers for the current RAML element with
	// the given value.
	//
	// Passing this method a nil value is equivalent to calling UnsetHeaders.
	SetHeaders(headers DataTypeMap) Response

	// UnsetHeaders removes all headers from the current RAML element.
	UnsetHeaders() Response

	Body() Body

	SetBody(body Body) Response

	UnsetBody() Response
}
