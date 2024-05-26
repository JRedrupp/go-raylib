package main

import "github.com/JRedrupp/go-raylib/src/bindings/raylib"

// ------------------------------------------------------------------------------------
// Program main entry point
// ------------------------------------------------------------------------------------
func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	raylib.SetConfigFlags(raylib.FLAG_MSAA_4X_HINT)
	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - cubic-bezier lines")

	startPoint := raylib.Vector2{X: 30, Y: 30}
	endPoint := raylib.Vector2{X: screenWidth - 30, Y: screenHeight - 30}
	moveStartPoint := false
	moveEndPoint := false

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		mouse := raylib.GetMousePosition()

		if raylib.CheckCollisionPointCircle(mouse, startPoint, 10.0) && raylib.IsMouseButtonDown(raylib.MOUSE_BUTTON_LEFT) {
			moveStartPoint = true
		} else if raylib.CheckCollisionPointCircle(mouse, endPoint, 10.0) && raylib.IsMouseButtonDown(raylib.MOUSE_BUTTON_LEFT) {
			moveEndPoint = true
		}

		if moveStartPoint {
			startPoint = mouse
			if raylib.IsMouseButtonReleased(raylib.MOUSE_BUTTON_LEFT) {
				moveStartPoint = false
			}
		}

		if moveEndPoint {
			endPoint = mouse
			if raylib.IsMouseButtonReleased(raylib.MOUSE_BUTTON_LEFT) {
				moveEndPoint = false
			}
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawText("MOVE START-END POINTS WITH MOUSE", 15, 20, 20, raylib.GRAY)

		// Draw line Cubic Bezier, in-out interpolation (easing), no control points
		raylib.DrawLineBezier(startPoint, endPoint, 4.0, raylib.BLUE)

		// Draw start-end spline circles with some details
		var v float32
		var c raylib.Color
		v = 8
		if raylib.CheckCollisionPointCircle(mouse, startPoint, 10.0) {
			v = 14
		}
		c = raylib.BLUE
		if moveStartPoint {
			c = raylib.RED
		}
		raylib.DrawCircleV(startPoint, v, c)

		if raylib.CheckCollisionPointCircle(mouse, endPoint, 10.0) {
			v = 14
		} else {
			v = 8
		}
		if moveEndPoint {
			c = raylib.RED
		} else {
			c = raylib.BLUE
		}
		raylib.DrawCircleV(endPoint, v, c)

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
