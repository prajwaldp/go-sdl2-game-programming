package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const bulletSpeed = 3

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}
	sr := newSpriteRenderer(bullet, renderer, "./sprites/bullet.svg")
	bullet.addComponent(sr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	bullet.active = false
	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bullet := newBullet(renderer)
		elements = append(elements, bullet)
		bulletPool = append(bulletPool, bullet)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bullet := range bulletPool {
		if !bullet.active {
			return bullet, true
		}
	}

	return nil, false
}
