package raygui

/*
#cgo CFLAGS: -g -I../../../raylib/src -I../../../raylib/examples/shapes
#cgo LDFLAGS: -L../../../raylib/src -lraylib -lGL -lm -lpthread -ldl -lrt -lX11
#include <stdlib.h>
#define RAYGUI_IMPLEMENTATION
#include "raygui.h"
*/
import "C"
import (
	"unsafe"

	"github.com/JRedrupp/go-raylib/src/bindings/raylib"
)

// RAYGUIAPI int GuiSliderBar(Rectangle bounds, const char *textLeft, const char *textRight, float *value, float minValue, float maxValue); // Slider Bar control, returns selected value
func GuiSliderBar(bounds raylib.Rectangle, textLeft string, textRight string, value *float32, minValue float32, maxValue float32) int {

	cTextLeft := C.CString(textLeft)
	defer C.free(unsafe.Pointer(cTextLeft))
	cTextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(cTextRight))
	cValue := (*C.float)(unsafe.Pointer(value))
	cMinValue := (C.float)(minValue)
	cMaxValue := (C.float)(maxValue)

	// TODO: Work out how we can use the .C() method on the raylib.Rectangle struct
	cBounds := C.struct_Rectangle{
		x:      (C.float)(bounds.X),
		y:      (C.float)(bounds.Y),
		width:  (C.float)(bounds.Width),
		height: (C.float)(bounds.Height),
	}

	return int(C.GuiSliderBar(cBounds, cTextLeft, cTextRight, cValue, cMinValue, cMaxValue))
}
