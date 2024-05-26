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
	const screenWidth = 800
	const screenHeight = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - draw ring")

	var center = raylib.Vector2{X: float32(raylib.GetScreenWidth()-300) / 2.0, Y: float32(raylib.GetScreenHeight()) / 2.0}

	var innerRadius float32 = 80.0
	var outerRadius float32 = 190.0

	var startAngle float32 = 0.0
	var endAngle float32 = 360.0
	var segments float32 = 0.0

	var drawRing bool = true
	var drawRingLines bool = false
	var drawCircleLines bool = false

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

		if drawRing {
			raylib.DrawRing(center, innerRadius, outerRadius, startAngle, endAngle, (int32)(segments), raylib.Fade(raylib.MAROON, 0.3))
		}
		if drawRingLines {
			raylib.DrawRingLines(center, innerRadius, outerRadius, startAngle, endAngle, (int32)(segments), raylib.Fade(raylib.BLACK, 0.4))
		}
		if drawCircleLines {
			raylib.DrawCircleSectorLines(center, outerRadius, startAngle, endAngle, (int32)(segments), raylib.Fade(raylib.BLACK, 0.4))
		}

		// Draw GUI controls
		//------------------------------------------------------------------------------
		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 40, Width: 120, Height: 20}, "StartAngle", "", &startAngle, -450, 450)
		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 70, Width: 120, Height: 20}, "EndAngle", "", &endAngle, -450, 450)

		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 140, Width: 120, Height: 20}, "InnerRadius", "", &innerRadius, 0, 100)
		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 170, Width: 120, Height: 20}, "OuterRadius", "", &outerRadius, 0, 200)

		raygui.GuiSliderBar(model.Rectangle{X: 600, Y: 240, Width: 120, Height: 20}, "Segments", "", &segments, 0, 100)

		raygui.GuiCheckBox(model.Rectangle{X: 600, Y: 320, Width: 20, Height: 20}, "Draw Ring", &drawRing)
		raygui.GuiCheckBox(model.Rectangle{X: 600, Y: 350, Width: 20, Height: 20}, "Draw RingLines", &drawRingLines)
		raygui.GuiCheckBox(model.Rectangle{X: 600, Y: 380, Width: 20, Height: 20}, "Draw CircleLines", &drawCircleLines)
		//------------------------------------------------------------------------------

		var minSegments = (int)(math.Ceil((float64)(endAngle-startAngle) / 90))
		var mode = "AUTO"
		if int(segments) >= minSegments {
			mode = "MANUAL"
		}

		var color = raylib.DARKGRAY
		if int(segments) >= minSegments {
			color = raylib.MAROON
		}

		raylib.DrawText(raylib.TextFormat("MODE: %s", mode), 600, 270, 10, color)

		raylib.DrawFPS(10, 10)

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------

	return
}
