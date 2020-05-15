package raml

const (
	keyFileTypes = "fileTypes"
	keyMinLength = "minLength"
	keyMaxLength = "maxLength"
)

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

func (f *File) ToRaml() (string, error) {
	return dataTypeRaml(f)
}

type fAlias File
func (f *File) MarshalYAML() (interface{}, error) {
	if f.canSimplify() {
		return f.Type, nil
	}
	return fAlias(*f), nil
}

//func (o *File) UnmarshalYAML(value *yaml.Node) error {
//	if err := o.Base.UnmarshalYAML(value); err != nil {
//		return err
//	}
//
//	for i := 0; i < len(value.Content); i++ {
//		key := value.Content[i].Value
//		i++
//		if err := o.assign(key, value.Content[i]); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func (o *File) assign(key string, val *yaml.Node) error {
//	switch key {
//	case keyFileTypes:
//		o.FileTypes = make([]string, 0, len(val.Content))
//		for i := range val.Content {
//			o.FileTypes[i] = val.Content[i].Value
//		}
//	case keyMinLength:
//		tmp, err := ParseUint(val)
//		if err != nil {
//			return err
//		}
//		o.MinLength = &tmp
//	case keyMaxLength:
//		tmp, err := ParseUint(val)
//		if err != nil {
//			return err
//		}
//		o.MaxLength = &tmp
//	}
//	return nil
//}

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
