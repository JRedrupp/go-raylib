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

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - raylib logo using shapes")

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		// TODO: Update your variables here
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawRectangle(screenWidth/2-128, screenHeight/2-128, 256, 256, raylib.BLACK)
		raylib.DrawRectangle(screenWidth/2-112, screenHeight/2-112, 224, 224, raylib.RAYWHITE)
		raylib.DrawText("raylib", screenWidth/2-44, screenHeight/2+48, 50, raylib.BLACK)

		raylib.DrawText("this is NOT a texture!", 350, 370, 10, raylib.GRAY)

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
