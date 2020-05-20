package rmod

const headerBase = "#%RAML 1.0"

type ModuleType string

// RAML Module Types
//noinspection GoUnusedConst
const (
	TypeRoot ModuleType = "root"

	TypeDataType     ModuleType = "DataType"
	TypeExtension    ModuleType = "Extension"
	TypeLibrary      ModuleType = "Library"
	TypeResourceType ModuleType = "ResourceType"
	TypeTrait        ModuleType = "Trait"
)

func (M ModuleType) ToHeader() string {
	if M == TypeRoot {
		return headerBase
	}

	return headerBase + " " + string(M)
}
