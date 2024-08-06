package XMLBuilder

import (
	"bufio"
	"io"
)

func Marshal(doc *Document, writer io.Writer) (err error) {
	buff := bufio.NewWriter(writer)

	_, err = buff.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	if err != nil {
		return err
	}
	_, err = buff.WriteRune('\n')
	if err != nil {
		return err
	}

	err = writeElement(buff, doc.Root, 0)
	if err != nil {
		return err
	}

	return nil
}

func writeElement(buff *bufio.Writer, element Element, depth int) (err error) {
	closed := false

	for i := 0; i < depth; i++ {
		_, err = buff.WriteRune('\t')
		if err != nil {
			return err
		}
	}

	_, err = buff.WriteString(StartElement)
	if err != nil {
		return err
	}
	_, err = buff.WriteString(xmlEncoder.Replace(element.Name))
	if err != nil {
		return err
	}

	err = writeAttributes(buff, element.Attributes)
	if err != nil {
		return err
	}

	if element.SelfClose && len(element.Value) == 0 && len(element.Children) == 0 {
		_, err = buff.WriteRune('/')
		if err != nil {
			return err
		}
		_, err = buff.WriteString(EndElement)
		if err != nil {
			return err
		}
		closed = true
	} else {
		_, err = buff.WriteString(EndElement)
		if err != nil {
			return err
		}
	}

	if len(element.Children) > 0 {
		err = writeChildren(buff, element.Children, depth+1)
		if err != nil {
			return err
		}
	} else if len(element.Value) > 0 {
		_, err = buff.WriteString(xmlEncoder.Replace(element.Value))
		if err != nil {
			return err
		}
	}

	if !closed {
		if depth == 0 {
			_, err = buff.WriteRune('\n')
			if err != nil {
				return err
			}
		}
		_, err = buff.WriteString(StartElement)
		if err != nil {
			return err
		}
		_, err = buff.WriteRune('/')
		if err != nil {
			return err
		}
		_, err = buff.WriteString(xmlEncoder.Replace(element.Name))
		if err != nil {
			return err
		}
		_, err = buff.WriteString(EndElement)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeAttributes(buff *bufio.Writer, attributes []Attribute) (err error) {
	for _, attr := range attributes {
		_, err = buff.WriteString(Space)
		if err != nil {
			return err
		}
		_, err = buff.WriteString(xmlEncoder.Replace(attr.Name))
		if err != nil {
			return err
		}
		_, err = buff.WriteString(Equals)
		if err != nil {
			return err
		}
		_, err = buff.WriteString(quote(attr.Value))
		if err != nil {
			return err
		}
	}

	return nil
}

func writeChildren(buff *bufio.Writer, children []Element, depth int) (err error) {
	for _, child := range children {
		_, err = buff.WriteRune('\n')
		if err != nil {
			return err
		}
		err = writeElement(buff, child, depth)
	}

	return nil
}
