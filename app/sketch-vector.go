package main

import "strconv"

type SketchVector struct {
	Sketch
	sp *SettingPane
}

func NewSketchVector(ctx *Canvas2d) *SketchVector {
	sp := NewSettingPane("settings", "Vector Settings")
	crossSize := NewInputControl("crossSize", "number", "Cross Size", 50)
	sp.AddInputControl(crossSize)
	arrowSize := NewInputControl("arrowSize", "number", "Arrow Size", 50)
	sp.AddInputControl(arrowSize)

	sketch := &SketchVector{}
	sketch.ctx = ctx
	sketch.sp = sp

	return sketch
}

func (sketch *SketchVector) RenderLoop() {

	sketch.angle += 0.1

	sketch.ctx.SetFillStyle(Color{R: 11, G: 5, B: 38})
	sketch.ctx.FillRect(0, 0, width, height)

	sketch.ctx.SetStrokeStyle(NewColor(255, 255, 255))

	val := sketch.sp.GetValue("crossSize").String()
	crossSize, _ := strconv.ParseFloat(val, 64)
	sketch.ctx.Save()
	sketch.ctx.Translate(width*0.5, height*0.5)
	sketch.ctx.SetStrokeStyle(NewColor(255, 255, 255))
	cross.Width = crossSize
	cross.Height = crossSize
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
