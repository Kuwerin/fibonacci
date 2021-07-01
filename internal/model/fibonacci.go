package model

type Fibonacci struct {
	X uint64 `form:"x" json:"x"`
	Y uint64 `form:"y" json:"y" binding:"required"`
}
