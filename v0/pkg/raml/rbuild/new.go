package rbuild

import (
	impl "github.com/Foxcapades/lib-go-raml-types/v0/internal/raml"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml"
)

// NewLibrary creates a new, empty instance of the raml.Library type.
func NewLibrary() raml.Library {
	return impl.NewLibrary()
}

// NewAPISpec creates a new, empty instance of the raml.APISpec type.
func NewAPISpec() raml.APISpec {
	return impl.NewApiSpec()
}

// NewAnyDataType creates a new, empty instance of the raml.AnyType type.
func NewAnyDataType() raml.AnyType {
	return impl.NewAnyType()
}

// NewArrayDataType creates a new, empty instance of the raml.ArrayType type.
func NewArrayDataType() raml.ArrayType {
	return impl.NewArrayType()
}

// NewBoolDataType creates a new, empty instance of the raml.BoolType type.
func NewBoolDataType() raml.BoolType {
	return impl.NewBoolType()
}

// NewIntegerDataType creates a new, empty instance of the raml.IntegerType
// type.
func NewIntegerDataType() raml.IntegerType {
	return impl.NewIntegerType()
}

// NewNumberDataType creates a new, empty instance of the raml.NumberType type.
func NewNumberDataType() raml.NumberType {
	return impl.NewNumberType()
}

// NewStringDataType creates a new, empty instance of the raml.StringType type.
func NewStringDataType() raml.StringType {
	return impl.NewStringType()
}
