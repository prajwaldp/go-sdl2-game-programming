package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64
	sr        *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()
	container := mover.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if container.position.x-mover.sr.width/2.0 > 0 {
			container.position.x -= mover.speed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if container.position.x+mover.sr.width/2.0 < screenWidth {
			container.position.x += mover.speed
		}
	}

	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (shooter *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()
	position := shooter.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(shooter.lastShot) >= shooter.cooldown {
			shooter.shoot(position.x+25, position.y-25)
			shooter.shoot(position.x-25, position.y-25)
			shooter.lastShot = time.Now()
		}
	}
	return nil
}

func (shooter *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (shooter *keyboardShooter) shoot(x, y float64) {
	if bullet, ok := bulletFromPool(); ok {
		bullet.active = true
		bullet.position.x = x
		bullet.position.y = y
		bullet.rotation = 270 * (math.Pi / 180)
	}
}
