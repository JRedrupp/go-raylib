package main

import "github.com/JRedrupp/go-raylib/src/raylib"

const MOUSE_SCALE_MARK_SIZE int = 12

func main() {

	//------------------------------------------------------------------------------------
	// Program main entry point
	//------------------------------------------------------------------------------------
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	raylib.InitWindow(screenWidth, screenHeight, "go-raylib [shapes] example - rectangle scaling mouse")

	rec := raylib.Rectangle{X: 100, Y: 100, Width: 200, Height: 80}

	var mousePosition = raylib.Vector2{X: 0, Y: 0}

	mouseScaleReady := false
	mouseScaleMode := false

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		mousePosition = raylib.GetMousePosition()

		if (raylib.CheckCollisionPointRec(mousePosition, raylib.Rectangle{X: rec.X + rec.Width - float32(MOUSE_SCALE_MARK_SIZE), Y: rec.Y + rec.Height - float32(MOUSE_SCALE_MARK_SIZE), Width: float32(MOUSE_SCALE_MARK_SIZE), Height: float32(MOUSE_SCALE_MARK_SIZE)})) {
			mouseScaleReady = true
			if raylib.IsMouseButtonPressed(raylib.MOUSE_BUTTON_LEFT) {
				mouseScaleMode = true
			}
		} else {
			mouseScaleReady = false
		}

		if mouseScaleMode {
			mouseScaleReady = true

			rec.Width = (mousePosition.X - rec.X)
			rec.Height = (mousePosition.Y - rec.Y)

			// Check minimum rec size
			if rec.Width < float32(MOUSE_SCALE_MARK_SIZE) {
				rec.Width = float32(MOUSE_SCALE_MARK_SIZE)
			}
			if rec.Height < float32(MOUSE_SCALE_MARK_SIZE) {
				rec.Height = float32(MOUSE_SCALE_MARK_SIZE)
			}

			// Check maximum rec size
			if rec.Width > (float32(raylib.GetScreenWidth()) - rec.X) {
				rec.Width = float32(raylib.GetScreenWidth()) - rec.X
			}
			if rec.Height > (float32(raylib.GetScreenHeight()) - rec.Y) {
				rec.Height = float32(raylib.GetScreenHeight()) - rec.Y
			}

			if raylib.IsMouseButtonReleased(raylib.MOUSE_BUTTON_LEFT) {
				mouseScaleMode = false
			}
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawText("Scale rectangle dragging from bottom-right corner!", 10, 10, 20, raylib.GRAY)

		raylib.DrawRectangleRec(rec, raylib.Fade(raylib.GREEN, 0.5))

		if mouseScaleReady {
			raylib.DrawRectangleLinesEx(rec, 1, raylib.RED)
			raylib.DrawTriangle(raylib.Vector2{X: rec.X + rec.Width - float32(MOUSE_SCALE_MARK_SIZE), Y: rec.Y + rec.Height},
				raylib.Vector2{X: rec.X + rec.Width, Y: rec.Y + rec.Height},
				raylib.Vector2{X: rec.X + rec.Width, Y: rec.Y + rec.Height - float32(MOUSE_SCALE_MARK_SIZE)}, raylib.RED)
		}

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
