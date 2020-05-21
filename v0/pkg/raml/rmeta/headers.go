package rmeta

// Raml Headers
//noinspection GoUnusedConst
const (
	headerBase      = "#%RAML 1.0"
	headerFoot      = "\n---\n"
	HeaderRoot      = headerBase + headerFoot
	HeaderDataType  = headerBase + " DataType" + headerFoot
	HeaderExtension = headerBase + " Extension" + headerFoot
	HeaderLibrary   = headerBase + " Library" + headerFoot
	HeaderOverlay   = headerBase + " Overlay" + headerFoot
	HeaderTrait     = headerBase + " Trait" + headerFoot
)
