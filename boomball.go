package main

import rl "github.com/gen2brain/raylib-go/raylib"

type boomball struct {
	texture rl.Texture2D
	x       float32
	y       float32
	angle   float32
	color   rl.Color
}

func (b *boomball) draw() {
	rl.DrawTexturePro(b.texture, rl.NewRectangle(0, 0, float32(b.texture.Width), float32(b.texture.Width)),
		rl.NewRectangle(float32(b.texture.Width)/2+b.x, float32(b.texture.Height)/2+b.y, float32(b.texture.Width), float32(b.texture.Width)), rl.NewVector2(float32(b.texture.Width)/2, float32(b.texture.Width)/2), b.angle, b.color)
}

func NewBoomball(x, y float32, texture rl.Texture2D) boomball {
	return boomball{texture: texture, x: x, y: y}
}
