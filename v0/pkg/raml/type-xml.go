package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

type Xml interface {
	IsAttribute() option.Bool
	SetIsAttribute(bool) Xml
	UnsetIsAttribute() Xml

	IsWrapped() option.Bool
	SetIsWrapped(bool) Xml
	UnsetIsWrapped() Xml

	Name() option.String
	SetName(string) Xml
	UnsetName() Xml

	Namespace() option.String
	SetNamespace(string) Xml
	UnsetNamespace() Xml

	Prefix() option.String
	SetPrefix(string) Xml
	UnsetPrefix() Xml
}
