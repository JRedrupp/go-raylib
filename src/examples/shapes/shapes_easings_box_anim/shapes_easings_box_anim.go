package main

import (
	"github.com/JRedrupp/go-raylib/src/bindings/model"
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

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - easings box anim")

	// Box variables to be animated with easings
	rec := model.Rectangle{X: float32(raylib.GetScreenWidth()) / 2.0, Y: -100, Width: 100, Height: 100}
	var rotation float32 = 0.0
	var alpha float32 = 1.0

	state := 0
	framesCounter := 0

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		switch state {
		case 0: // Move box down to center of screen
			{
				framesCounter++

				// NOTE: Remember that 3rd parameter of easing function refers to
				// desired value variation, do not confuse it with expected final value!
				rec.Y = reasings.EaseElasticOut(float32(framesCounter), -100, float32(raylib.GetScreenHeight())/2.0+100, 120)

				if framesCounter >= 120 {
					framesCounter = 0
					state = 1
				}
			}
		case 1: // Scale box to an horizontal bar
			{
				framesCounter++
				rec.Height = reasings.EaseBounceOut(float32(framesCounter), 100, -90, 120)
				rec.Width = reasings.EaseBounceOut(float32(framesCounter), 100, float32(raylib.GetScreenWidth()), 120)

				if framesCounter >= 120 {
					framesCounter = 0
					state = 2
				}
			}
		case 2: // Rotate horizontal bar rectangle
			{
				framesCounter++
				rotation = reasings.EaseQuadOut(float32(framesCounter), 0.0, 270.0, 240)

				if framesCounter >= 240 {
					framesCounter = 0
					state = 3
				}
			}
		case 3: // Increase bar size to fill all screen
			{
				framesCounter++
				rec.Height = reasings.EaseCircOut(float32(framesCounter), 10, float32(raylib.GetScreenWidth()), 120)

				if framesCounter >= 120 {
					framesCounter = 0
					state = 4
				}
			}
		case 4: // Fade out animation
			{
				framesCounter++
				alpha = reasings.EaseSineOut(float32(framesCounter), 1.0, -1.0, 160)

				if framesCounter >= 160 {
					framesCounter = 0
					state = 5
				}
			}
		}

		// Reset animation at any moment
		if raylib.IsKeyPressed(raylib.KEY_SPACE) {
			rec = model.Rectangle{X: float32(raylib.GetScreenWidth()) / 2.0, Y: -100, Width: 100, Height: 100}
			rotation = 0.0
			alpha = 1.0
			state = 0
			framesCounter = 0
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawRectanglePro(rec, raylib.Vector2{X: rec.Width / 2, Y: rec.Height / 2}, rotation, raylib.Fade(raylib.BLACK, alpha))

		raylib.DrawText("PRESS [SPACE] TO RESET BOX ANIMATION!", 10, raylib.GetScreenHeight()-25, 20, raylib.LIGHTGRAY)

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
