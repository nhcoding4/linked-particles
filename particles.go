package main

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Particle struct {
	width        int32
	height       int32
	radius       int32
	x            int32
	y            int32
	movementX    int32
	movementY    int32
	mouse        *Mouse
	outlineColor rl.Color
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) init() {
	p.setMovementSpeed()
	p.radius = int32(rand.Intn(10)) + 5
	p.setPosition()
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) collision() {
	if 0+p.radius > p.x {
		p.x = 1 + p.radius
		p.movementX *= -1
	}
	if p.x+p.radius > p.width {
		p.x = p.width - p.radius - 1
		p.movementX *= -1
	}

	if 0+p.radius > p.y {
		p.y = 1 + p.radius
		p.movementY *= -1
	}
	if p.y+p.radius > p.height {
		p.y = p.height - p.radius - 1
		p.movementY *= -1
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) draw() {
	color1 := rl.ColorFromHSV(float32(p.x), 1.0, 1.0)
	color2 := rl.ColorFromHSV(float32(p.y), 1.0, 1.0)

	rl.DrawCircleGradient(p.x, p.y, float32(p.radius), color1, color2)
	rl.DrawCircleLines(p.x, p.y, float32(p.radius), p.outlineColor)
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) movement() {
	p.x += p.movementX
	p.y += p.movementY
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) pushParticle() {
	mx, my := p.mouse.getLocation()
	dx := p.x - *mx
	dy := p.y - *my

	distance := math.Sqrt(float64(dx*dx + dy*dy))

	if distance < float64(p.radius+*p.mouse.getRadius()) {
		force := int32(3)

		if p.x < *mx && p.x > p.radius+10 {
			p.x -= force
		}
		if p.x > *mx && p.x < p.width-(p.radius+10) {
			p.x += force
		}
		if p.y < *my && p.y > p.radius+10 {
			p.y -= force
		}
		if p.y > *my && p.y < p.height-(p.radius+10) {
			p.y += force
		}
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) setMovementSpeed() {
	p.movementX = int32(rand.Intn(4)) - 2
	p.movementY = int32(rand.Intn(4)) - 2
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) setPosition() {
	p.x = int32(rand.Intn(int(p.width)-int(p.radius))) + p.radius
	p.y = int32(rand.Intn(int(p.height)-int(p.radius))) + p.radius
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) update() {
	p.movement()
	if *p.mouse.activationStatus() {
		p.pushParticle()
	}
	p.collision()
}

// ---------------------------------------------------------------------------------------------------------------------

func (p *Particle) updateWindowSize() {
	p.width = int32(rl.GetScreenWidth())
	p.height = int32(rl.GetScreenHeight())
}

// ---------------------------------------------------------------------------------------------------------------------
