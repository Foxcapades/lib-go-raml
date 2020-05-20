package assign

import (
	"fmt"
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/util/xyml"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func AsString(v *yaml.Node, ptr *string) error {
	logrus.Trace("assign.AsString")

	if err := xyml.RequireString(v); err != nil {
		return err
	}

	*ptr = v.Value

	return nil
}

func AsStringPtr(v *yaml.Node, ptr **string) error {
	logrus.Trace("assign.AsStringPtr")

	if v.Tag != xyml.String {
		return fmt.Errorf(errReqType, xyml.String, v.Tag, v.Line, v.Column)
	}

	tmp := v.Value
	*ptr = &tmp

	return nil
}
