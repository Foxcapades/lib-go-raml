package raml

type hasAnnotationTypes interface {
	AnnotationTypes() UntypedMap
}

type hasImports interface {
	// Uses returns the map of defined library imports for the current RAML
	// document.
	Uses() StringMap
}

type hasResourceTypes interface {
	ResourceTypes() UntypedMap
}

type hasSecuritySchemes interface {
	SecuritySchemes() UntypedMap
}

type hasTraits interface {
	Traits() UntypedMap
}

type hasTypes interface {
	// Types returns the value of the `types` or `schemas` property defined on
	// the current document.
	Types() DataTypeMap

	// Schemas is an alias for Types using the deprecated property name.
	Schemas() DataTypeMap
}
