package mymath

import "math"

func Abs(x float64) float64 {
	return math.Abs(x)
}

func Acos(x float64) float64 {
	return math.Acos(x)
}

func Acosh(x float64) float64 {
	return math.Acosh(x)
}

func Asin(x float64) float64 {
	return math.Asin(x)
}

func Asinh(x float64) float64 {
	return math.Asinh(x)
}

func Atan(x float64) float64 {
	return math.Atan(x)
}

func Atan2(x float64, y float64) float64 {
	return math.Atan2(x, y)
}

func Atanh(x float64) float64 {
	return math.Atanh(x)
}

func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Copysign(x float64, y float64) float64 {
	return math.Copysign(x, y)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func Cosh(x float64) float64 {
	return math.Cosh(x)
}

func Dim(x float64, y float64) float64 {
	return math.Dim(x, y)
}

func Erf(x float64) float64 {
	return math.Erf(x)
}

func Erfc(x float64) float64 {
	return math.Erfc(x)
}

func Erfcinv(x float64) float64 {
	return math.Erfcinv(x)
}

func Erfinv(x float64) float64 {
	return math.Erfinv(x)
}

func Exp(x float64) float64 {
	return math.Exp(x)
}

func Exp2(x float64) float64 {
	return math.Exp2(x)
}

func Expm1(x float64) float64 {
	return math.Expm1(x)
}

func FMA(x float64, y float64, z float64) float64 {
	return math.FMA(x, y, z)
}

func Float32bits(x float32) uint32 {
	return math.Float32bits(x)
}

func Float32frombits(x uint32) float32 {
	return math.Float32frombits(x)
}

func Float64bits(x float64) uint64 {
	return math.Float64bits(x)
}

func Float64frombits(x uint64) float64 {
	return math.Float64frombits(x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Frexp(x float64) (float64, int) {
	return math.Frexp(x)
}

func Gamma(x float64) float64 {
	return math.Gamma(x)
}

func Hypot(x float64, y float64) float64 {
	return math.Hypot(x, y)
}

func Ilogb(x float64) int {
	return math.Ilogb(x)
}

func Inf(x int) float64 {
	return math.Inf(x)
}

func IsInf(f float64, sign int) bool {
	return math.IsInf(f, sign)
}

func IsNaN(f float64) bool {
	return math.IsNaN(f)
}

func J0(x float64) float64 {
	return math.J0(x)
}

func J1(x float64) float64 {
	return math.J1(x)
}

func Jn(n int, x float64) float64 {
	return math.Jn(n, x)
}

func Ldexp(x float64, n int) float64 {
	return math.Ldexp(x, n)
}

func Lgamma(x float64) (float64, int) {
	return math.Lgamma(x)
}

func Log(x float64) float64 {
	return math.Log(x)
}

func Log10(x float64) float64 {
	return math.Log10(x)
}

func Log1p(x float64) float64 {
	return math.Log1p(x)
}

func Log2(x float64) float64 {
	return math.Log2(x)
}

func Logb(x float64) float64 {
	return math.Logb(x)
}

func Max(x float64, y float64) float64 {
	return math.Max(x, y)
}

func Mod(x float64, y float64) float64 {
	return math.Mod(x, y)
}

func Modf(x float64, y float64) (float64, float64) {
	return math.Modf(x)
}

func Min(x float64, y float64) float64 {
	return math.Min(x, y)
}

func NaN() float64 {
	return math.NaN()
}

func Nextafter(x float64, y float64) float64 {
	return math.Nextafter(x, y)
}

func Nextafter32(x float32, y float32) float32 {
	return math.Nextafter32(x, y)
}

func Pow(x float64, y float64) float64 {
	return math.Pow(x, y)
}

func Pow10(n int) float64 {
	return math.Pow10(n)
}

func Remainder(x float64, y float64) float64 {
	return math.Remainder(x, y)
}

func Round(x float64) float64 {
	return math.Round(x)
}

func RoundToEven(x float64) float64 {
	return math.RoundToEven(x)
}

func Signbit(x float64) bool {
	return math.Signbit(x)
}

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Sincos(x float64) (float64, float64) {
	return math.Sincos(x)
}

func Sinh(x float64) float64 {
	return math.Sinh(x)
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Tan(x float64) float64 {
	return math.Tan(x)
}

func Tanh(x float64) float64 {
	return math.Tanh(x)
}

func Trunc(x float64) float64 {
	return math.Trunc(x)
}

func Y0(x float64) float64 {
	return math.Y0(x)
}

func Y1(x float64) float64 {
	return math.Y1(x)
}

func Yn(n int, x float64) float64 {
	return math.Yn(n, x)
}
