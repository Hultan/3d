package three_d

import (
	"errors"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

type gtkBuilder struct {
	Builder *gtk.Builder
}

// newBuilder : Creates a new gtkBuilder
func newBuilder(glade string) (*gtkBuilder, error) {
	builder, err := gtk.BuilderNewFromString(glade)
	if err != nil {
		return nil, err
	}

	return &gtkBuilder{builder}, nil
}

// getObject : Gets a gtk object by name
func (g *gtkBuilder) getObject(name string) glib.IObject {
	if g.Builder == nil {
		panic(errors.New("need to call CreateBuilder first"))
	}
	obj, err := g.Builder.GetObject(name)
	if err != nil {
		panic(err)
	}

	return obj
}
