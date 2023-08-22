package main

import "fmt"

type Drawable interface {
	Draw()
}

type Graphic struct{}

func (g *Graphic) Draw() {
	fmt.Println("Drawing graphic")
}

type BorderGraphic struct {
	graphic Drawable
}

func (g *BorderGraphic) Draw() {
	fmt.Println("Drawing border")
	g.graphic.Draw()
}

type ScrollGraphic struct {
	graphic Drawable
}

func (g *ScrollGraphic) Draw() {
	fmt.Println("Drawing scroll")
	g.graphic.Draw()
}

func main() {
	bareGraphic := Graphic{}

	bareGraphic.Draw()
	fmt.Println()

	borderGraphic := BorderGraphic{graphic: &Graphic{}}

	borderGraphic.Draw()
	fmt.Println()

	scrollGraphic := ScrollGraphic{graphic: &Graphic{}}

	scrollGraphic.Draw()
	fmt.Println()

	borderScrollGraphic := BorderGraphic{graphic: &ScrollGraphic{graphic: &Graphic{}}}

	borderScrollGraphic.Draw()
	fmt.Println()

	borderBorderScrollGraphic := BorderGraphic{graphic: &BorderGraphic{graphic: &ScrollGraphic{graphic: &Graphic{}}}}

	borderBorderScrollGraphic.Draw()
	fmt.Println()
}
