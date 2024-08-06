package XMLBuilder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElementBuilder(t *testing.T) {
	doc := XDocument(
		XElement(
			"ExtensionInfo",
			XElement("Name").
				Content("My First Extension"),
		).Close(),
	)

	expects := &Document{Root: Element{
		Attribute: &Attribute{Name: "ExtensionInfo"},
		Children: []Element{
			{
				Attribute: &Attribute{Name: "Name", Value: "My First Extension"},
				SelfClose: false,
			},
		},
		SelfClose: false,
	}}

	assert.Equal(t, expects, doc)
}
