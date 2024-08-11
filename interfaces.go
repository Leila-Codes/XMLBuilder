package XMLBuilder

type ElementBuilder interface {
	Element(name string) ElementBuilder
	Attr(name, value string) ElementBuilder
	Content(data string) ElementBuilder
	Close(selfClosing bool) ElementBuilder
}
