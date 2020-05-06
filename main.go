package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 800
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Cannot initialize sdl. %v", err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("My Cool Game", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		log.Fatalf("Cannot create window. %v", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Cannot create renderer. %v", err)
	}
	defer renderer.Destroy()

	player := newPlayer(renderer)
	elements = append(elements, player)

	var i, j float64

	for ; i < 3; i++ {
		for j = 0; j < 5; j++ {
			x := (j/5)*screenWidth + enemySize/2.0
			y := i * enemySize
			enemy := newEnemy(renderer, x, y)
			elements = append(elements, enemy)
		}
	}

	initBulletPool(renderer)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quitting")
				running = false
				break
			}
		}

		// renderer.SetDrawColor(23, 3, 18, 255)
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, element := range elements {
			if element.active {
				err := element.update()
				if err != nil {
					log.Fatalf("Cannot update element. %v", err)
				}

				err = element.draw(renderer)
				if err != nil {
					log.Fatalf("Cannot draw element. %v", err)
				}
			}
		}
		renderer.Present()
	}
}
