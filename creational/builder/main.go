package builder

type (
	Font          any
	TextConverter interface {
		convertChar(c byte)
		convertFontChange(Font Font)
		convertParagraph()
	}
)

type RTFReader struct {
	textConverter TextConverter
}

func (a ASCIIConverter) convertChar(c byte) {
	panic("not implemented") // TODO: Implement
}

func (a ASCIIConverter) convertFontChange(Font Font) {
	panic("not implemented") // TODO: Implement
}

func (a ASCIIConverter) convertParagraph() {
	panic("not implemented") // TODO: Implement
}

type IASCIIConverter interface {
	GetAsciiText()
}

func (a ASCIIConverter) GetAsciiText() {
	panic("not implemented") // TODO: Implement
}

type ASCIIConverter struct{}

type Token string

const (
	CHAR Token = "CHAR"
	FONT Token = "FONT"
	PARA Token = "PARA"
)

func (r RTFReader) ParseRTF(t Token) {
	switch t {
	case CHAR:
		r.textConverter.convertChar(0)
	case FONT:
		r.textConverter.convertFontChange(0)
	case PARA:
		r.textConverter.convertParagraph()
	}
}

func main() {
	reader := RTFReader{
		textConverter: ASCIIConverter{},
	}
	reader.ParseRTF("Token")
}
