package raml

import (
	"github.com/Foxcapades/lib-go-raml/v0/pkg/raml"
	"github.com/Foxcapades/lib-go-yaml/v1/pkg/xyml"
	"gopkg.in/yaml.v3"
)

func UnmarshalAnyExampleMapRAML(aMap raml.AnyExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewAnyExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}
func UnmarshalArrayExampleMapRAML(aMap raml.ArrayExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewArrayExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalBoolExampleMapRAML(aMap raml.BoolExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewBoolExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalCustomExampleMapRAML(aMap raml.CustomExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewCustomExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalDateOnlyExampleMapRAML(aMap raml.DateOnlyExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewDateOnlyExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalDatetimeOnlyExampleMapRAML(aMap raml.DatetimeOnlyExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewDatetimeOnlyExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalDatetimeExampleMapRAML(aMap raml.DatetimeExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewDatetimeExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalFileExampleMapRAML(aMap raml.FileExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewFileExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalIntegerExampleMapRAML(aMap raml.IntegerExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewIntegerExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalNumberExampleMapRAML(aMap raml.NumberExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewNumberExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalObjectExampleMapRAML(aMap raml.ObjectExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewObjectExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalStringExampleMapRAML(aMap raml.StringExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewStringExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalTimeOnlyExampleMapRAML(aMap raml.TimeOnlyExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewTimeOnlyExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}

func UnmarshalUnionExampleMapRAML(aMap raml.UnionExampleMap, y *yaml.Node) error {
	return xyml.MapForEach(y, func(k, v *yaml.Node) error {
		if err := xyml.RequireString(k); err != nil {
			return err
		}

		tmp := NewUnionExample()

		if err := tmp.UnmarshalRAML(v); err != nil {
			return err
		}

		aMap.Put(k.Value, tmp)

		return nil
	})
}
