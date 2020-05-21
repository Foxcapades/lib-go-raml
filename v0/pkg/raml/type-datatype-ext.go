package raml

type ExtendedDataType interface {
	hasAnnotations
	hasDescription
	displayed

	DataType

	// FacetDefinitions returns a map of the custom facets
	// defined on the current DataType definition.
	FacetDefinitions() FacetMap

	// XML returns the xml serialization settings for the
	// current DataType definition, or nil if no settings have
	// been applied.
	XML() XML
}
