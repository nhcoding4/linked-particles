package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	width           int32
	height          int32
	title           string
	targetFPS       int32
	backgroundColor rl.Color
	effect          *Effect
	mouse           *Mouse
	mouseRadius     int32
	totalParticles  int32
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) init() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(g.width, g.height, g.title)
	rl.SetTargetFPS(g.targetFPS)
	g.createMouse()
	g.createEffect()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) createEffect() {
	g.effect = &Effect{
		width:           g.width,
		height:          g.height,
		backgroundColor: rl.White,
		totalParticles:  g.totalParticles,
		mouse:           g.mouse,
	}
	g.effect.init()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) createMouse() {
	g.mouse = &Mouse{
		radius: g.mouseRadius,
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) run() {
	for !rl.WindowShouldClose() {
		g.updateState()

		rl.BeginDrawing()
		rl.ClearBackground(g.backgroundColor)
		g.effect.draw()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) updateState() {
	if rl.IsWindowResized() {
		g.updateWindowSize()
		g.effect.setPosition()
	}

	g.mouse.checkLeftClickDown()
	if *g.mouse.activationStatus() {
		g.mouse.update()
	}

	g.effect.update()
}

// ---------------------------------------------------------------------------------------------------------------------

func (g *Game) updateWindowSize() {
	g.width = int32(rl.GetScreenWidth())
	g.height = int32(rl.GetScreenHeight())
	g.effect.updateWindowSize()
}

// ---------------------------------------------------------------------------------------------------------------------
