package main

import (
	"fmt"
	"math"
)

func main() {
	maxInt8 := math.MaxInt8
	minInt8 := math.MinInt8
	maxInt16 := math.MaxInt16
	minInt16 := math.MinInt16
	maxInt32 := math.MaxInt32
	minInt32 := math.MinInt32
	maxInt64 := math.MaxInt64
	minInt64 := math.MinInt64
	maxUint8 := math.MaxUint8
	maxUint16 := math.MaxUint16
	maxUint32 := math.MaxUint32
	maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64

	fmt.Printf("Range of Int8 :: %v to %v.\n", minInt8, maxInt8)
	fmt.Printf("Range of Int16 :: %v to %v.\n", minInt16, maxInt16)
	fmt.Printf("Range of Int32 :: %v to %v.\n", minInt32, maxInt32)
	fmt.Printf("Range of Int64 :: %v to %v.\n", minInt64, maxInt64)
	fmt.Printf("Max Uint8 :: %v.\n", maxUint8)
	fmt.Printf("Max Uint16 :: %v.\n", maxUint16)
	fmt.Printf("Max Uint32 :: %v.\n", maxUint32)
	fmt.Printf("Max Float32 :: %f.\n", maxFloat32)
	fmt.Printf("Max Float64 :: %3.2f.\n", maxFloat64)
}
