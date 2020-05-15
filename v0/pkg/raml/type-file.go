package raml

const (
	TypeFile = "file"

	DefaultFileMinLength = uint(0)
	DefaultFileMaxLength = uint(2_147_483_647)
)

func NewFile() *File {
	return &File{Base: Base{Type: TypeFile}}
}

type File struct {
	Base

	FileTypes []string `yaml:"fileTypes,omitempty"`
	MinLength *uint    `yaml:"minLength,omitempty"`
	MaxLength *uint    `yaml:"maxLength,omitempty"`
}

func (f *File) ToRAML() (string, error) {
	return dataTypeRaml(f)
}

type fAlias File

func (f *File) MarshalYAML() (interface{}, error) {
	if f.canSimplify() {
		return f.Type, nil
	}
	return fAlias(*f), nil
}

func (f *File) canSimplify() bool {
	if !f.Base.canSimplify() {
		return false
	}
	if f.FileTypes != nil && len(f.FileTypes) > 0 {
		return false
	}
	if f.MinLength != nil && *f.MinLength != DefaultFileMinLength {
		return false
	}
	if f.MaxLength != nil && *f.MaxLength != DefaultFileMaxLength {
		return false
	}

	return true
}

func (f *File) GetType() string {
	return TypeFile
}
