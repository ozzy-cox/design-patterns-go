package main

import "fmt"

type Graphic interface {
	Draw()
	GetExtent() (int, int)
	Store()
	Load()
}

type Image struct {
	height int
	width  int
}

func (i Image) Draw() {
	fmt.Println("Drawing image")
}

func (i Image) GetExtent() (int, int) {
	return i.height, i.width
}

func (i Image) Store() {
	fmt.Println("Storing image")
}

func (i Image) Load() {
	fmt.Println("Loading image")
}

type ImageProxy struct {
	image  *Image
	height int
	width  int
}

func (i ImageProxy) Draw() {
	fmt.Println("Redirecting draw")
	if i.image == nil {
		i.image = &Image{0, 0}
	}
	i.image.Draw()
}

func (i ImageProxy) GetExtent() (int, int) {
	if i.image == nil {
		return i.height, i.width
	}
	return i.image.GetExtent()
}

func (i ImageProxy) Store() {
	fmt.Println("Redirecting store")
	if i.image == nil {
		i.image = &Image{0, 0}
	}
	i.image.Store()
}

func (i ImageProxy) Load() {
	fmt.Println("Redirecting load")
	if i.image == nil {
		i.image = &Image{0, 0}
	}
	i.image.Load()
}

func main() {
}
