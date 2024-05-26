package reasings

/*
#cgo CFLAGS: -I../../../raylib/examples/shapes
#include "reasings.h"
*/
import "C"

// EASEDEF float EaseElasticIn(float t, float b, float c, float d) // Ease: Elastic In
func EaseElasticIn(t float32, b float32, c float32, d float32) float32 {
	return float32(C.EaseElasticIn(C.float(t), C.float(b), C.float(c), C.float(d)))
}

// EASEDEF float EaseElasticOut(float t, float b, float c, float d) // Ease: Elastic Out
func EaseElasticOut(t float32, b float32, c float32, d float32) float32 {
	return float32(C.EaseElasticOut(C.float(t), C.float(b), C.float(c), C.float(d)))
}

// EASEDEF float EaseCubicOut(float t, float b, float c, float d) { t = t/d - 1.0f; return (c*(t*t*t + 1.0f) + b); }    // Ease: Cubic Out
func EaseCubicOut(t float32, b float32, c float32, d float32) float32 {
	return float32(C.EaseCubicOut(C.float(t), C.float(b), C.float(c), C.float(d)))
}

// EASEDEF float EaseBounceOut(float t, float b, float c, float d) // Ease: Bounce Out
func EaseBounceOut(t float32, b float32, c float32, d float32) float32 {
	return float32(C.EaseBounceOut(C.float(t), C.float(b), C.float(c), C.float(d)))
}

// EASEDEF float EaseQuadOut(float t, float b, float c, float d) { t /= d; return (-c*t*(t - 2.0f) + b); }              // Ease: Quadratic Out
func EaseQuadOut(t float32, b float32, c float32, d float32) float32 {
	return float32(C.EaseQuadOut(C.float(t), C.float(b), C.float(c), C.float(d)))
}

// EASEDEF float EaseCircOut(float t, float b, float c, float d) { t = t/d - 1.0f; return (c*sqrtf(1.0f - t*t) + b); }  // Ease: Circular Out
func EaseCircOut(t float32, b float32, c float32, d float32) float32 {
	return float32(C.EaseCircOut(C.float(t), C.float(b), C.float(c), C.float(d)))
}

// EASEDEF float EaseSineOut(float t, float b, float c, float d) { return (c*sinf(t/d*(PI/2.0f)) + b); }                // Ease: Sine Out
func EaseSineOut(t float32, b float32, c float32, d float32) float32 {
	return float32(C.EaseSineOut(C.float(t), C.float(b), C.float(c), C.float(d)))
}

// EASEDEF float EaseLinearIn(float t, float b, float c, float d) { return (c*t/d + b); }                              // Ease: Linear In
func EaseLinearIn(t float32, b float32, c float32, d float32) float32 {
	return float32(C.EaseLinearIn(C.float(t), C.float(b), C.float(c), C.float(d)))
}
