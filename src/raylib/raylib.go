package raylib

/*
#cgo CFLAGS: -g -I../../raylib/src
#cgo LDFLAGS: -L../../raylib/src -lraylib -lGL -lm -lpthread -ldl -lrt -lX11
#include <stdlib.h>
#include "raylib.h"
*/
import "C"
import "unsafe"

// InitWindow - Initialize window and OpenGL context
func InitWindow(width int32, height int32, title string) {

	// TODO: Store Strings in a map to prevent having to malloc and free every time
	cTitle := C.CString(title)
	defer C.free(unsafe.Pointer(cTitle))

	C.InitWindow(C.int(width), C.int(height), cTitle)
}

// WindowShouldClose - bool WindowShouldClose(void);
func WindowShouldClose() bool {
	return bool(C.WindowShouldClose())
}

// CloseWindow - Close window and unload OpenGL context
func CloseWindow() {
	C.CloseWindow()
}

// BeginDrawing - Setup canvas (framebuffer) to start drawing
func BeginDrawing() {
	C.BeginDrawing()
}

func EndDrawing() {
	C.EndDrawing()
}

func SetTargetFPS(fps int) {
	C.SetTargetFPS(C.int(fps))
}

func GetMousePosition() Vector2 {
	pos := C.GetMousePosition()
	return Vector2{float32(pos.x), float32(pos.y)}
}

func CheckCollisionPointRec(point Vector2, rec Rectangle) bool {
	return bool(C.CheckCollisionPointRec(point.c(), rec.c()))
}

type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

func (r Rectangle) c() C.struct_Rectangle {
	return C.struct_Rectangle{C.float(r.X), C.float(r.Y), C.float(r.Width), C.float(r.Height)}
}

type ConfigFlags int

const (
	FLAG_VSYNC_HINT               ConfigFlags = 0x00000040 // Set to try enabling V-Sync on GPU
	FLAG_FULLSCREEN_MODE          ConfigFlags = 0x00000002 // Set to run program in fullscreen
	FLAG_WINDOW_RESIZABLE         ConfigFlags = 0x00000004 // Set to allow resizable window
	FLAG_WINDOW_UNDECORATED       ConfigFlags = 0x00000008 // Set to disable window decoration (frame and buttons)
	FLAG_WINDOW_HIDDEN            ConfigFlags = 0x00000080 // Set to hide window
	FLAG_WINDOW_MINIMIZED         ConfigFlags = 0x00000200 // Set to minimize window (iconify)
	FLAG_WINDOW_MAXIMIZED         ConfigFlags = 0x00000400 // Set to maximize window (expanded to monitor)
	FLAG_WINDOW_UNFOCUSED         ConfigFlags = 0x00000800 // Set to window non focused
	FLAG_WINDOW_TOPMOST           ConfigFlags = 0x00001000 // Set to window always on top
	FLAG_WINDOW_ALWAYS_RUN        ConfigFlags = 0x00000100 // Set to allow windows running while minimized
	FLAG_WINDOW_TRANSPARENT       ConfigFlags = 0x00000010 // Set to allow transparent framebuffer
	FLAG_WINDOW_HIGHDPI           ConfigFlags = 0x00002000 // Set to support HighDPI
	FLAG_WINDOW_MOUSE_PASSTHROUGH ConfigFlags = 0x00004000 // Set to support mouse passthrough only supported when FLAG_WINDOW_UNDECORATED
	FLAG_BORDERLESS_WINDOWED_MODE ConfigFlags = 0x00008000 // Set to run program in borderless windowed mode
	FLAG_MSAA_4X_HINT             ConfigFlags = 0x00000020 // Set to try enabling MSAA 4X
	FLAG_INTERLACED_HINT          ConfigFlags = 0x00010000 // Set to try enabling interlaced video format (for V3D)

)

// RLAPI void SetConfigFlags(unsigned int flags);                    // Setup init configuration flags (view FLAGS)
func SetConfigFlags(flags ConfigFlags) {
	C.SetConfigFlags(C.uint(flags))
}

type Button int

const (
	MOUSE_BUTTON_LEFT Button = iota
	MOUSE_BUTTON_RIGHT
	MOUSE_BUTTON_MIDDLE
	MOUSE_BUTTON_SIDE    // Mouse button side (advanced mouse device)
	MOUSE_BUTTON_EXTRA   // Mouse button extra (advanced mouse device)
	MOUSE_BUTTON_FORWARD // Mouse button forward (advanced mouse device)
	MOUSE_BUTTON_BACK
)

func IsMouseButtonPressed(button Button) bool {
	return bool(C.IsMouseButtonPressed(C.int(button)))
}

type KeyboardKey int

const (
	KEY_NULL = 0 // Key: NULL used for no key pressed
	// Alphanumeric keys
	KEY_APOSTROPHE    = 39 // Key: '
	KEY_COMMA         = 44 // Key:
	KEY_MINUS         = 45 // Key: -
	KEY_PERIOD        = 46 // Key: .
	KEY_SLASH         = 47 // Key: /
	KEY_ZERO          = 48 // Key: 0
	KEY_ONE           = 49 // Key: 1
	KEY_TWO           = 50 // Key: 2
	KEY_THREE         = 51 // Key: 3
	KEY_FOUR          = 52 // Key: 4
	KEY_FIVE          = 53 // Key: 5
	KEY_SIX           = 54 // Key: 6
	KEY_SEVEN         = 55 // Key: 7
	KEY_EIGHT         = 56 // Key: 8
	KEY_NINE          = 57 // Key: 9
	KEY_SEMICOLON     = 59 // Key: ;
	KEY_EQUAL         = 61 // Key: =
	KEY_A             = 65 // Key: A | a
	KEY_B             = 66 // Key: B | b
	KEY_C             = 67 // Key: C | c
	KEY_D             = 68 // Key: D | d
	KEY_E             = 69 // Key: E | e
	KEY_F             = 70 // Key: F | f
	KEY_G             = 71 // Key: G | g
	KEY_H             = 72 // Key: H | h
	KEY_I             = 73 // Key: I | i
	KEY_J             = 74 // Key: J | j
	KEY_K             = 75 // Key: K | k
	KEY_L             = 76 // Key: L | l
	KEY_M             = 77 // Key: M | m
	KEY_N             = 78 // Key: N | n
	KEY_O             = 79 // Key: O | o
	KEY_P             = 80 // Key: P | p
	KEY_Q             = 81 // Key: Q | q
	KEY_R             = 82 // Key: R | r
	KEY_S             = 83 // Key: S | s
	KEY_T             = 84 // Key: T | t
	KEY_U             = 85 // Key: U | u
	KEY_V             = 86 // Key: V | v
	KEY_W             = 87 // Key: W | w
	KEY_X             = 88 // Key: X | x
	KEY_Y             = 89 // Key: Y | y
	KEY_Z             = 90 // Key: Z | z
	KEY_LEFT_BRACKET  = 91 // Key: [
	KEY_BACKSLASH     = 92 // Key: '\'
	KEY_RIGHT_BRACKET = 93 // Key: ]
	KEY_GRAVE         = 96 // Key: `
	// Function keys
	KEY_SPACE         = 32  // Key: Space
	KEY_ESCAPE        = 256 // Key: Esc
	KEY_ENTER         = 257 // Key: Enter
	KEY_TAB           = 258 // Key: Tab
	KEY_BACKSPACE     = 259 // Key: Backspace
	KEY_INSERT        = 260 // Key: Ins
	KEY_DELETE        = 261 // Key: Del
	KEY_RIGHT         = 262 // Key: Cursor right
	KEY_LEFT          = 263 // Key: Cursor left
	KEY_DOWN          = 264 // Key: Cursor down
	KEY_UP            = 265 // Key: Cursor up
	KEY_PAGE_UP       = 266 // Key: Page up
	KEY_PAGE_DOWN     = 267 // Key: Page down
	KEY_HOME          = 268 // Key: Home
	KEY_END           = 269 // Key: End
	KEY_CAPS_LOCK     = 280 // Key: Caps lock
	KEY_SCROLL_LOCK   = 281 // Key: Scroll down
	KEY_NUM_LOCK      = 282 // Key: Num lock
	KEY_PRINT_SCREEN  = 283 // Key: Print screen
	KEY_PAUSE         = 284 // Key: Pause
	KEY_F1            = 290 // Key: F1
	KEY_F2            = 291 // Key: F2
	KEY_F3            = 292 // Key: F3
	KEY_F4            = 293 // Key: F4
	KEY_F5            = 294 // Key: F5
	KEY_F6            = 295 // Key: F6
	KEY_F7            = 296 // Key: F7
	KEY_F8            = 297 // Key: F8
	KEY_F9            = 298 // Key: F9
	KEY_F10           = 299 // Key: F10
	KEY_F11           = 300 // Key: F11
	KEY_F12           = 301 // Key: F12
	KEY_LEFT_SHIFT    = 340 // Key: Shift left
	KEY_LEFT_CONTROL  = 341 // Key: Control left
	KEY_LEFT_ALT      = 342 // Key: Alt left
	KEY_LEFT_SUPER    = 343 // Key: Super left
	KEY_RIGHT_SHIFT   = 344 // Key: Shift right
	KEY_RIGHT_CONTROL = 345 // Key: Control right
	KEY_RIGHT_ALT     = 346 // Key: Alt right
	KEY_RIGHT_SUPER   = 347 // Key: Super right
	KEY_KB_MENU       = 348 // Key: KB menu
	// Keypad keys
	KEY_KP_0        = 320 // Key: Keypad 0
	KEY_KP_1        = 321 // Key: Keypad 1
	KEY_KP_2        = 322 // Key: Keypad 2
	KEY_KP_3        = 323 // Key: Keypad 3
	KEY_KP_4        = 324 // Key: Keypad 4
	KEY_KP_5        = 325 // Key: Keypad 5
	KEY_KP_6        = 326 // Key: Keypad 6
	KEY_KP_7        = 327 // Key: Keypad 7
	KEY_KP_8        = 328 // Key: Keypad 8
	KEY_KP_9        = 329 // Key: Keypad 9
	KEY_KP_DECIMAL  = 330 // Key: Keypad .
	KEY_KP_DIVIDE   = 331 // Key: Keypad /
	KEY_KP_MULTIPLY = 332 // Key: Keypad *
	KEY_KP_SUBTRACT = 333 // Key: Keypad -
	KEY_KP_ADD      = 334 // Key: Keypad +
	KEY_KP_ENTER    = 335 // Key: Keypad Enter
	KEY_KP_EQUAL    = 336 // Key: Keypad =
	// Android key buttons
	KEY_BACK        = 4  // Key: Android back button
	KEY_MENU        = 5  // Key: Android menu button
	KEY_VOLUME_UP   = 24 // Key: Android volume up button
	KEY_VOLUME_DOWN = 25 // Key: Android volume down button
)

// RLAPI bool IsKeyPressed(int key);                             // Check if a key has been pressed once
func IsKeyPressed(key KeyboardKey) bool {
	return bool(C.IsKeyPressed(C.int(key)))
}

func GetScreenWidth() int32 {
	return int32(C.GetScreenWidth())
}

// RLAPI void DrawFPS(int posX, int posY);                                                     // Draw current FPS
func DrawFPS(posX int32, posY int32) {
	C.DrawFPS(C.int(posX), C.int(posY))
}

// DrawCircle(int centerX, int centerY, float radius, Color color)
func DrawCircle(centerX int32, centerY int32, radius float32, color Color) {
	C.DrawCircle(C.int(centerX), C.int(centerY), C.float(radius), color.c())
}

// RLAPI void DrawCircleV(Vector2 center, float radius, Color color);                                       // Draw a color-filled circle (Vector version)
func DrawCircleV(center Vector2, radius float32, color Color) {
	C.DrawCircleV(center.c(), C.float(radius), color.c())
}

// void DrawCircleGradient(int centerX, int centerY, float radius, Color color1, Color color2);       // Draw a gradient-filled circle
func DrawCircleGradient(centerX int32, centerY int32, radius float32, color1 Color, color2 Color) {
	C.DrawCircleGradient(C.int(centerX), C.int(centerY), C.float(radius), color1.c(), color2.c())
}

// void DrawCircleLines(int centerX, int centerY, float radius, Color color);                         // Draw circle outline
func DrawCircleLines(centerX int32, centerY int32, radius float32, color Color) {
	C.DrawCircleLines(C.int(centerX), C.int(centerY), C.float(radius), color.c())
}

func GetScreenHeight() int32 {
	return int32(C.GetScreenHeight())
}

func IsMouseButtonReleased(button Button) bool {
	return bool(C.IsMouseButtonReleased(C.int(button)))
}

func DrawText(text string, x int32, y int32, fontSize int32, color Color) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	C.DrawText(cText, C.int(x), C.int(y), C.int(fontSize), color.c())
}

// void DrawRectangle(int posX, int posY, int width, int height, Color color);                        // Draw a color-filled rectangle
func DrawRectangle(posX int32, posY int32, width int32, height int32, color Color) {
	C.DrawRectangle(C.int(posX), C.int(posY), C.int(width), C.int(height), color.c())
}

// void DrawRectangleGradientH(int posX, int posY, int width, int height, Color color1, Color color2);// Draw a horizontal-gradient-filled rectangle
func DrawRectangleGradientH(posX int32, posY int32, width int32, height int32, color1 Color, color2 Color) {
	C.DrawRectangleGradientH(C.int(posX), C.int(posY), C.int(width), C.int(height), color1.c(), color2.c())
}

// RLAPI void DrawRectangleLines(int posX, int posY, int width, int height, Color color);                   // Draw rectangle outline
func DrawRectangleLines(posX int32, posY int32, width int32, height int32, color Color) {
	C.DrawRectangleLines(C.int(posX), C.int(posY), C.int(width), C.int(height), color.c())
}

// RLAPI void DrawTriangleLines(Vector2 v1, Vector2 v2, Vector2 v3, Color color);                           // Draw triangle outline (vertex in counter-clockwise order!)
func DrawTriangleLines(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	C.DrawTriangleLines(v1.c(), v2.c(), v3.c(), color.c())
}

// RLAPI void DrawPoly(Vector2 center, int sides, float radius, float rotation, Color color);               // Draw a regular polygon (Vector version)
func DrawPoly(center Vector2, sides int32, radius float32, rotation float32, color Color) {
	C.DrawPoly(center.c(), C.int(sides), C.float(radius), C.float(rotation), color.c())
}

// RLAPI void DrawPolyLines(Vector2 center, int sides, float radius, float rotation, Color color);          // Draw a polygon outline of n sides
func DrawPolyLines(center Vector2, sides int32, radius float32, rotation float32, color Color) {
	C.DrawPolyLines(center.c(), C.int(sides), C.float(radius), C.float(rotation), color.c())
}

// RLAPI void DrawPolyLinesEx(Vector2 center, int sides, float radius, float rotation, float lineThick, Color color); // Draw a polygon outline of n sides with extended parameters
func DrawPolyLinesEx(center Vector2, sides int32, radius float32, rotation float32, lineThick float32, color Color) {
	C.DrawPolyLinesEx(center.c(), C.int(sides), C.float(radius), C.float(rotation), C.float(lineThick), color.c())
}

// RLAPI void DrawLine(int startPosX, int startPosY, int endPosX, int endPosY, Color color);                // Draw a line
func DrawLine(startPosX int32, startPosY int32, endPosX int32, endPosY int32, color Color) {
	C.DrawLine(C.int(startPosX), C.int(startPosY), C.int(endPosX), C.int(endPosY), color.c())
}

func DrawRectangleRec(rec Rectangle, color Color) {
	C.DrawRectangleRec(rec.c(), color.c())
}

func DrawRectangleLinesEx(rec Rectangle, lineThick float32, color Color) {
	C.DrawRectangleLinesEx(rec.c(), C.float(lineThick), color.c())
}

func DrawTriangle(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	C.DrawTriangle(v1.c(), v2.c(), v3.c(), color.c())
}

// Easy enough to do internally
func Fade(color Color, alpha float32) Color {
	return Color{r: color.r, g: color.g, b: color.b, a: uint8(alpha * 255)}
}

type Vector2 struct {
	X float32
	Y float32
}

func (v Vector2) c() C.struct_Vector2 {
	return C.struct_Vector2{C.float(v.X), C.float(v.Y)}
}

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

func (c Color) c() C.struct_Color {
	return C.struct_Color{C.uchar(c.r), C.uchar(c.g), C.uchar(c.b), C.uchar(c.a)}
}

// Raylib Colors
var (
	LIGHTGRAY  = Color{r: 200, g: 200, b: 200, a: 255} // Light Gray
	GRAY       = Color{r: 130, g: 130, b: 130, a: 255} // Gray
	DARKGRAY   = Color{r: 80, g: 80, b: 80, a: 255}    // Dark Gray
	YELLOW     = Color{r: 253, g: 249, b: 0, a: 255}   // Yellow
	GOLD       = Color{r: 255, g: 203, b: 0, a: 255}   // Gold
	ORANGE     = Color{r: 255, g: 161, b: 0, a: 255}   // Orange
	PINK       = Color{r: 255, g: 109, b: 194, a: 255} // Pink
	RED        = Color{r: 230, g: 41, b: 55, a: 255}   // Red
	MAROON     = Color{r: 190, g: 33, b: 55, a: 255}   // Maroon
	GREEN      = Color{r: 0, g: 228, b: 48, a: 255}    // Green
	LIME       = Color{r: 0, g: 158, b: 47, a: 255}    // Lime
	DARKGREEN  = Color{r: 0, g: 117, b: 44, a: 255}    // Dark Green
	SKYBLUE    = Color{r: 102, g: 191, b: 255, a: 255} // Sky Blue
	BLUE       = Color{r: 0, g: 121, b: 241, a: 255}   // Blue
	DARKBLUE   = Color{r: 0, g: 82, b: 172, a: 255}    // Dark Blue
	PURPLE     = Color{r: 200, g: 122, b: 255, a: 255} // Purple
	VIOLET     = Color{r: 135, g: 60, b: 190, a: 255}  // Violet
	DARKPURPLE = Color{r: 112, g: 31, b: 126, a: 255}  // Dark Purple
	BEIGE      = Color{r: 211, g: 176, b: 131, a: 255} // Beige
	BROWN      = Color{r: 127, g: 106, b: 79, a: 255}  // Brown
	DARKBROWN  = Color{r: 76, g: 63, b: 47, a: 255}    // Dark Brown
	WHITE      = Color{r: 255, g: 255, b: 255, a: 255} // White
	BLACK      = Color{r: 0, g: 0, b: 0, a: 255}       // Black
	BLANK      = Color{r: 0, g: 0, b: 0, a: 0}         // Blank (Transparent)
	MAGENTA    = Color{r: 255, g: 0, b: 255, a: 255}   // Magenta
	RAYWHITE   = Color{r: 245, g: 245, b: 245, a: 255} // My own White (raylib logo)
)

// ClearBackground - Clear background with color
func ClearBackground(color Color) {
	C.ClearBackground(color.c())
}