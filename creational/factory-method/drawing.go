package main

type DrawingDocument struct{}

func (d DrawingDocument) Open() {
	panic("not implemented") // TODO: Implement
}

func (d DrawingDocument) Close() {
	panic("not implemented") // TODO: Implement
}

func (d DrawingDocument) Read() {
	panic("not implemented") // TODO: Implement
}

func (d DrawingDocument) Print() {
	panic("not implemented") // TODO: Implement
}

type DrawingApplication struct{}

func (a DrawingApplication) createDocument() *IDocument {
	panic("not implemented") // TODO: Implement
}
