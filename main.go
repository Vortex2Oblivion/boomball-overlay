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
	var boomball boomball = NewBoomball(0, 0, boomballTexture)

	var windowIcon rl.Image = *rl.LoadImage("7Dogs.png")
	rl.SetWindowIcon(windowIcon)
	rl.UnloadImage(&windowIcon)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(color.RGBA{})

		boomball.draw()
		boomball.x = float32(math.Tan(rl.GetTime()*1.5))*(1280-float32(boomballTexture.Width)*3) + float32(rl.GetScreenWidth())/2 - float32(boomballTexture.Width)/2
		boomball.y = float32(math.Sin(rl.GetTime()*1.5))*(720-float32(boomballTexture.Height)*3) + float32(rl.GetScreenHeight())/2 - float32(boomballTexture.Height)/2
		boomball.angle = float32(math.Tan(rl.GetTime()*1.5)) * 45
		boomball.color = rl.ColorFromHSV(float32(rl.GetTime())*15, 1, 1)

		if rl.IsKeyPressed(rl.KeyF1) {
			boomball.rainbowMode = !boomball.rainbowMode
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
