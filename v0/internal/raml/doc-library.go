package raml

import (
	"errors"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func NewLibrary() *Library {
	return &Library{
		hasAnnotations: makeAnnotations(),
		hasAnnTypes:    makeAnnTypes(),
		hasExtra:       makeExtra(),
		hasResTypes:    makeResTypes(),
		hasSecSchemes:  makeSecSchemes(),
		hasTraits:      makeHasTraits(),
		hasTypes:       makeHasTypes(),
		hasUsage:       makeUsage(),
		hasUses:        makeUses(),
	}
}

type Library struct {
	hasAnnotations
	hasAnnTypes
	hasExtra
	hasResTypes
	hasSecSchemes
	hasTraits
	hasTypes
	hasUsage
	hasUses
}

func (l *Library) UnmarshalYAML(root *yaml.Node) error {
	logrus.Trace("internal.Library.UnmarshalYAML")

	if root.Kind != yaml.DocumentNode {
		return errors.New("cannot unmarshal document from a non-root node")
	}

	root = root.Content[0]

	if root.Kind != yaml.MappingNode {
		return errors.New("invalid document structure, expected !!map, got " + root.Tag)
	}

	for i := 0; i < len(root.Content); i += 2 {
		key := root.Content[i]
		val := root.Content[i+1]

		if err := l.assign(key, val); err != nil {
			return err
		}
	}

	return nil
}

func (l Library) MarshalYAML() (interface{}, error) {
	logrus.Trace("internal.Library.MarshalYAML")
	out := NewAnyMap()
	l.hasUsage.out(out)
	l.hasUses.out(out)
	l.hasSecSchemes.out(out)
	l.hasTraits.out(out)
	l.hasAnnTypes.out(out)
	l.hasAnnotations.out(out)
	l.hasResTypes.out(out)
	l.hasExtra.out(out)
	l.hasTypes.out(out)
	return out, nil
}

func (l *Library) assign(k, v *yaml.Node) error {
	logrus.Trace("internal.Library.assign")
	switch k.Value {
	case rmeta.KeyAnnotationTypes:
		return l.hasAnnTypes.in(v)
	case rmeta.KeyResourceTypes:
		return l.hasResTypes.in(v)
	case rmeta.KeyTypes, rmeta.KeySchemas:
		return l.hasTypes.in(v)
	case rmeta.KeyTraits:
		return l.hasTraits.in(v)
	case rmeta.KeyUses:
		return l.hasUses.in(v)
	case rmeta.KeyUsage:
		return l.hasUsage.in(v)
	case rmeta.KeySecuritySchemes:
		return l.hasSecSchemes.in(v)
	}

	if used, err := l.hasAnnotations.in(k.Value, v); err != nil {
		return err
	} else if used {
		return nil
	}

	l.hasExtra.in(k, v)
	return nil
}
