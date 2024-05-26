package main

import (
	"github.com/JRedrupp/go-raylib/src/bindings/model"
	"github.com/JRedrupp/go-raylib/src/bindings/raylib"
	"github.com/JRedrupp/go-raylib/src/bindings/reasings"
)

const RECS_WIDTH int32 = 50
const RECS_HEIGHT int32 = 50

const MAX_RECS_X int32 = 800 / RECS_WIDTH
const MAX_RECS_Y int32 = 450 / RECS_HEIGHT

const PLAY_TIME_IN_FRAMES int32 = 240 // At 60 fps = 4 seconds

// ------------------------------------------------------------------------------------
// Program main entry point
// ------------------------------------------------------------------------------------
func main() {
	// Initialization
	//--------------------------------------------------------------------------------------
	const screenWidth = 800
	const screenHeight = 450

	raylib.InitWindow(screenWidth, screenHeight, "raylib [shapes] example - easings rectangle array")

	recs := make([]model.Rectangle, MAX_RECS_X*MAX_RECS_Y)
	for y := int32(0); y < MAX_RECS_Y; y++ {
		for x := int32(0); x < MAX_RECS_X; x++ {
			recs[y*MAX_RECS_X+x].X = float32(RECS_WIDTH)/2.0 + float32(RECS_WIDTH*x)
			recs[y*MAX_RECS_X+x].Y = float32(RECS_HEIGHT)/2.0 + float32(RECS_HEIGHT*y)
			recs[y*MAX_RECS_X+x].Width = float32(RECS_WIDTH)
			recs[y*MAX_RECS_X+x].Height = float32(RECS_HEIGHT)
		}
	}

	var rotation float32 = 0.0
	framesCounter := 0
	state := 0 // Rectangles animation state: 0-Playing, 1-Finished

	raylib.SetTargetFPS(60) // Set our game to run at 60 frames-per-second
	//--------------------------------------------------------------------------------------

	// Main game loop
	for !raylib.WindowShouldClose() { // Detect window close button or ESC key
		// Update
		//----------------------------------------------------------------------------------
		if state == 0 {
			framesCounter++

			for i := int32(0); i < MAX_RECS_X*MAX_RECS_Y; i++ {
				recs[i].Height = reasings.EaseCircOut(float32(framesCounter), float32(RECS_HEIGHT), -float32(RECS_HEIGHT), float32(PLAY_TIME_IN_FRAMES))
				recs[i].Width = reasings.EaseCircOut(float32(framesCounter), float32(RECS_WIDTH), -float32(RECS_WIDTH), float32(PLAY_TIME_IN_FRAMES))

				if recs[i].Height < 0 {
					recs[i].Height = 0
				}
				if recs[i].Width < 0 {
					recs[i].Width = 0
				}

				if (recs[i].Height == 0) && (recs[i].Width == 0) {
					state = 1
				} // Finish playing

				rotation = reasings.EaseLinearIn(float32(framesCounter), 0.0, 360.0, float32(PLAY_TIME_IN_FRAMES))
			}
		} else if (state == 1) && raylib.IsKeyPressed(raylib.KEY_SPACE) {
			// When animation has finished, press space to restart
			framesCounter = 0

			for i := int32(0); i < MAX_RECS_X*MAX_RECS_Y; i++ {
				recs[i].Height = float32(RECS_HEIGHT)
				recs[i].Width = float32(RECS_WIDTH)
			}

			state = 0
		}
		//----------------------------------------------------------------------------------

		// Draw
		//----------------------------------------------------------------------------------
		raylib.BeginDrawing()

		raylib.ClearBackground(raylib.RAYWHITE)

		if state == 0 {
			for i := int32(0); i < MAX_RECS_X*MAX_RECS_Y; i++ {
				raylib.DrawRectanglePro(recs[i], raylib.Vector2{X: recs[i].Width / 2, Y: recs[i].Height / 2}, rotation, raylib.RED)
			}
		} else if state == 1 {
			raylib.DrawText("PRESS [SPACE] TO PLAY AGAIN!", 240, 200, 20, raylib.GRAY)
		}

		raylib.EndDrawing()
		//----------------------------------------------------------------------------------
	}

	// De-Initialization
	//--------------------------------------------------------------------------------------
	raylib.CloseWindow() // Close window and OpenGL context
	//--------------------------------------------------------------------------------------
}
