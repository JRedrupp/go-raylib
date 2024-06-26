package raylib

/*
#cgo CFLAGS: -g -I../../../raylib/src
#cgo LDFLAGS: -L../../../raylib/src -lraylib -lGL -lm -lpthread -ldl -lrt -lX11
#include <stdlib.h>
#include "raylib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/JRedrupp/go-raylib/src/bindings/model"
)

func rect_to_c(r model.Rectangle) C.Rectangle {
	return C.Rectangle{C.float(r.X), C.float(r.Y), C.float(r.Width), C.float(r.Height)}
}

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

func CheckCollisionPointRec(point Vector2, rec model.Rectangle) bool {
	return bool(C.CheckCollisionPointRec(point.c(), rect_to_c(rec)))
}

// RLAPI bool CheckCollisionPointCircle(Vector2 point, Vector2 center, float radius);                       // Check if point is inside circle
func CheckCollisionPointCircle(point Vector2, center Vector2, radius float32) bool {
	return bool(C.CheckCollisionPointCircle(point.c(), center.c(), C.float(radius)))
}

// RLAPI int GetMouseX(void);                                    // Get mouse position X
func GetMouseX() int32 {
	return int32(C.GetMouseX())
}

// RLAPI int GetMouseY(void);                                    // Get mouse position Y
func GetMouseY() int32 {
	return int32(C.GetMouseY())
}

// RLAPI bool CheckCollisionRecs(Rectangle rec1, model.Rectangle rec2);                                           // Check collision between two model.Rectangles
func CheckCollisionRecs(rec1 model.Rectangle, rec2 model.Rectangle) bool {
	return bool(C.CheckCollisionRecs(rect_to_c(rec1), rect_to_c(rec2)))
}

// RLAPI model.Rectangle GetCollisionRec(Rectangle rec1, model.Rectangle rec2);                                         // Get collision model.Rectangle for two model.Rectangles collision
func GetCollisionRec(rec1 model.Rectangle, rec2 model.Rectangle) model.Rectangle {
	rec := C.GetCollisionRec(rect_to_c(rec1), rect_to_c(rec2))
	return model.Rectangle{X: float32(rec.x), Y: float32(rec.y), Width: float32(rec.width), Height: float32(rec.height)}
}

// RLAPI int MeasureText(const char *text, int fontSize);                                      // Measure string width for default font
func MeasureText(text string, fontSize int32) int32 {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return int32(C.MeasureText(cText, C.int(fontSize)))
}

// RLAPI const char *TextSubtext(const char *text, int position, int length);                  // Get a piece of a text string
func TextSubtext(text string, position int32, length int32) string {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	return C.GoString(C.TextSubtext(cText, C.int(position), C.int(length)))
}

// RLAPI const char *TextFormat(const char *text, ...);                                        // Text formatting with variables (sprintf() style)
func TextFormat(text string, args ...interface{}) string {
	// For now just use fmt.Sprintf
	return fmt.Sprintf(text, args...)

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

// RLAPI bool IsMouseButtonDown(int button);                     // Check if a mouse button is being pressed
func IsMouseButtonDown(button Button) bool {
	return bool(C.IsMouseButtonDown(C.int(button)))
}

type KeyboardKey int

const (
	KEY_NULL = C.KEY_NULL // Key: NULL used for no key pressed
	// Alphanumeric keys
	KEY_APOSTROPHE    = C.KEY_APOSTROPHE    // Key: '
	KEY_COMMA         = C.KEY_COMMA         // Key:
	KEY_MINUS         = C.KEY_MINUS         // Key: -
	KEY_PERIOD        = C.KEY_PERIOD        // Key: .
	KEY_SLASH         = C.KEY_SLASH         // Key: /
	KEY_ZERO          = C.KEY_ZERO          // Key: 0
	KEY_ONE           = C.KEY_ONE           // Key: 1
	KEY_TWO           = C.KEY_TWO           // Key: 2
	KEY_THREE         = C.KEY_THREE         // Key: 3
	KEY_FOUR          = C.KEY_FOUR          // Key: 4
	KEY_FIVE          = C.KEY_FIVE          // Key: 5
	KEY_SIX           = C.KEY_SIX           // Key: 6
	KEY_SEVEN         = C.KEY_SEVEN         // Key: 7
	KEY_EIGHT         = C.KEY_EIGHT         // Key: 8
	KEY_NINE          = C.KEY_NINE          // Key: 9
	KEY_SEMICOLON     = C.KEY_SEMICOLON     // Key: ;
	KEY_EQUAL         = C.KEY_EQUAL         // Key: =
	KEY_A             = C.KEY_A             // Key: A | a
	KEY_B             = C.KEY_B             // Key: B | b
	KEY_C             = C.KEY_C             // Key: C | c
	KEY_D             = C.KEY_D             // Key: D | d
	KEY_E             = C.KEY_E             // Key: E | e
	KEY_F             = C.KEY_F             // Key: F | f
	KEY_G             = C.KEY_G             // Key: G | g
	KEY_H             = C.KEY_H             // Key: H | h
	KEY_I             = C.KEY_I             // Key: I | i
	KEY_J             = C.KEY_J             // Key: J | j
	KEY_K             = C.KEY_K             // Key: K | k
	KEY_L             = C.KEY_L             // Key: L | l
	KEY_M             = C.KEY_M             // Key: M | m
	KEY_N             = C.KEY_N             // Key: N | n
	KEY_O             = C.KEY_O             // Key: O | o
	KEY_P             = C.KEY_P             // Key: P | p
	KEY_Q             = C.KEY_Q             // Key: Q | q
	KEY_R             = C.KEY_R             // Key: R | r
	KEY_S             = C.KEY_S             // Key: S | s
	KEY_T             = C.KEY_T             // Key: T | t
	KEY_U             = C.KEY_U             // Key: U | u
	KEY_V             = C.KEY_V             // Key: V | v
	KEY_W             = C.KEY_W             // Key: W | w
	KEY_X             = C.KEY_X             // Key: X | x
	KEY_Y             = C.KEY_Y             // Key: Y | y
	KEY_Z             = C.KEY_Z             // Key: Z | z
	KEY_LEFT_BRACKET  = C.KEY_LEFT_BRACKET  // Key: [
	KEY_BACKSLASH     = C.KEY_BACKSLASH     // Key: '\'
	KEY_RIGHT_BRACKET = C.KEY_RIGHT_BRACKET // Key: ]
	KEY_GRAVE         = C.KEY_GRAVE         // Key: `
	// Function keys
	KEY_SPACE         = C.KEY_SPACE         // Key: Space
	KEY_ESCAPE        = C.KEY_ESCAPE        // Key: Esc
	KEY_ENTER         = C.KEY_ENTER         // Key: Enter
	KEY_TAB           = C.KEY_TAB           // Key: Tab
	KEY_BACKSPACE     = C.KEY_BACKSPACE     // Key: Backspace
	KEY_INSERT        = C.KEY_INSERT        // Key: Ins
	KEY_DELETE        = C.KEY_DELETE        // Key: Del
	KEY_RIGHT         = C.KEY_RIGHT         // Key: Cursor right
	KEY_LEFT          = C.KEY_LEFT          // Key: Cursor left
	KEY_DOWN          = C.KEY_DOWN          // Key: Cursor down
	KEY_UP            = C.KEY_UP            // Key: Cursor up
	KEY_PAGE_UP       = C.KEY_PAGE_UP       // Key: Page up
	KEY_PAGE_DOWN     = C.KEY_PAGE_DOWN     // Key: Page down
	KEY_HOME          = C.KEY_HOME          // Key: Home
	KEY_END           = C.KEY_END           // Key: End
	KEY_CAPS_LOCK     = C.KEY_CAPS_LOCK     // Key: Caps lock
	KEY_SCROLL_LOCK   = C.KEY_SCROLL_LOCK   // Key: Scroll down
	KEY_NUM_LOCK      = C.KEY_NUM_LOCK      // Key: Num lock
	KEY_PRINT_SCREEN  = C.KEY_PRINT_SCREEN  // Key: Print screen
	KEY_PAUSE         = C.KEY_PAUSE         // Key: Pause
	KEY_F1            = C.KEY_F1            // Key: F1
	KEY_F2            = C.KEY_F2            // Key: F2
	KEY_F3            = C.KEY_F3            // Key: F3
	KEY_F4            = C.KEY_F4            // Key: F4
	KEY_F5            = C.KEY_F5            // Key: F5
	KEY_F6            = C.KEY_F6            // Key: F6
	KEY_F7            = C.KEY_F7            // Key: F7
	KEY_F8            = C.KEY_F8            // Key: F8
	KEY_F9            = C.KEY_F9            // Key: F9
	KEY_F10           = C.KEY_F10           // Key: F10
	KEY_F11           = C.KEY_F11           // Key: F11
	KEY_F12           = C.KEY_F12           // Key: F12
	KEY_LEFT_SHIFT    = C.KEY_LEFT_SHIFT    // Key: Shift left
	KEY_LEFT_CONTROL  = C.KEY_LEFT_CONTROL  // Key: Control left
	KEY_LEFT_ALT      = C.KEY_LEFT_ALT      // Key: Alt left
	KEY_LEFT_SUPER    = C.KEY_LEFT_SUPER    // Key: Super left
	KEY_RIGHT_SHIFT   = C.KEY_RIGHT_SHIFT   // Key: Shift right
	KEY_RIGHT_CONTROL = C.KEY_RIGHT_CONTROL // Key: Control right
	KEY_RIGHT_ALT     = C.KEY_RIGHT_ALT     // Key: Alt right
	KEY_RIGHT_SUPER   = C.KEY_RIGHT_SUPER   // Key: Super right
	KEY_KB_MENU       = C.KEY_KB_MENU       // Key: KB menu
	// Keypad keys
	KEY_KP_0        = C.KEY_KP_0        // Key: Keypad 0
	KEY_KP_1        = C.KEY_KP_1        // Key: Keypad 1
	KEY_KP_2        = C.KEY_KP_2        // Key: Keypad 2
	KEY_KP_3        = C.KEY_KP_3        // Key: Keypad 3
	KEY_KP_4        = C.KEY_KP_4        // Key: Keypad 4
	KEY_KP_5        = C.KEY_KP_5        // Key: Keypad 5
	KEY_KP_6        = C.KEY_KP_6        // Key: Keypad 6
	KEY_KP_7        = C.KEY_KP_7        // Key: Keypad 7
	KEY_KP_8        = C.KEY_KP_8        // Key: Keypad 8
	KEY_KP_9        = C.KEY_KP_9        // Key: Keypad 9
	KEY_KP_DECIMAL  = C.KEY_KP_DECIMAL  // Key: Keypad .
	KEY_KP_DIVIDE   = C.KEY_KP_DIVIDE   // Key: Keypad /
	KEY_KP_MULTIPLY = C.KEY_KP_MULTIPLY // Key: Keypad *
	KEY_KP_SUBTRACT = C.KEY_KP_SUBTRACT // Key: Keypad -
	KEY_KP_ADD      = C.KEY_KP_ADD      // Key: Keypad +
	KEY_KP_ENTER    = C.KEY_KP_ENTER    // Key: Keypad Enter
	KEY_KP_EQUAL    = C.KEY_KP_EQUAL    // Key: Keypad =
	// Android key buttons
	KEY_BACK        = C.KEY_BACK        // Key: Android back button
	KEY_MENU        = C.KEY_MENU        // Key: Android menu button
	KEY_VOLUME_UP   = C.KEY_VOLUME_UP   // Key: Android volume up button
	KEY_VOLUME_DOWN = C.KEY_VOLUME_DOWN // Key: Android volume down button
)

// RLAPI bool IsKeyPressed(int key);                             // Check if a key has been pressed once
func IsKeyPressed(key KeyboardKey) bool {
	return bool(C.IsKeyPressed(C.int(key)))
}

// RLAPI bool IsKeyDown(int key);                                // Check if a key is being pressed
func IsKeyDown(key KeyboardKey) bool {
	return bool(C.IsKeyDown(C.int(key)))
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

// RLAPI void DrawCircleSector(Vector2 center, float radius, float startAngle, float endAngle, int segments, Color color);      // Draw a piece of a circle
func DrawCircleSector(center Vector2, radius float32, startAngle float32, endAngle float32, segments int32, color Color) {
	C.DrawCircleSector(center.c(), C.float(radius), C.float(startAngle), C.float(endAngle), C.int(segments), color.c())
}

// RLAPI void DrawCircleSectorLines(Vector2 center, float radius, float startAngle, float endAngle, int segments, Color color); // Draw circle sector outline
func DrawCircleSectorLines(center Vector2, radius float32, startAngle float32, endAngle float32, segments int32, color Color) {
	C.DrawCircleSectorLines(center.c(), C.float(radius), C.float(startAngle), C.float(endAngle), C.int(segments), color.c())
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

// void DrawRectangle(int posX, int posY, int width, int height, Color color);                        // Draw a color-filled model.Rectangle
func DrawRectangle(posX int32, posY int32, width int32, height int32, color Color) {
	C.DrawRectangle(C.int(posX), C.int(posY), C.int(width), C.int(height), color.c())
}

// RLAPI void DrawRectanglePro(Rectangle rec, Vector2 origin, float rotation, Color color);                 // Draw a color-filled rectangle with pro parameters
func DrawRectanglePro(rec model.Rectangle, origin Vector2, rotation float32, color Color) {
	C.DrawRectanglePro(rect_to_c(rec), origin.c(), C.float(rotation), color.c())
}

// void DrawRectangleGradientH(int posX, int posY, int width, int height, Color color1, Color color2);// Draw a horizontal-gradient-filled model.Rectangle
func DrawRectangleGradientH(posX int32, posY int32, width int32, height int32, color1 Color, color2 Color) {
	C.DrawRectangleGradientH(C.int(posX), C.int(posY), C.int(width), C.int(height), color1.c(), color2.c())
}

// RLAPI void DrawRectangleLines(int posX, int posY, int width, int height, Color color);                   // Draw model.Rectangle outline
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

// RLAPI void DrawLineBezier(Vector2 startPos, Vector2 endPos, float thick, Color color);                   // Draw line segment cubic-bezier in-out interpolation
func DrawLineBezier(startPos Vector2, endPos Vector2, thick float32, color Color) {
	C.DrawLineBezier(startPos.c(), endPos.c(), C.float(thick), color.c())
}

func DrawRectangleRec(rec model.Rectangle, color Color) {
	C.DrawRectangleRec(rect_to_c(rec), color.c())
}

func DrawRectangleLinesEx(rec model.Rectangle, lineThick float32, color Color) {
	C.DrawRectangleLinesEx(rect_to_c(rec), C.float(lineThick), color.c())
}

// RLAPI void DrawRectangleRounded(Rectangle rec, float roundness, int segments, Color color);              // Draw model.Rectangle with rounded edges
func DrawRectangleRounded(rec model.Rectangle, roundness float32, segments int32, color Color) {
	C.DrawRectangleRounded(rect_to_c(rec), C.float(roundness), C.int(segments), color.c())
}

// RLAPI void DrawRectangleRoundedLines(Rectangle rec, float roundness, int segments, Color color);         // Draw rectangle lines with rounded edges
func DrawRectangleRoundedLines(rec model.Rectangle, roundness float32, segments int32, color Color) {
	C.DrawRectangleRoundedLines(rect_to_c(rec), C.float(roundness), C.int(segments), color.c())
}

func DrawTriangle(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	C.DrawTriangle(v1.c(), v2.c(), v3.c(), color.c())
}

// RLAPI void DrawRing(Vector2 center, float innerRadius, float outerRadius, float startAngle, float endAngle, int segments, Color color); // Draw ring
func DrawRing(center Vector2, innerRadius float32, outerRadius float32, startAngle float32, endAngle float32, segments int32, color Color) {
	C.DrawRing(center.c(), C.float(innerRadius), C.float(outerRadius), C.float(startAngle), C.float(endAngle), C.int(segments), color.c())
}

// RLAPI void DrawRingLines(Vector2 center, float innerRadius, float outerRadius, float startAngle, float endAngle, int segments, Color color);    // Draw ring outline
func DrawRingLines(center Vector2, innerRadius float32, outerRadius float32, startAngle float32, endAngle float32, segments int32, color Color) {
	C.DrawRingLines(center.c(), C.float(innerRadius), C.float(outerRadius), C.float(startAngle), C.float(endAngle), C.int(segments), color.c())
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

// Func to convert from C Color to Go Color
func colorFromC(c C.struct_Color) Color {
	return Color{r: uint8(c.r), g: uint8(c.g), b: uint8(c.b), a: uint8(c.a)}
}

// Raylib Colors
var (
	LIGHTGRAY  = colorFromC(C.LIGHTGRAY)  // Light Gray
	GRAY       = colorFromC(C.GRAY)       // Gray
	DARKGRAY   = colorFromC(C.DARKGRAY)   // Dark Gray
	YELLOW     = colorFromC(C.YELLOW)     // Yellow
	GOLD       = colorFromC(C.GOLD)       // Gold
	ORANGE     = colorFromC(C.ORANGE)     // Orange
	PINK       = colorFromC(C.PINK)       // Pink
	RED        = colorFromC(C.RED)        // Red
	MAROON     = colorFromC(C.MAROON)     // Maroon
	GREEN      = colorFromC(C.GREEN)      // Green
	LIME       = colorFromC(C.LIME)       // Lime
	DARKGREEN  = colorFromC(C.DARKGREEN)  // Dark Green
	SKYBLUE    = colorFromC(C.SKYBLUE)    // Sky Blue
	BLUE       = colorFromC(C.BLUE)       // Blue
	DARKBLUE   = colorFromC(C.DARKBLUE)   // Dark Blue
	PURPLE     = colorFromC(C.PURPLE)     // Purple
	VIOLET     = colorFromC(C.VIOLET)     // Violet
	DARKPURPLE = colorFromC(C.DARKPURPLE) // Dark Purple
	BEIGE      = colorFromC(C.BEIGE)      // Beige
	BROWN      = colorFromC(C.BROWN)      // Brown
	DARKBROWN  = colorFromC(C.DARKBROWN)  // Dark Brown
	WHITE      = colorFromC(C.WHITE)      // White
	BLACK      = colorFromC(C.BLACK)      // Black
	BLANK      = colorFromC(C.BLANK)      // Blank (Transparent)
	MAGENTA    = colorFromC(C.MAGENTA)    // Magenta
	RAYWHITE   = colorFromC(C.RAYWHITE)   // My own White (raylib logo)
)

// ClearBackground - Clear background with color
func ClearBackground(color Color) {
	C.ClearBackground(color.c())
}
