package XMLBuilder

type ElementBuilder struct {
	elem *Element
}

// XAttribute appends a new attribute with the given name and value to the current XML element.
func (builder *ElementBuilder) XAttribute(name, value string) *ElementBuilder {
	attr := Attribute{name, value}

	if builder.elem.Attributes == nil {
		builder.elem.Attributes = make([]Attribute, 1)
		builder.elem.Attributes[0] = attr
	} else {
		builder.elem.Attributes = append(builder.elem.Attributes, attr)
	}

	return builder
}

// Content sets the text content of the current element, and automatically closes the element as there's nothing further to write.
func (builder *ElementBuilder) Content(content string) Element {
	builder.elem.Value = content
	return *builder.elem
}

// Close closes the current element in progress, returning the finalised element.
func (builder *ElementBuilder) Close(selfClosing ...bool) Element {
	if len(selfClosing) > 0 {
		builder.elem.SelfClose = selfClosing[0]
	}

	return *builder.elem
}

// XElement constructs a new XML element with the given name and (optionally) a list of child nodes.
func XElement(
	name string,
	children ...Element,
) *ElementBuilder {
	return &ElementBuilder{
		&Element{
			Attribute: &Attribute{Name: name},
			Children:  children,
		},
	}
}

// XDocument constructs a brand-new XML document with the given element as it's root.
func XDocument(root Element) *Document {
	return &Document{
		root,
	}
}
