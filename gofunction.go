package gofunction

import "github.com/PrimForDev/gofunction/operation"

func Addition(a int32, b int32) int32 {
	var result = a + b
	return result
}

func Sub(a int32, b int32) int32 {
	var result = operation.Subtraction(a, b)
	return result
}
