package raml

import (
	"github.com/Foxcapades/lib-go-raml-types/v0/internal/xlog"
	"github.com/Foxcapades/lib-go-raml-types/v0/pkg/raml/rmeta"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func NewLibrary(log *logrus.Entry) *Library {
	log = xlog.WithType(log, "internal.Library")
	return &Library{
		log:            log,
		hasAnnotations: makeAnnotations(log),
		hasAnnTypes:    makeAnnTypes(log),
		hasExtra:       makeExtra(log),
		hasResTypes:    makeResTypes(log),
		hasSecSchemes:  makeSecSchemes(log),
		hasTraits:      makeHasTraits(log),
		hasTypes:       makeHasTypes(log),
		hasUsage:       makeUsage(log),
		hasUses:        makeUses(log),
	}
}

type Library struct {
	log *logrus.Entry
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

func (l *Library) UnmarshalYAML(fn func(interface{}) error) error {
	l.log.Trace("internal.Library.UnmarshalYAML")
	var slice yaml.MapSlice

	if err := fn(&slice); err != nil {
		return err
	}

	for i := range slice {
		log := xlog.AddPath(l.log, slice[i].Key)
		if err := l.assign(slice[i].Key, slice[i].Value, log); err != nil {
			return err
		}
	}

	return nil
}

func (l Library) MarshalYAML() (interface{}, error) {
	l.log.Trace("internal.Library.MarshalYAML")
	out := NewAnyMap(l.log)
	l.hasUsage.out(out)
	l.hasUses.out(out)
	l.hasSecSchemes.out(out)
	l.hasTraits.out(out)
	l.hasAnnTypes.out(out)
	l.hasAnnotations.out(out)
	l.hasResTypes.out(out)
	l.hasExtra.out(out)
	l.hasTypes.out(out, l.log)
	return out, nil
}

func (l *Library) assign(k, v interface{}, log *logrus.Entry) error {
	log.Trace("internal.Library.assign")
	switch k {
	case rmeta.KeyAnnotationTypes:
		return l.hasAnnTypes.in(v, l.log)
	case rmeta.KeyResourceTypes:
		return l.hasResTypes.in(v, l.log)
	case rmeta.KeyTypes, rmeta.KeySchemas:
		return l.hasTypes.in(v, l.log)
	case rmeta.KeyTraits:
		return l.hasTraits.in(v, l.log)
	case rmeta.KeyUses:
		return l.hasUses.in(v, l.log)
	case rmeta.KeyUsage:
		return l.hasUsage.in(v)
	case rmeta.KeySecuritySchemes:
		return l.hasSecSchemes.in(v, l.log)
	}

	if str, ok := k.(string); ok {
		if used, err := l.hasAnnotations.in(str, v); err != nil {
			return err
		} else if used {
			return nil
		}
	}

	l.hasExtra.in(k, v)
	return nil
}
