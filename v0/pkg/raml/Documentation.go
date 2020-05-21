package raml

type Documentation interface {
	Marshaler
	Unmarshaler

	// Title returns the value of the title property set on the current RAML
	// element.
	Title() string

	// SetTitle sets the value of the title property on the current RAML element.
	SetTitle(title string) Documentation

	// ExtraFacets returns a mutable map of extra/custom facets applied to the
	// current RAML element.
	ExtraFacets() AnyMap
	// Content returns the value of the content property defined on the current
	// RAML element.
	Content() string

	// SetContent sets the value of the content property on the current RAML
	// element.
	SetContent(content string) Documentation
}
