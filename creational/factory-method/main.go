package main

type Document struct{}

func (d Document) Open() {
	panic("not implemented") // TODO: Implement
}

func (d Document) Close() {
	panic("not implemented") // TODO: Implement
}

func (d Document) Read() {
	panic("not implemented") // TODO: Implement
}

func (d Document) Print() {
	panic("not implemented") // TODO: Implement
}

type IDocument interface {
	Open()
	Close()
	Read()
	Print()
}

type IApplication interface {
	createDocument() *IDocument
}

type Application struct{}

func (*Application) createDocument() IDocument {
	return Document{}
}

func main() {
	var app IApplication

	app = DrawingApplication{}

	app.createDocument()
}
