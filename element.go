package XMLBuilder

type Attribute struct {
	Name, Value string
}

type Element struct {
	*Attribute
	Attributes []Attribute
	Children   []Element
	SelfClose  bool
}

type Document struct {
	Root Element
}
