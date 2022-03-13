package main

import "math"

var cross = NewShape([]Vec{
	{-0.4, -1},
	{0.4, -1},
	{0.4, -0.4},
	{1, -0.4},
	{1, 0.4},
	{0.4, 0.4},
	{0.4, 1},
	{-0.4, 1},
	{-0.4, 0.4},
	{-1, 0.4},
	{-1, -0.4},
	{-0.4, -0.4},
	{-0.4, -1},
})

var arrow = NewShape([]Vec{
	{0, -1},
	{0.4, -0.2},
	{0.2, -0.2},
	{0.2, 1},
	{-0.2, 1},
	{-0.2, -0.2},
	{-0.4, -0.2},
	{0, -1},
})

// var trailVertices = [][]float64{
// 	{-0.4, -1},
// 	{0.4, -0.4},
// }

type Vec struct {
	X float64
	Y float64
}

type Shape struct {
	Vectors []Vec
	Pos     Vec
	Width   float64
	Height  float64
	Angle   float64
}

func NewShape(vectors []Vec) Shape {
	return Shape{
		Vectors: vectors,
	}
}

func (shape Shape) Stroke(ctx *Canvas2d) {
	ctx.BeginPath()

	if len(shape.Vectors) < 2 {
		ctx.Rect(0, 0, 1, 1)
		ctx.Stroke()
		return
	}

	if shape.Angle == 0 {
		x := shape.Vectors[0].X * shape.Width * 0.5
		y := shape.Vectors[0].Y * shape.Height * 0.5

		ctx.MoveTo(x, y)
		for i := 1; i < len(shape.Vectors); i++ {
			x = shape.Vectors[i].X * shape.Width * 0.5
			y = shape.Vectors[i].Y * shape.Height * 0.5
			ctx.LineTo(x, y)
		}
	}

	rad := shape.Angle * math.Pi / 180

	x := shape.Vectors[0].X * shape.Width * 0.5
	y := shape.Vectors[0].Y * shape.Height * 0.5

	rx := x*math.Cos(rad) - y*math.Sin(rad)
	ry := y*math.Cos(rad) + x*math.Sin(rad)
	ctx.MoveTo(rx, ry)

	for i := 1; i < len(shape.Vectors); i++ {
		x = shape.Vectors[i].X * shape.Width * 0.5
		y = shape.Vectors[i].Y * shape.Height * 0.5
		rx := x*math.Cos(rad) - y*math.Sin(rad)
		ry := y*math.Cos(rad) + x*math.Sin(rad)

		ctx.LineTo(rx, ry)
	}

	ctx.Stroke()
}

func (shape Shape) Fill(ctx *Canvas2d) {
	ctx.BeginPath()

	if len(shape.Vectors) < 2 {
		ctx.Rect(0, 0, 1, 1)
		ctx.Stroke()
		return
	}

	if shape.Angle == 0 {
		x := shape.Vectors[0].X * shape.Width * 0.5
		y := shape.Vectors[0].Y * shape.Height * 0.5

		ctx.MoveTo(x, y)
		for i := 1; i < len(shape.Vectors); i++ {
			x = shape.Vectors[i].X * shape.Width * 0.5
			y = shape.Vectors[i].Y * shape.Height * 0.5
			ctx.LineTo(x, y)
		}
	}

	rad := shape.Angle * math.Pi / 180

	x := shape.Vectors[0].X * shape.Width * 0.5
	y := shape.Vectors[0].Y * shape.Height * 0.5

	rx := x*math.Cos(rad) - y*math.Sin(rad)
	ry := y*math.Cos(rad) + x*math.Sin(rad)
	ctx.MoveTo(rx, ry)

	for i := 1; i < len(shape.Vectors); i++ {
		x = shape.Vectors[i].X * shape.Width * 0.5
		y = shape.Vectors[i].Y * shape.Height * 0.5
		rx := x*math.Cos(rad) - y*math.Sin(rad)
		ry := y*math.Cos(rad) + x*math.Sin(rad)

		ctx.LineTo(rx, ry)
	}

	ctx.Fill()
}
