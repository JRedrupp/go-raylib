package main

import (
	"fmt"

	"github.com/JRedrupp/go-raylib/src/bindings/model"
	"github.com/JRedrupp/go-raylib/src/bindings/raygui"
	"github.com/JRedrupp/go-raylib/src/bindings/raylib"
)

func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth int32 = 800
	const screenHeight int32 = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - draw rectangle rounded")

	var roundness float32 = 0.2
	var width float32 = 200.0
	var height float32 = 100.0
	var segments float32 = 0.0
	var lineThick float32 = 1.0

	var drawRect bool = false
	var drawRoundedRect bool = true
	var drawRoundedLines bool = false

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------

		var rec = model.Rectangle{X: ((float32)(raylib.GetScreenWidth()) - width - 250) / 2, Y: (float32(raylib.GetScreenHeight()) - height) / 2.0, Width: (float32)(width), Height: (float32)(height)}
		//----------------------------------------------------------------------------------
		if raylib.IsKeyDown(raylib.KEY_A) {
			fmt.Println("X:", rec.X, "Y:", rec.Y, "Width:", rec.Width, "Height:", rec.Height)
		}
		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawLine(560, 0, 560, raylib.GetScreenHeight(), raylib.Fade(raylib.LIGHTGRAY, 0.6))
		raylib.DrawRectangle(560, 0, raylib.GetScreenWidth()-500, raylib.GetScreenHeight(), raylib.Fade(raylib.LIGHTGRAY, 0.3))

		if drawRect {
			raylib.DrawRectangleRec(rec, raylib.Fade(raylib.GOLD, 0.6))
		}
		if drawRoundedRect {
			raylib.DrawRectangleRounded(rec, roundness, (int32)(segments), raylib.Fade(raylib.MAROON, 0.2))
		}
		if drawRoundedLines {
			raylib.DrawRectangleRoundedLines(rec, roundness, (int32)(segments), raylib.Fade(raylib.MAROON, 0.4))
		}

		// Draw GUI controls
		//------------------------------------------------------------------------------
		raygui.GuiSliderBar(model.Rectangle{X: 640, Y: 40, Width: 105, Height: 20}, "Width", "", &width, 0, (float32)(raylib.GetScreenWidth()-300))
		raygui.GuiSliderBar(model.Rectangle{X: 640, Y: 70, Width: 105, Height: 20}, "Height", "", &height, 0, (float32)(raylib.GetScreenHeight()-50))
		raygui.GuiSliderBar(model.Rectangle{X: 640, Y: 140, Width: 105, Height: 20}, "Roundness", "", &roundness, 0.0, 1.0)
		raygui.GuiSliderBar(model.Rectangle{X: 640, Y: 170, Width: 105, Height: 20}, "Thickness", "", &lineThick, 0, 20)
		raygui.GuiSliderBar(model.Rectangle{X: 640, Y: 240, Width: 105, Height: 20}, "Segments", "", &segments, 0, 60)

		raygui.GuiCheckBox(model.Rectangle{X: 640, Y: 320, Width: 20, Height: 20}, "DrawRoundedRect", &drawRoundedRect)
		raygui.GuiCheckBox(model.Rectangle{X: 640, Y: 350, Width: 20, Height: 20}, "DrawRoundedLines", &drawRoundedLines)
		raygui.GuiCheckBox(model.Rectangle{X: 640, Y: 380, Width: 20, Height: 20}, "DrawRect", &drawRect)
		//------------------------------------------------------------------------------
		mode := "AUTO"
		col := raylib.DARKGRAY
		if segments >= 4 {
			mode = "MANUAL"
			col = raylib.MAROON
		}

		raylib.DrawText(raylib.TextFormat("MODE: %s", mode), 640, 280, 10, col)

		raylib.DrawFPS(10, 10)

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------

}
