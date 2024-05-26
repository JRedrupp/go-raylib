package main

import (
	"math"

	"github.com/JRedrupp/go-raylib/src/bindings/raylib"
)

// ------------------------------------------------------------------------------------
// Program main entry point
// ------------------------------------------------------------------------------------
func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - following eyes")

	scleraLeftPosition := raylib.Vector2{X: float32(raylib.GetScreenWidth())/2.0 - 100.0, Y: float32(raylib.GetScreenHeight()) / 2.0}
	scleraRightPosition := raylib.Vector2{X: float32(raylib.GetScreenWidth())/2.0 + 100.0, Y: float32(raylib.GetScreenHeight()) / 2.0}
	scleraRadius := float32(80)

	irisLeftPosition := raylib.Vector2{X: float32(raylib.GetScreenWidth())/2.0 - 100.0, Y: float32(raylib.GetScreenHeight()) / 2.0}
	irisRightPosition := raylib.Vector2{X: float32(raylib.GetScreenWidth())/2.0 + 100.0, Y: float32(raylib.GetScreenHeight()) / 2.0}
	irisRadius := float32(24)

	angle := float32(0.0)
	dx, dy, dxx, dyy := float32(0.0), float32(0.0), float32(0.0), float32(0.0)

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		irisLeftPosition = raylib.GetMousePosition()
		irisRightPosition = raylib.GetMousePosition()

		// Check not inside the left eye sclera
		if !raylib.CheckCollisionPointCircle(irisLeftPosition, scleraLeftPosition, scleraRadius-irisRadius) {
			dx = irisLeftPosition.X - scleraLeftPosition.X
			dy = irisLeftPosition.Y - scleraLeftPosition.Y
			angle = float32(math.Atan2(float64(dy), float64(dx)))

			dxx = (scleraRadius - irisRadius) * float32(math.Cos(float64(angle)))
			dyy = (scleraRadius - irisRadius) * float32(math.Sin(float64(angle)))

			irisLeftPosition.X = scleraLeftPosition.X + dxx
			irisLeftPosition.Y = scleraLeftPosition.Y + dyy
		}

		// Check not inside the right eye sclera
		if !raylib.CheckCollisionPointCircle(irisRightPosition, scleraRightPosition, scleraRadius-irisRadius) {
			dx = irisRightPosition.X - scleraRightPosition.X
			dy = irisRightPosition.Y - scleraRightPosition.Y

			angle = float32(math.Atan2(float64(dy), float64(dx)))

			dxx = (scleraRadius - irisRadius) * float32(math.Cos(float64(angle)))
			dyy = (scleraRadius - irisRadius) * float32(math.Sin(float64(angle)))

			irisRightPosition.X = scleraRightPosition.X + dxx
			irisRightPosition.Y = scleraRightPosition.Y + dyy
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawCircleV(scleraLeftPosition, scleraRadius, raylib.LIGHTGRAY)
		raylib.DrawCircleV(irisLeftPosition, irisRadius, raylib.BROWN)
		raylib.DrawCircleV(irisLeftPosition, 10, raylib.BLACK)

		raylib.DrawCircleV(scleraRightPosition, scleraRadius, raylib.LIGHTGRAY)
		raylib.DrawCircleV(irisRightPosition, irisRadius, raylib.DARKGREEN)
		raylib.DrawCircleV(irisRightPosition, 10, raylib.BLACK)

		raylib.DrawFPS(10, 10)

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
