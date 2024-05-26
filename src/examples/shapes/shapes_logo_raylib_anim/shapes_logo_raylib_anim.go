package main

import "github.com/JRedrupp/go-raylib/src/bindings/raylib"

func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - raylib logo animation")

	const logoPositionX int32 = screenWidth/2 - 128
	const logoPositionY int32 = screenHeight/2 - 128

	framesCounter := 0
	lettersCount := int32(0)

	topSideRecWidth := int32(16)
	leftSideRecHeight := int32(16)

	bottomSideRecWidth := int32(16)
	rightSideRecHeight := int32(16)

	state := 0            // Tracking animation states (State Machine)
	alpha := float32(1.0) // Useful for fading

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		if state == 0 { // State 0: Small box blinking
			framesCounter++

			if framesCounter == 120 {
				state = 1
				framesCounter = 0 // Reset counter... will be used later...
			}
		} else if state == 1 { // State 1: Top and left bars growing

			topSideRecWidth += 4
			leftSideRecHeight += 4

			if topSideRecWidth == 256 {
				state = 2
			}
		} else if state == 2 { // State 2: Bottom and right bars growing

			bottomSideRecWidth += 4
			rightSideRecHeight += 4

			if bottomSideRecWidth == 256 {
				state = 3
			}
		} else if state == 3 { // State 3: Letters appearing (one by one)

			framesCounter++

			if framesCounter/12 == 0 { // Every 12 frames, one more letter!

				lettersCount++
				framesCounter = 0
			}

			if lettersCount >= 10 { // When all letters have appeared, just fade out everything

				alpha -= 0.02

				if alpha <= 0.0 {
					alpha = 0.0
					state = 4
				}
			}
		} else if state == 4 { // State 4: Reset and Replay

			if raylib.IsKeyPressed(raylib.KEY_R) {
				framesCounter = 0
				lettersCount = 0

				topSideRecWidth = 16
				leftSideRecHeight = 16

				bottomSideRecWidth = 16
				rightSideRecHeight = 16

				alpha = 1.0
				state = 0 // Return to State 0
			}
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		if state == 0 {
			if (framesCounter/15)%2 == 0 {
				raylib.DrawRectangle(logoPositionX, logoPositionY, 16, 16, raylib.BLACK)
			}
		} else if state == 1 {
			raylib.DrawRectangle(logoPositionX, logoPositionY, topSideRecWidth, 16, raylib.BLACK)
			raylib.DrawRectangle(logoPositionX, logoPositionY, 16, leftSideRecHeight, raylib.BLACK)
		} else if state == 2 {

			raylib.DrawRectangle(logoPositionX, logoPositionY, topSideRecWidth, 16, raylib.BLACK)
			raylib.DrawRectangle(logoPositionX, logoPositionY, 16, leftSideRecHeight, raylib.BLACK)

			raylib.DrawRectangle(logoPositionX+240, logoPositionY, 16, rightSideRecHeight, raylib.BLACK)
			raylib.DrawRectangle(logoPositionX, logoPositionY+240, bottomSideRecWidth, 16, raylib.BLACK)
		} else if state == 3 {
			raylib.DrawRectangle(logoPositionX, logoPositionY, topSideRecWidth, 16, raylib.Fade(raylib.BLACK, alpha))
			raylib.DrawRectangle(logoPositionX, logoPositionY+16, 16, leftSideRecHeight-32, raylib.Fade(raylib.BLACK, alpha))

			raylib.DrawRectangle(logoPositionX+240, logoPositionY+16, 16, rightSideRecHeight-32, raylib.Fade(raylib.BLACK, alpha))
			raylib.DrawRectangle(logoPositionX, logoPositionY+240, bottomSideRecWidth, 16, raylib.Fade(raylib.BLACK, alpha))

			raylib.DrawRectangle(raylib.GetScreenWidth()/2-112, raylib.GetScreenHeight()/2-112, 224, 224, raylib.Fade(raylib.RAYWHITE, alpha))

			raylib.DrawText(raylib.TextSubtext("raylib", 0, lettersCount), raylib.GetScreenWidth()/2-44, raylib.GetScreenHeight()/2+48, 50, raylib.Fade(raylib.BLACK, alpha))
		} else if state == 4 {

			raylib.DrawText("[R] REPLAY", 340, 200, 20, raylib.GRAY)
		}

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
