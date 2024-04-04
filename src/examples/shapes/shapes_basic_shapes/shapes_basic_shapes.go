package main

import "github.com/JRedrupp/go-snake/src/raylib"

func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - basic shapes drawing")

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		rotation := float32(0.2)
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawText("some basic shapes available on raylib", 20, 20, 20, raylib.DARKGRAY)

		// Circle shapes and lines
		raylib.DrawCircle(screenWidth/5, 120, 35, raylib.DARKBLUE)
		raylib.DrawCircleGradient(screenWidth/5, 220, 60, raylib.GREEN, raylib.SKYBLUE)
		raylib.DrawCircleLines(screenWidth/5, 340, 80, raylib.DARKBLUE)

		// Rectangle shapes and lines
		raylib.DrawRectangle(screenWidth/4*2-60, 100, 120, 60, raylib.RED)
		raylib.DrawRectangleGradientH(screenWidth/4*2-90, 170, 180, 130, raylib.MAROON, raylib.GOLD)
		raylib.DrawRectangleLines(screenWidth/4*2-40, 320, 80, 60, raylib.ORANGE) // NOTE: Uses QUADS internally, not lines

		// Triangle shapes and lines
		raylib.DrawTriangle(raylib.Vector2{X: screenWidth / 4.0 * 3.0, Y: 80.0},
			raylib.Vector2{X: screenWidth/4.0*3.0 - 60.0, Y: 150.0},
			raylib.Vector2{X: screenWidth/4.0*3.0 + 60.0, Y: 150.0}, raylib.VIOLET)

		raylib.DrawTriangleLines(raylib.Vector2{X: screenWidth / 4.0 * 3.0, Y: 160.0},
			raylib.Vector2{X: screenWidth/4.0*3.0 - 20.0, Y: 230.0},
			raylib.Vector2{X: screenWidth/4.0*3.0 + 20.0, Y: 230.0}, raylib.DARKBLUE)

		// Polygon shapes and lines
		raylib.DrawPoly(raylib.Vector2{X: screenWidth / 4.0 * 3, Y: 330}, 6, 80, rotation, raylib.BROWN)
		raylib.DrawPolyLines(raylib.Vector2{X: screenWidth / 4.0 * 3, Y: 330}, 6, 90, rotation, raylib.BROWN)
		raylib.DrawPolyLinesEx(raylib.Vector2{X: screenWidth / 4.0 * 3, Y: 330}, 6, 85, rotation, 6, raylib.BEIGE)

		// NOTE: We draw all LINES based shapes together to optimize internal drawing,
		// this way, all LINES are rendered in a single draw pass
		raylib.DrawLine(18, 42, screenWidth-18, 42, raylib.BLACK)
		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
