package main

import (
	"github.com/JRedrupp/go-raylib/src/bindings/model"
	"github.com/JRedrupp/go-raylib/src/bindings/raylib"
)

func main() {
	// Initialization
	//---------------------------------------------------------
	const screenWidth int32 = 800
	const screenHeight int32 = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - collision area")

	// Box A: Moving box
	boxA := model.Rectangle{X: 10, Y: float32(raylib.GetScreenHeight()/2.0 - 50), Width: 200, Height: 100}
	boxASpeedX := 4

	// Box B: Mouse moved box
	boxB := model.Rectangle{X: float32(raylib.GetScreenWidth()/2.0 - 30), Y: float32(raylib.GetScreenHeight()/2.0 - 30), Width: 60, Height: 60}

	boxCollision := model.Rectangle{X: 0, Y: 0, Width: 0, Height: 0} // Collision model.Rectangle

	screenUpperLimit := int32(40) // Top menu limits

	pause := false     // Movement pause
	collision := false // Collision detection

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//----------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update
		//-----------------------------------------------------
		// Move box if not paused
		if !pause {
			boxA.X += float32(boxASpeedX)
		}

		// Bounce box on x screen limits
		if ((boxA.X + boxA.Width) >= float32(raylib.GetScreenWidth())) || (boxA.X <= 0) {
			boxASpeedX *= -1
		}

		// Update player-controlled-box (box02)
		boxB.X = float32(raylib.GetMouseX()) - boxB.Width/2
		boxB.Y = float32(raylib.GetMouseY()) - boxB.Height/2

		// Make sure Box B does not go out of move area limits
		if (boxB.X + boxB.Width) >= float32(raylib.GetScreenWidth()) {
			boxB.X = float32(raylib.GetScreenWidth()) - boxB.Width
		} else if boxB.X <= 0 {
			boxB.X = 0
		}

		if (boxB.Y + boxB.Height) >= float32(raylib.GetScreenHeight()) {
			boxB.Y = float32(raylib.GetScreenHeight()) - boxB.Height
		} else if boxB.Y <= float32(screenUpperLimit) {
			boxB.Y = float32(screenUpperLimit)
		}

		// Check boxes collision
		collision = raylib.CheckCollisionRecs(boxA, boxB)

		// Get collision model.Rectangle (only on collision)
		if collision {
			boxCollision = raylib.GetCollisionRec(boxA, boxB)
		}

		// Pause Box A movement
		if raylib.IsKeyPressed(raylib.KEY_SPACE) {
			pause = !pause
		}
		//-----------------------------------------------------

		// Draw
		//-----------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		col := raylib.BLACK
		if collision {
			col = raylib.RED
		}
		raylib.DrawRectangle(0, 0, screenWidth, screenUpperLimit, col)

		raylib.DrawRectangleRec(boxA, raylib.GOLD)
		raylib.DrawRectangleRec(boxB, raylib.BLUE)

		if collision {
			// Draw collision area
			raylib.DrawRectangleRec(boxCollision, raylib.LIME)

			// Draw collision message
			raylib.DrawText("COLLISION!", raylib.GetScreenWidth()/2-raylib.MeasureText("COLLISION!", 20)/2, screenUpperLimit/2-10, 20, raylib.BLACK)

			// Draw collision area
			raylib.DrawText(raylib.TextFormat("Collision Area: %f", boxCollision.Width*boxCollision.Height), raylib.GetScreenWidth()/2-100, screenUpperLimit+10, 20, raylib.BLACK)
		}

		// Draw help instructions
		raylib.DrawText("Press SPACE to PAUSE/RESUME", 20, screenHeight-35, 20, raylib.LIGHTGRAY)

		raylib.DrawFPS(10, 10)

		raylib.EndDrawing()
		//-----------------------------------------------------
	}

	// De-Initialization
	//---------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//----------------------------------------------------------
}
