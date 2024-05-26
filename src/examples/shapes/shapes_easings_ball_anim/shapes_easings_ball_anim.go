package main

import (
	"github.com/JRedrupp/go-raylib/src/bindings/raylib"
	"github.com/JRedrupp/go-raylib/src/bindings/reasings"
)

// ------------------------------------------------------------------------------------
// Program main entry point
// ------------------------------------------------------------------------------------
func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - easings ball anim")

	// Ball variable value to be animated with easings
	var ballPositionX int32 = -100
	var ballRadius int32 = 20
	var ballAlpha float32 = 0.0

	state := 0
	framesCounter := 0

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		if state == 0 {
			framesCounter++
			ballPositionX = int32(reasings.EaseElasticOut((float32)(framesCounter), -100, screenWidth/2.0+100, 120))

			if framesCounter >= 120 {
				framesCounter = 0
				state = 1
			}
		} else if state == 1 { // Increase ball radius with easing
			framesCounter++
			ballRadius = int32(reasings.EaseElasticIn((float32)(framesCounter), 20, 500, 200))

			if framesCounter >= 200 {
				framesCounter = 0
				state = 2
			}
		} else if state == 2 { // Change ball alpha with easing (background color blending)
			framesCounter++
			ballAlpha = reasings.EaseCubicOut((float32)(framesCounter), 0.0, 1.0, 200)

			if framesCounter >= 200 {
				framesCounter = 0
				state = 3
			}
		} else if state == 3 { // Reset state to play again
			if raylib.IsKeyPressed(raylib.KEY_ENTER) {
				// Reset required variables to play again
				ballPositionX = -100
				ballRadius = 20
				ballAlpha = 0.0
				state = 0
			}
		}

		if raylib.IsKeyPressed(raylib.KEY_R) {
			framesCounter = 0
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		if state >= 2 {
			raylib.DrawRectangle(0, 0, screenWidth, screenHeight, raylib.GREEN)
		}
		raylib.DrawCircle(ballPositionX, 200, (float32)(ballRadius), raylib.Fade(raylib.RED, (1.0-ballAlpha)))

		if state == 3 {
			raylib.DrawText("PRESS [ENTER] TO PLAY AGAIN!", 240, 200, 20, raylib.BLACK)
		}

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
