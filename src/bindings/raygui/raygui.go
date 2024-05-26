package raygui

/*
#cgo CFLAGS: -g -I../../../raylib/src -I../../../raylib/examples/shapes -w
#cgo LDFLAGS: -L../../../raylib/src -lraylib -lGL -lm -lpthread -ldl -lrt -lX11
#include <stdlib.h>
#define RAYGUI_IMPLEMENTATION
#include "raygui.h"
*/
import "C"
import (
	"unsafe"

	"github.com/JRedrupp/go-raylib/src/bindings/model"
)

func rect_to_c(r model.Rectangle) C.Rectangle {
	return C.Rectangle{C.float(r.X), C.float(r.Y), C.float(r.Width), C.float(r.Height)}
}

// RAYGUIAPI int GuiSliderBar(Rectangle bounds, const char *textLeft, const char *textRight, float *value, float minValue, float maxValue); // Slider Bar control, returns selected value
func GuiSliderBar(bounds model.Rectangle, textLeft string, textRight string, value *float32, minValue float32, maxValue float32) int {

	cTextLeft := C.CString(textLeft)
	defer C.free(unsafe.Pointer(cTextLeft))
	cTextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(cTextRight))
	cValue := (*C.float)(unsafe.Pointer(value))
	cMinValue := (C.float)(minValue)
	cMaxValue := (C.float)(maxValue)

	return int(C.GuiSliderBar(rect_to_c(bounds), cTextLeft, cTextRight, cValue, cMinValue, cMaxValue))
}

// RAYGUIAPI int GuiCheckBox(Rectangle bounds, const char *text, bool *checked);                          // Check Box control, returns true when active
func GuiCheckBox(bounds model.Rectangle, text string, checked *bool) int {

	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))
	cChecked := (*C.bool)(unsafe.Pointer(checked))

	return int(C.GuiCheckBox(rect_to_c(bounds), cText, cChecked))
}
