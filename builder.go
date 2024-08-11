package XMLBuilder

import (
	"bufio"
)

type elementBuilder struct {
	parent *elementBuilder
	name   string
	writer *bufio.Writer
	closed bool
}

const (
	xmlVersionHeader = `<?xml version="1.0" encoding="UTF-8"?>`
)

func (eb *elementBuilder) Element(name string) ElementBuilder {
	if !eb.closed {
		_, err := eb.writer.WriteRune(EndElement)
		if err != nil {
			panic(err)
		}
		eb.closed = true
	}
	_, err := eb.writer.WriteString("\n\t")
	if err != nil {
		panic(err)
	}
	_, err = eb.writer.WriteRune(StartElement)
	if err != nil {
		panic(err)
	}
	_, err = eb.writer.WriteString(name)
	if err != nil {
		panic(err)
	}

	return &elementBuilder{
		parent: eb,
		writer: eb.writer,
		name:   name,
	}
}

func (eb *elementBuilder) Attr(name, value string) ElementBuilder {
	_, err := eb.writer.WriteRune(Space)
	if err != nil {
		panic(err)
	}
	_, err = eb.writer.WriteString(name)
	if err != nil {
		panic(err)
	}
	_, err = eb.writer.WriteRune(Equals)
	if err != nil {
		panic(err)
	}
	_, err = eb.writer.WriteString(quote(value))
	if err != nil {
		panic(err)
	}

	return eb
}

func (eb *elementBuilder) Content(content string) ElementBuilder {
	if !eb.closed {
		_, err := eb.writer.WriteRune(EndElement)
		if err != nil {
			panic(err)
		}
		eb.closed = true
	}
	_, err := eb.writer.WriteString(xmlEncoder.Replace(content))
	if err != nil {
		panic(err)
	}

	return eb
}

func (eb *elementBuilder) Close(selfClosing bool) ElementBuilder {
	if selfClosing {
		_, err := eb.writer.WriteRune('/')
		if err != nil {
			panic(err)
		}
		_, err = eb.writer.WriteRune(EndElement)
		if err != nil {
			panic(err)
		}
	} else {
		if !eb.closed {
			_, err := eb.writer.WriteRune(EndElement)
			if err != nil {
				panic(err)
			}
		}
		if eb.parent == nil {
			_, err := eb.writer.WriteString("\n")
			if err != nil {
				panic(err)
			}
		}
		_, err := eb.writer.WriteRune(StartElement)
		if err != nil {
			panic(err)
		}
		_, err = eb.writer.WriteRune('/')
		if err != nil {
			panic(err)
		}
		_, err = eb.writer.WriteString(eb.name)
		if err != nil {
			panic(err)
		}
		_, err = eb.writer.WriteRune(EndElement)
		if err != nil {
			panic(err)
		}
	}

	// if this is the root element, flush the buffer
	if eb.parent == nil {
		err := eb.writer.Flush()
		if err != nil {
			panic(err)
		}
	}

	return eb.parent
}
