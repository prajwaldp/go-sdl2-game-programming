package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bulletMover struct {
	container *element
	speed     float64
}

func newBulletMover(container *element, speed float64) *bulletMover {
	return &bulletMover{
		container: container,
		speed:     speed,
	}
}

func (mover *bulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *bulletMover) onUpdate() error {
	container := mover.container
	container.position.x += bulletSpeed * math.Cos(container.rotation)
	container.position.y += bulletSpeed * math.Sin(container.rotation)

	if container.position.x > screenWidth || container.position.x < 0 ||
		container.position.y > screenHeight || container.position.y < 0 {
		container.active = false
	}

	return nil
}
