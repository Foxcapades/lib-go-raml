package raml

import "gopkg.in/yaml.v2"

// Library represents the contents of a RAML library file.
type Library interface {
	yaml.Unmarshaler
	yaml.Marshaler

	hasAnnotations
	hasAnnotationTypes
	hasFacets
	hasImports
	hasResourceTypes
	hasSecuritySchemes
	hasTraits
	hasTypes
	hasUsage
}
