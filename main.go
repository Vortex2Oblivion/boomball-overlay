package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowTransparent)
	rl.SetConfigFlags(rl.FlagWindowMousePassthrough)
	rl.SetConfigFlags(rl.FlagWindowTopmost)
	rl.InitWindow(int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), "Look at him go!")

	rl.SetTargetFPS(int32(rl.GetMonitorRefreshRate(rl.GetCurrentMonitor())))

	var boomballTexture rl.Texture2D = rl.LoadTexture("7Dogs.png")
	rl.SetTextureFilter(boomballTexture, rl.FilterBilinear)
	var boomballs []boomball = []boomball{NewBoomball(0, 0, boomballTexture)}

	var windowIcon rl.Image = *rl.LoadImage("7Dogs.png")
	rl.SetWindowIcon(windowIcon)
	rl.UnloadImage(&windowIcon)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(color.RGBA{})

		for i := range len(boomballs) {
			var f = float64(i)
			boomballs[i].draw()
			boomballs[i].x = float32(math.Tan(f*500+rl.GetTime()*1.5))*(1280-float32(boomballTexture.Width)*3) + float32(rl.GetScreenWidth())/2 - float32(boomballTexture.Width)/2
			boomballs[i].y = float32(math.Sin(f*500+rl.GetTime()*1.5))*(720-float32(boomballTexture.Height)*3.25) + float32(rl.GetScreenHeight())/2 - float32(boomballTexture.Height)/2
			boomballs[i].angle = float32(math.Tan(f*15+rl.GetTime()*1.5)) * 45
			boomballs[i].color = rl.ColorFromHSV(float32(rl.GetTime()+f*15)*15, 1, 1)

			if rl.IsKeyPressed(rl.KeyF1) {
				boomballs[i].rainbowMode = !boomballs[i].rainbowMode
			}
		}

		if rl.IsKeyPressed(rl.KeyF2) {
			boomballs = append(boomballs, NewBoomball(0, 0, boomballTexture))
			boomballs[len(boomballs)-1].rainbowMode = boomballs[0].rainbowMode
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
