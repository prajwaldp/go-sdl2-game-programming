package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container     *element
	texture       *sdl.Texture
	width, height float64
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	texture, err := newTextureFromImg(filename, renderer)
	if err != nil {
		log.Fatalf("Cannot create texture from image. %v", err)
	}
	_, _, width, height, err := texture.Query()

	if err != nil {
		log.Fatalf("Cannot query texture. %v", err)
	}

	return &spriteRenderer{
		container: container,
		texture:   texture,
		width:     float64(width),
		height:    float64(height),
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	x := sr.container.position.x - sr.width/2.0
	y := sr.container.position.y - sr.height/2.0

	renderer.CopyEx(
		sr.texture,
		&sdl.Rect{X: 0, Y: 0, W: int32(sr.width), H: int32(sr.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(sr.width), H: int32(sr.height)},
		sr.container.rotation,
		&sdl.Point{X: int32(sr.width) / 2, Y: int32(sr.height) / 2},
		sdl.FLIP_NONE,
	)

	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}
