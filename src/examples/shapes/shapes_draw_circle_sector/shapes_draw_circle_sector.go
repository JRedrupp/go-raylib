package main

import (
	"math"

	"github.com/JRedrupp/go-raylib/src/bindings/model"
	"github.com/JRedrupp/go-raylib/src/bindings/raygui"
	"github.com/JRedrupp/go-raylib/src/bindings/raylib"
)

func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth int32 = 800
	const screenHeight int32 = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - draw circle sector")

	center := raylib.Vector2{X: (float32(raylib.GetScreenWidth()-300) / 2.0), Y: float32(raylib.GetScreenHeight() / 2.0)}

	outerRadius := float32(180.0)
	startAngle := float32(0.0)
	endAngle := float32(180.0)
	segments := float32(10.0)
	minSegments := float32(4)

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		// NOTE: All variables update happens inside GUI control functions
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawLine(500, 0, 500, raylib.GetScreenHeight(), raylib.Fade(raylib.LIGHTGRAY, 0.6))
		raylib.DrawRectangle(500, 0, raylib.GetScreenWidth()-500, raylib.GetScreenHeight(), raylib.Fade(raylib.LIGHTGRAY, 0.3))

		raylib.DrawCircleSector(center, outerRadius, startAngle, endAngle, int32(segments), raylib.Fade(raylib.MAROON, 0.3))
		raylib.DrawCircleSectorLines(center, outerRadius, startAngle, endAngle, int32(segments), raylib.Fade(raylib.MAROON, 0.6))

		// Draw GUI controls
		//------------------------------------------------------------------------------
		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 40, Width: 120, Height: 20}, "StartAngle", "", &startAngle, 0, 720)
		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 70, Width: 120, Height: 20}, "EndAngle", "", &endAngle, 0, 720)

		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 140, Width: 120, Height: 20}, "Radius", "", &outerRadius, 0, 200)
		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 170, Width: 120, Height: 20}, "Segments", "", &segments, 0, 100)
		//------------------------------------------------------------------------------

		minSegments = float32(math.Trunc(math.Ceil(float64((endAngle - startAngle) / 90))))
		modeType := "AUTO"
		modeColor := raylib.MAROON
		if segments >= minSegments {
			modeType = "MANUAL"
			modeColor = raylib.DARKGRAY
		}
		raylib.DrawText(raylib.TextFormat("MODE: %s", modeType), 600, 200, 10, modeColor)

		raylib.DrawFPS(10, 10)

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
