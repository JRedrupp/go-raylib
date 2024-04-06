package main

import "github.com/JRedrupp/go-raylib/src/bindings/raylib"

const MAX_COLORS_COUNT int = 21

func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth int32 = 800
	const screenHeight int32 = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - colors palette")

	colors := [MAX_COLORS_COUNT]raylib.Color{
		raylib.DARKGRAY, raylib.MAROON, raylib.ORANGE, raylib.DARKGREEN, raylib.DARKBLUE, raylib.DARKPURPLE, raylib.DARKBROWN,
		raylib.GRAY, raylib.RED, raylib.GOLD, raylib.LIME, raylib.BLUE, raylib.VIOLET, raylib.BROWN, raylib.LIGHTGRAY, raylib.PINK, raylib.YELLOW,
		raylib.GREEN, raylib.SKYBLUE, raylib.PURPLE, raylib.BEIGE,
	}

	colorNames := [MAX_COLORS_COUNT]string{
		"DARKGRAY", "MAROON", "ORANGE", "DARKGREEN", "DARKBLUE", "DARKPURPLE",
		"DARKBROWN", "GRAY", "RED", "GOLD", "LIME", "BLUE", "VIOLET", "BROWN",
		"LIGHTGRAY", "PINK", "YELLOW", "GREEN", "SKYBLUE", "PURPLE", "BEIGE",
	}

	colorsRecs := [MAX_COLORS_COUNT]raylib.Rectangle{} // Rectangles array

	// Fills colorsRecs data (for every rectangle)
	for i := 0; i < MAX_COLORS_COUNT; i++ {
		colorsRecs[i].X = 20.0 + 100.0*float32(i%7) + 10.0*float32(i%7)
		colorsRecs[i].Y = 80.0 + 100.0*float32(i/7) + 10.0*float32(i/7)
		colorsRecs[i].Width = 100.0
		colorsRecs[i].Height = 100.0
	}

	colorState := [MAX_COLORS_COUNT]bool{false} // Color state: 0-DEFAULT, 1-MOUSE_HOVER

	var mousePoint raylib.Vector2

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------
		mousePoint = raylib.GetMousePosition()
		for i := 0; i < MAX_COLORS_COUNT; i++ {
			if raylib.CheckCollisionPointRec(mousePoint, colorsRecs[i]) {
				colorState[i] = true
			} else {
				colorState[i] = false
			}
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		raylib.DrawText("raylib colors palette", 28, 42, 20, raylib.BLACK)
		raylib.DrawText("press SPACE to see all colors", raylib.GetScreenWidth()-180, raylib.GetScreenHeight()-40, 10, raylib.GRAY)

		for i := 0; i < MAX_COLORS_COUNT; i++ {
			var opacity float32 = 1.0
			if colorState[i] {
				opacity = 0.6
			}
			raylib.DrawRectangleRec(colorsRecs[i], raylib.Fade(colors[i], opacity))

			if raylib.IsKeyDown(raylib.KEY_SPACE) || colorState[i] {
				raylib.DrawRectangle(int32(colorsRecs[i].X), int32(colorsRecs[i].Y+colorsRecs[i].Height-26), int32(colorsRecs[i].Width), 20, raylib.BLACK)
				raylib.DrawRectangleLinesEx(colorsRecs[i], 6, raylib.Fade(raylib.BLACK, 0.3))
				raylib.DrawText(colorNames[i], int32(colorsRecs[i].X+colorsRecs[i].Width)-raylib.MeasureText(colorNames[i], 10)-12,
					int32(colorsRecs[i].Y+colorsRecs[i].Height-20), 10, colors[i])
			}
		}

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
