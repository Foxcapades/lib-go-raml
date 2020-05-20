package rmeta

type DataTypeKind string

const (
	TypeAny          DataTypeKind = "any"
	TypeArray        DataTypeKind = "array"
	TypeBool         DataTypeKind = "boolean"
	TypeDateOnly     DataTypeKind = "date-only"
	TypeDatetime     DataTypeKind = "datetime"
	TypeDatetimeOnly DataTypeKind = "datetime-only"
	TypeFile         DataTypeKind = "file"
	TypeInteger      DataTypeKind = "integer"
	TypeNil          DataTypeKind = "nil"
	TypeNumber       DataTypeKind = "number"
	TypeObject       DataTypeKind = "object"
	TypeString       DataTypeKind = "string"
	TypeTimeOnly     DataTypeKind = "time-only"

	// Meta types
	TypeUnion   DataTypeKind = "union"
	TypeMulti   DataTypeKind = "multi"
	TypeCustom  DataTypeKind = "custom"
	TypeInclude DataTypeKind = "include"
)
