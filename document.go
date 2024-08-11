package XMLBuilder

import (
	"bufio"
	"io"
)

type Document struct {
	writer *bufio.Writer
}

func NewDocument(output io.Writer) *Document {
	buff := bufio.NewWriter(output)

	buff.WriteString(xmlVersionHeader)
	buff.WriteRune('\n')

	return &Document{buff}
}

func (d *Document) Element(name string) ElementBuilder {
	d.writer.WriteRune(StartElement)
	d.writer.WriteString(name)

	return &elementBuilder{
		parent: nil,
		writer: d.writer,
		name:   name,
	}
}
