package main

import "github.com/JRedrupp/go-raylib/src/bindings/raylib"

func main() {
	// Initialization
	//---------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	raylib.SetConfigFlags(raylib.FLAG_MSAA_4X_HINT)
	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - bouncing ball")

	ballPosition := raylib.Vector2{X: float32(raylib.GetScreenWidth()) / 2.0, Y: float32(raylib.GetScreenHeight()) / 2.0}
	ballSpeed := raylib.Vector2{X: 5.0, Y: 4.0}
	ballRadius := int32(20)

	pause := false
	framesCounter := 0

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//----------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update
		//-----------------------------------------------------
		if raylib.IsKeyPressed(raylib.KEY_SPACE) {
			pause = !pause
		}

		if !pause {
			ballPosition.X += ballSpeed.X
			ballPosition.Y += ballSpeed.Y

			// Check walls collision for bouncing
			if ballPosition.X >= float32((raylib.GetScreenWidth()-ballRadius)) || (ballPosition.X <= float32(ballRadius)) {
				ballSpeed.X *= -1.0
			}
			if (ballPosition.Y >= float32((raylib.GetScreenHeight() - ballRadius))) || (ballPosition.Y <= float32(ballRadius)) {
				ballSpeed.Y *= -1.0
			}
		} else {
			framesCounter++
		}
		//-----------------------------------------------------

		// Draw
		//-----------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawCircleV(ballPosition, float32(ballRadius), raylib.MAROON)
		//DrawText("PRESS SPACE to PAUSE BALL MOVEMENT", 10, GetScreenHeight() - 25, 20, LIGHTGRAY);

		// On pause, we draw a blinking message
		if pause && (framesCounter/30)%2 == 0 {
			raylib.DrawText("PAUSED", 350, 200, 30, raylib.GRAY)
		}

		raylib.DrawFPS(10, 10)

		raylib.EndDrawing()
		//-----------------------------------------------------
	}

	// De-Initialization
	//---------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//----------------------------------------------------------
}
