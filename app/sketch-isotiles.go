package main

import "math"

// const (
// 	tileWidth  float64 = 100
// 	tileHeight float64 = 50
// )

type SketchIsoTiles struct {
	Sketch
	Settings *SettingPane
	Grid     []*Tile
}

type Tile struct {
	topV   Shape
	leftV  Shape
	rightV Shape
	topC   Color
	leftC  Color
	rightC Color
}

func NewSketchIsoTiles(ctx *Canvas2d) *SketchIsoTiles {

	tile := Tile{
		topV:   NewShape([]Vec{{0, -1}, {1, 0}, {0, 1}, {-1, 0}, {0, -1}}),
		leftV:  NewShape([]Vec{{-1, -1}, {-1, -0.6}, {0, 0.4}, {0, 0}}),
		rightV: NewShape([]Vec{{1, -1}, {1, -0.6}, {0, 0.4}, {0, 0}}),
		topC:   NewColor(255, 0, 0),
		leftC:  NewColor(100, 100, 100),
		rightC: NewColor(200, 200, 200),
	}

	sp := NewSettingPane("settings", "ISO Tiles")
	sp.AddInputControl("rows", "number", "Rows", int64(200))
	sp.AddInputControl("cols", "number", "Cols", int64(10))
	sp.AddInputControl("tileWidth", "number", "Tile Width", float64(50))
	sp.AddInputControl("tileHeight", "number", "Tile Height", float64(25))

	sketch := &SketchIsoTiles{}
	sketch.ctx = ctx
	sketch.Settings = sp

	tiles := sp.GetValueAsInt("rows", 200) + sp.GetValueAsInt("cols", 10)
	sketch.Grid = make([]*Tile, tiles)
	for i := int64(0); i < tiles; i++ {
		sketch.Grid[i] = &tile
	}

	return sketch
}

func (sketch *SketchIsoTiles) Activate() {
	sketch.Settings.Activate()
}

func (sketch *SketchIsoTiles) UpdateSetting(id string, value interface{}) {
	sketch.Settings.SetValue(id, value)
}

func (sketch *SketchIsoTiles) Render() {

	title := sketch.Settings.GetValueAsString("title", "go-sketch")
	h1 := GetDocument().GetElementById("sketchTitle")
	h1.Set("textContent", title)

	ctx := sketch.ctx

	ctx.SetFillStyle(Color{R: 11, G: 5, B: 38})
	ctx.FillRect(0, 0, width, height)

	// rows := sketch.Settings.GetValueAsInt("rows", 10)
	cols := sketch.Settings.GetValueAsInt("cols", 10)
	tileWidth := sketch.Settings.GetValueAsFloat("tileWidth", 50)
	tileHeight := sketch.Settings.GetValueAsFloat("tileHeight", 50)
	for i := int64(0); i < 30; i++ {
		tx := float64(i % cols)
		ty := math.Floor(float64(i) / float64(cols))
		tile := sketch.Grid[i]
		x, y := ProjectIsoXY(400, 400, tx, ty, tileWidth, tileHeight)
		tile.SetSize(tileWidth, tileHeight)
		ctx.Save()
		ctx.Translate(x, y)
		ctx.SetFillStyle(sketch.Grid[0].topC)
		tile.Fill(ctx)
		ctx.Restore()
	}

}

func (tile *Tile) SetSize(width, height float64) {
	tile.topV.Width = width
	tile.topV.Height = height
	tile.leftV.Width = width * 0.5
	tile.leftV.Height = height * 0.4
	tile.rightV.Width = width * 0.5
	tile.rightV.Height = height * 0.4

}

func (tile *Tile) Fill(ctx *Canvas2d) {
	tile.topV.Fill(ctx)
}
