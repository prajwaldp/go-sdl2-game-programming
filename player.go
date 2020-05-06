package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 3
	playerSize  = 64
)

type player struct {
	texture *sdl.Texture
	x, y    float64
}

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}

	player.position = coordinate{
		x: screenWidth / 2.0,
		y: screenHeight - playerSize,
	}

	player.active = true

	sr := newSpriteRenderer(player, renderer, "./sprites/player.svg")
	player.addComponent(sr)

	mover := newKeyboardMover(player, playerSpeed)
	player.addComponent(mover)

	shooter := newKeyboardShooter(player, time.Millisecond*250)
	player.addComponent(shooter)

	return player
}
