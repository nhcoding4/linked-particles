package main

import rl "github.com/gen2brain/raylib-go/raylib"

// ---------------------------------------------------------------------------------------------------------------------

func main() {
	game := Game{
		width:           1000,
		height:          1000,
		title:           "Particles",
		targetFPS:       60,
		backgroundColor: rl.Black,
		totalParticles:  1500,
		mouseRadius:     200,
	}
	game.init()
	game.run()
}

// ---------------------------------------------------------------------------------------------------------------------
