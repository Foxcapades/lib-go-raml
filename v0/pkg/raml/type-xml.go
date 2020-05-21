package raml

import "github.com/Foxcapades/goop/v1/pkg/option"

type XML interface {
	IsAttribute() option.Bool
	SetIsAttribute(bool) XML
	UnsetIsAttribute() XML

	IsWrapped() option.Bool
	SetIsWrapped(bool) XML
	UnsetIsWrapped() XML

	Name() option.String
	SetName(string) XML
	UnsetName() XML

	Namespace() option.String
	SetNamespace(string) XML
	UnsetNamespace() XML

	Prefix() option.String
	SetPrefix(string) XML
	UnsetPrefix() XML
}
