package main

type SketchVector struct {
	Sketch
}

func NewSketchVector(ctx *Canvas2d) *SketchVector {
	return &SketchVector{
		Sketch{
			ctx: ctx,
		},
	}
}

func (sketch *SketchVector) RenderLoop() {

	sketch.angle += 0.1

	sketch.ctx.SetFillStyle(Color{R: 11, G: 5, B: 38})
	sketch.ctx.FillRect(0, 0, width, height)

	sketch.ctx.SetStrokeStyle(NewColor(255, 255, 255))

	sketch.ctx.Save()
	sketch.ctx.Translate(width*0.5, height*0.5)
	sketch.ctx.SetStrokeStyle(NewColor(255, 255, 255))
	cross.Width = 200
	cross.Height = 200
	cross.Angle = sketch.angle
	cross.Stroke(sketch.ctx)
	sketch.ctx.Restore()

	sketch.ctx.Save()
	sketch.ctx.Translate(width*0.5, height*0.5)
	sketch.ctx.SetStrokeStyle(NewColor(50, 255, 500))
	arrow.Width = 600
	arrow.Height = 600
	arrow.Angle = 360 - sketch.angle
	arrow.Stroke(sketch.ctx)
	sketch.ctx.Restore()

}