package main

import (
	"log"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type coordinate struct {
	x, y float64
}

// a component is anything that implements these two methods - onUpdate
// and onDraw
type component interface {
	onUpdate() error
	onDraw(renderder *sdl.Renderer) error
}

type element struct {
	position   coordinate
	rotation   float64
	active     bool
	components []component
}

func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *element) update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (elem *element) addComponent(new component) {
	for _, comp := range elem.components {
		if reflect.TypeOf(new) == reflect.TypeOf(comp) {
			log.Fatalf("Cannot add new component with existing type. %v",
				reflect.TypeOf(new))
		}
	}
	elem.components = append(elem.components, new)
}

func (elem *element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	log.Fatalf("No component with type %v", typ)
	return nil
}

var elements []*element
