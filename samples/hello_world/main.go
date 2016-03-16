// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	font, err := driver.CreateFont(gxfont.Default, 75)
	if err != nil {
		panic(err)
	}

	window := theme.CreateWindow(380, 100, "Main Application")
	// bg := gxui.Color{}
	// bg.A = 1
	// bg.R = .5
	// brsh := gxui.CreateBrush(bg)

	// window.SetBackgroundBrush(brsh)
	_ = font
	label := theme.CreateLabel()
	// label.SetFont(font)
	label.SetText("Hello world")
	label.SetSize(math.Size{200, 100})

	// window.AddChild(label)
	// btn := theme.CreateButton()
	b := theme.CreateButton()
	b.SetText("         Push Me             ")
	b.SetBorderPen(gxui.WhitePen)
	b.SetSize(math.Size{1024, 20})
	b.OnClick(func(gxui.MouseEvent) { log.Print("Hello") })
	c := window.AddChild(b)
	c.Offset = math.Point{10, 10}

	// b.SetSizeMode(gxui.Fill)
	// b.OnClick(func(gxui.MouseEvent) { action(); update() })

	layout := theme.CreateLinearLayout()
	layout.SetSizeMode(gxui.Fill)
	layout.SetDirection(gxui.TopToBottom)
	layout.SetHorizontalAlignment(gxui.AlignLeft)

	// btn.SetText("HELLO WORLD OK")
	// btn.DesiredSize(math.Size{50, 50}, math.Size{100, 100})
	// btn.SetBackgroundBrush(brsh)
	// child.Control.SetSize(math.Size{100, 50})

	canvas := driver.CreateCanvas(math.Size{W: 500, H: 100})

	drawStar(canvas, math.Point{X: 0, Y: 0}, 50, 0.2, 6)
	drawStar(canvas, math.Point{X: 150, Y: 20}, 70, 0.5, 7)
	drawStar(canvas, math.Point{X: 310, Y: 80}, 25, 0.9, 5)
	drawStar(canvas, math.Point{X: 500, Y: 100}, 45, 0, 6)

	// drawMoon(canvas, math.Point{X: 400, Y: 300}, 200)
	canvas.Complete()

	image := theme.CreateImage()
	image.SetCanvas(canvas)
	image.SetBorderPen(gxui.Pen{2, gxui.Yellow})
	layout.AddChild(image)
	layout.AddChild(label)
	layout.AddChild(b)

	// window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(layout)
	layout.SetSizeMode(gxui.ExpandToContent)
	layout.SetHorizontalAlignment(gxui.AlignCenter)
	layout.SetBorderPen(gxui.WhitePen)

	// ticker := time.NewTicker(time.Millisecond * 30)

	// fn := func(e gxui.MouseEvent) {
	// 	c := gxui.Color{1, 1, 1, .5}
	// 	c.R += float32(float64(e.ScrollY) * .5 / 40)
	// 	if c.R == 1 {
	// 		c.R = 0
	// 	}
	// 	log.Print(e.ScrollY)

	// 	// btn.SetBackgroundBrush(gxui.CreateBrush(c))
	// 	label.SetColor(c)
	// }

	// btn.OnMouseScroll(fn)
	// label.OnMouseScroll(fn)

	// go func() {
	// 	phase := float32(0)
	// 	for _ = range ticker.C {
	// 		c := gxui.Color{
	// 			R: 0.75 + 0.25*math.Cosf((phase+0.000)*math.TwoPi),
	// 			G: 0.75 + 0.25*math.Cosf((phase+0.333)*math.TwoPi),
	// 			B: 0.75 + 0.25*math.Cosf((phase+0.666)*math.TwoPi),
	// 			A: 0.50 + 0.50*math.Cosf(phase*10),
	// 		}
	// 		phase += 0.01
	// 		driver.Call(func() {
	// 			label.SetColor(c)

	// 		})
	// 	}
	// }()

	// window.OnClose(ticker.Stop)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}

func drawStar(canvas gxui.Canvas, center math.Point, radius, rotation float32, points int) {
	p := make(gxui.Polygon, points*2)
	for i := 0; i < points*2; i++ {
		frac := float32(i) / float32(points*2)
		α := frac*math.TwoPi + rotation
		r := []float32{radius, radius / 2}[i&1]
		p[i] = gxui.PolygonVertex{
			Position: math.Point{
				X: center.X + int(r*math.Cosf(α)),
				Y: center.Y + int(r*math.Sinf(α)),
			},
			RoundedRadius: []float32{0, 50}[i&1],
		}
	}
	canvas.DrawPolygon(p, gxui.CreatePen(3, gxui.Red), gxui.CreateBrush(gxui.Yellow))

}
