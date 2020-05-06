package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func newTextureFromImg(path string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	image, err := img.Load(path)
	if err != nil {
		return &sdl.Texture{}, fmt.Errorf("Cannot load image. %v", err)
	}
	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		return &sdl.Texture{}, fmt.Errorf("Cannot create texture from image. %v", err)
	}

	return texture, nil
}
