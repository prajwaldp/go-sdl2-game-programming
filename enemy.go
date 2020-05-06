package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	enemySpeed = 5
	enemySize  = 128
)

type enemy struct {
	texture *sdl.Texture
	x, y    float64
}

func newEnemy(renderer *sdl.Renderer, x, y float64) *element {
	enemy := &element{}
	enemy.position = coordinate{x: x, y: y}
	enemy.active = true

	sr := newSpriteRenderer(enemy, renderer, "./sprites/enemy.svg")
	enemy.addComponent(sr)

	return enemy
}
